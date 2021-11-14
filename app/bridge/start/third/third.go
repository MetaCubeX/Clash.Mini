package third

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	path "path/filepath"
	"regexp"
	"strings"

	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/second"

	"github.com/Clash-Mini/Clash.Mini/cmd/breaker"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
	httpUtils "github.com/Clash-Mini/Clash.Mini/util/http"
	"github.com/Clash-Mini/Clash.Mini/util/loopback"
	protocolUtils "github.com/Clash-Mini/Clash.Mini/util/protocol"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"
	"github.com/Clash-Mini/Clash.Mini/util/uac"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type ClashProtocol struct {
	Url  string `query:"url"`
	Name string `query:"name,omitempty"`
}

var (
	protocolRegexp = regexp.MustCompile(`clash://install-config/?\?(.*)`)
)

func init() {
	log.Infoln("[bridge] third")

	uacChecks()
}

func uacChecks() {
	// 检查是否UAC需求
	bindProtocol()
	bindRegisterProtocol()
	bindLoopback()

	uac.RunWhenAdmin()
}

func bindProtocol() {
	uac.BindFuncWithArg("--protocol", uac.AnyTime, func(arg *uac.Arg, args []string) (done bool) {
		var err error
		var msg string
		arg.EqualValue = strings.TrimSpace(arg.EqualValue)
		clashProtocol := &ClashProtocol{}
		var into, alerted bool
		into, err = checkProtocol(clashProtocol, arg.EqualValue)
		if err == nil {
			if into {
				log.Infoln("[protocol] into protocol: %v", clashProtocol)
				rlt := walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips),
					i18n.TData(cI18n.UacMsgBoxProtocolInstallConfigConfirmMsg, &i18n.Data{Data: map[string]interface{}{
						"Name": clashProtocol.Name,
						"Url":  clashProtocol.Url,
					}}), walk.MsgBoxIconQuestion|walk.MsgBoxOKCancel)
				if rlt != win.IDOK {
					log.Infoln("[protocol] user skipped")
					return
				}
				if err == nil {
					alerted, err = installConfig(clashProtocol)
				}
				if err == nil {
					msg = i18n.TData(cI18n.UacMsgBoxProtocolInstallConfigSuccessfulConfigMsg, &i18n.Data{Data: map[string]interface{}{
						"Name": clashProtocol.Name,
					}})
				} else {
					log.Errorln("[protocol] " + err.Error())
					if !alerted {
						msg = i18n.TData(cI18n.UacMsgBoxProtocolInstallConfigFailedConfigMsg, &i18n.Data{Data: map[string]interface{}{
							"Name":  clashProtocol.Name,
							"Error": err.Error(),
						}})
					}
				}
			} else {
				msg = "not found any valid info of protocol clash://"
				err = fmt.Errorf(msg)
			}
		} else {
			msg = err.Error()
		}
		if !alerted {
			alert(err == nil, msg)
		}
		if err != nil {
			os.Exit(1)
		}
		return true
	})
}

func bindRegisterProtocol() {
	uac.BindFuncWithArg("--uac-protocol-enable", uac.OnlyUac, func(arg *uac.Arg, args []string) (done bool) {
		err := protocolUtils.RegisterCommandProtocol(true)
		var msg string
		if err == nil {
			msg = "registered Clash protocol"
		} else {
			msg = err.Error()
		}
		alert(err == nil, msg)
		if err != nil {
			os.Exit(1)
		}
		return true
	})
	uac.BindFuncWithArg("--uac-protocol-disable", uac.OnlyUac, func(arg *uac.Arg, args []string) (done bool) {
		err := protocolUtils.RegisterCommandProtocol(false)
		var msg string
		if err == nil {
			msg = "unregistered Clash protocol"
		} else {
			msg = err.Error()
		}
		alert(err == nil, msg)
		if err != nil {
			os.Exit(1)
		}
		return true
	})
}

func bindLoopback() {
	uac.BindFuncWithArg("--uac-loopback-enable", uac.OnlyUac, func(arg *uac.Arg, args []string) (done bool) {
		ticker := loopback.Breaker(breaker.ON)
		select {
		case <-ticker.C:
		}
		return true
	})
	uac.BindFuncWithArg("--uac-loopback-disable", uac.OnlyUac, func(arg *uac.Arg, args []string) (done bool) {
		ticker := loopback.Breaker(breaker.OFF)
		select {
		case <-ticker.C:
		}
		return true
	})
}

func alert(ok bool, msg string) {
	var icon walk.MsgBoxStyle
	if ok {
		icon = walk.MsgBoxIconInformation
	} else {
		icon = walk.MsgBoxIconError
	}
	walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips), msg, icon)
}

