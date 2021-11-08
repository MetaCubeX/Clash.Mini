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

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
	httpUtils "github.com/Clash-Mini/Clash.Mini/util/http"
	"github.com/Clash-Mini/Clash.Mini/util/protocol"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"
	"github.com/Clash-Mini/Clash.Mini/util/uac"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type ClashProtocol struct {
	Url 	string	`query:"url"`
	Name 	string 	`query:"name,omitempty"`
}

var (
	protocolRegexp		= regexp.MustCompile(`clash://install-config\?(.*)`)
)

func init() {
	uacChecks()
	preChecks()
}

func uacChecks() {
	// 检查是否UAC需求
	//uacUtils.CheckAndRunMeElevated(taskExe, taskArgs)
	//uac.BindFuncWithArg("--uac-task", func(arg *uac.Arg, args []string) {
	//
	//})
	uac.BindFuncWithArg("--uac-protocol-enable", func(arg *uac.Arg, args []string) (done bool) {
		err := protocol.RegisterCommandProtocol(true)
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
	uac.BindFuncWithArg("--uac-protocol-disable", func(arg *uac.Arg, args []string) (done bool) {
		err := protocol.RegisterCommandProtocol(false)
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

	uac.RunWhenAdmin()
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

func preChecks() {
	// 检查是否协议注册需求
	clashProtocol := &ClashProtocol{}
	into, err := checkProtocol(clashProtocol)
	if into {
		log.Infoln("[protocol] into protocol: %v", clashProtocol)
		rlt := walk.MsgBox(nil, i18n.T(cI18n.UacMsgBoxTitle), fmt.Sprintf(i18n.T(cI18n.UacMsgBoxProtocolInstallConfigConfirmMsg), common.Protocol), walk.MsgBoxIconQuestion | walk.MsgBoxYesNo)
		if rlt != win.IDYES {
			log.Infoln("[protocol] user skipped")
			return
		}
		var alerted bool
		if err == nil {
			alerted, err = installConfig(clashProtocol)
		}
		if err == nil {
			walk.MsgBox(nil, i18n.T(cI18n.UacMsgBoxTitle), fmt.Sprintf(i18n.T(cI18n.UacMsgBoxProtocolInstallConfigSuccessfulConfigMsg), clashProtocol.Url), walk.MsgBoxIconInformation)
			os.Exit(0)
		} else {
			log.Errorln("[protocol] " + err.Error())
			if !alerted {
				walk.MsgBox(nil, i18n.T(cI18n.UacMsgBoxTitle), fmt.Sprintf(i18n.T(cI18n.UacMsgBoxProtocolInstallConfigFailedConfigMsg), err.Error()), walk.MsgBoxIconError)
			}
			os.Exit(1)
		}
	}
}

func installConfig(clashProtocol *ClashProtocol) (alerted bool, err error) {
	configPath := path.Join(constant.ProfileDir, clashProtocol.Name + constant.ConfigSuffix)
	// 检查是否存在，询问是否覆盖
	exists, err := fileUtils.IsExists(configPath)
	if err != nil {
		return false, err
	}
	if exists {
		log.Infoln("[protocol] config is exists: %s", clashProtocol.Name)
		rlt := walk.MsgBox(nil, i18n.T(cI18n.UacMsgBoxTitle), fmt.Sprintf(i18n.T(cI18n.UacMsgBoxProtocolInstallConfigOverwriteMsg), common.Protocol), walk.MsgBoxIconQuestion | walk.MsgBoxYesNo)
		if rlt != win.IDYES {
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

func checkProtocol(clashProtocol *ClashProtocol) (into bool, err error) {
	if len(common.Protocol) > 0 {
		into = true
		Protocol := strings.TrimSpace(common.Protocol)
		log.Infoln(Protocol)
		if protocolRegexp.MatchString(Protocol) {
			protocolQueryString := protocolRegexp.FindStringSubmatch(Protocol)[1]
			log.Infoln(protocolQueryString)
			err = util.UnmarshalByValues(protocolQueryString, clashProtocol)
			if err != nil {
				log.Errorln(err.Error())
				return into, err
			}
			clashProtocol.Url, err = url.PathUnescape(clashProtocol.Url)
			log.Infoln("[protocol] clashProtocol: %v", clashProtocol)
		}
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