func installConfig(clashProtocol *ClashProtocol) (alerted bool, err error) {
	configPath := path.Join(constant.ProfileDir, clashProtocol.Name+constant.ConfigSuffix)
	// 检查是否存在，询问是否覆盖
	exists, err := fileUtils.IsExists(configPath)
	if err != nil {
		return false, err
	}
	if exists {
		log.Infoln("[protocol] config is exists: %s", clashProtocol.Name)
		rlt := walk.MsgBox(nil, i18n.T(cI18n.UacMsgBoxTitle),
			i18n.TData(cI18n.UacMsgBoxProtocolInstallConfigOverwriteMsg, &i18n.Data{Data: map[string]interface{}{
				"Name": clashProtocol.Name,
				"Url":  clashProtocol.Url,
			}}), walk.MsgBoxIconQuestion|walk.MsgBoxOKCancel)
		if rlt != win.IDOK {
			log.Infoln("[protocol] user skipped overwrite")
			return false, nil
		}
	}
	// download
	req, err := httpUtils.NewRequest(http.MethodGet, clashProtocol.Url, nil)
	if err != nil {
		return false, err
	}
	httpUtils.AddClashHeader(req)
	rsp, err := httpUtils.DoRequest(req)
	defer httpUtils.DeferSafeCloseResponseBody(rsp)
	if err != nil {
		return false, err
	}
	var rspBody string
	if rsp != nil {
		rspBody = string(stringUtils.IgnoreErrorBytes(ioutil.ReadAll(rsp.Body)))
	}
	statusCode := -1
	if rsp != nil {
		statusCode = rsp.StatusCode
	}
	if err != nil || (statusCode != http.StatusOK) {
		log.Warnln("[protocol] AddConfig Do error: %v, request url: %s, response: [%d] %s",
			err, req.URL.String(), statusCode, rspBody)
		var errMsg string
		if err == http.ErrHandlerTimeout ||
			(util.EqualsAny(statusCode, http.StatusInternalServerError, http.StatusServiceUnavailable)) {
			errMsg = i18n.T(cI18n.MenuConfigWindowAddConfigUrlTimeout)
		} else if err == http.ErrNoLocation || err == http.ErrMissingFile ||
			(statusCode == http.StatusNotFound) {
			errMsg = i18n.T(cI18n.MenuConfigWindowAddConfigUrlCodeFail)
		} else {
			errMsg = i18n.T(cI18n.MenuConfigWindowAddConfigUrlDownloadFail)
		}
		return false, fmt.Errorf(errMsg)
	}
	if err == nil {
		if statusCode == 200 {
			var isClashConfig bool
			isClashConfig, err = regexp.MatchString(`proxy-groups`, rspBody)
			if err != nil || !isClashConfig {
				if err != nil {
					log.Errorln(err.Error())
				}
				walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips), i18n.T(cI18n.MenuConfigWindowAddConfigUrlNotClash), walk.MsgBoxIconError)
				return true, fmt.Errorf("[protocol] %s: %v", i18n.T(cI18n.MenuConfigWindowAddConfigUrlNotClash), err)
			}
			//save to profile
			rspBodyReader := ioutil.NopCloser(strings.NewReader(rspBody))
			var f *os.File
			f, err = os.Create(configPath)
			if err != nil {
				return false, err
			}
			_, err = f.WriteString(fmt.Sprintf("# Clash.Mini : %s\n", clashProtocol.Url))
			_, err = io.Copy(f, rspBodyReader)
			err = f.Close()
		} else {
			log.Errorln("[protocol] response invalid")
			err = fmt.Errorf("response invalid")
		}
	}
	return false, err
}

func checkProtocol(clashProtocol *ClashProtocol, protocolInfo string) (into bool, err error) {
	if len(protocolInfo) > 0 {
		into = true
		log.Infoln(protocolInfo)
		if len(protocolInfo) == 0 {
			return false, fmt.Errorf("[protocol] url is blank")
		}
		log.Warnln("[%s] : %t", protocolInfo, protocolRegexp.MatchString(protocolInfo))
		if !protocolRegexp.MatchString(protocolInfo) {
			return false, fmt.Errorf("[protocol] not found any valid info of protocol \"clash://\", "+
				"url is not supported: \n%s", protocolInfo)
		}
		protocolQueryString := protocolRegexp.FindStringSubmatch(protocolInfo)[1]
		log.Infoln(protocolQueryString)
		err = util.UnmarshalByValues(protocolQueryString, clashProtocol)
		if err != nil {
			log.Errorln(err.Error())
			return into, err
		}
		clashProtocol.Url, err = url.PathUnescape(clashProtocol.Url)
		clashProtocol.Url = strings.TrimSpace(clashProtocol.Url)
		log.Infoln("[protocol] clashProtocol: %v", clashProtocol)
		if len(clashProtocol.Url) == 0 {
			return false, fmt.Errorf("[protocol] url is blank")
		}
		if !httpUtils.UrlRegexp.MatchString(clashProtocol.Url) {
			return false, fmt.Errorf("[protocol] url is invalid")
		}
		if len(clashProtocol.Name) == 0 {
			regexpGroups := httpUtils.UrlRegexp.FindStringSubmatch(clashProtocol.Url)
			if len(regexpGroups) > 1 {
				clashProtocol.Name = regexpGroups[1]
			}
		}
	}
	return into, nil
}
