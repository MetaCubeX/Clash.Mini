package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	path "path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func AddConfig() {
	var AddMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	var oUrlName *walk.LineEdit
	err := MainWindow{
		Visible:  false,
		AssignTo: &AddMenuConfig,
		Name:     "AddConfig",
		Title:    util.GetSubTitle(i18n.T(cI18n.MenuConfigWindowAddConfig)),
		Icon:     appIcon,
		Font: Font{
			Family:    "Microsoft YaHei",
			PointSize: 9,
		},
		Layout: VBox{Alignment: AlignHCenterVCenter}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: i18n.T(cI18n.MenuConfigWindowConfigName),
					},
					LineEdit{
						AssignTo: &oUrlName,
					},
					Label{
						Text: i18n.T(cI18n.MenuConfigWindowSubscriptionUrl),
					},
					LineEdit{
						AssignTo: &oUrl,
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: i18n.T(cI18n.MenuConfigWindowAddConfigBottomAdd),
						OnClicked: func() {
							urlMatched, _ := regexp.MatchString("^https?://(\\w.+.)?(\\w.+\\.\\w.+)", oUrl.Text())
							if oUrlName != nil && oUrl != nil && urlMatched {
								client := &http.Client{Timeout: 10 * time.Second}
								req, _ := http.NewRequest(http.MethodGet, oUrl.Text(), nil)
								req.Header.Add("User-Agent", "clash")
								rsp, err := client.Do(req)
								defer rsp.Body.Close()
								var rspBody string
								if rsp != nil {
									rspBody = string(util.IgnoreErrorBytes(ioutil.ReadAll(rsp.Body)))
								}
								if err != nil || (rsp != nil && rsp.StatusCode != http.StatusOK) {
									log.Warnln("AddConfig Do error: %v, request url: %s, response: [%s] %s",
										err, req.URL.String(), rsp.StatusCode, rspBody)
									var errMsg string
									if err == http.ErrHandlerTimeout ||
										(rsp != nil && rsp.StatusCode == http.StatusInternalServerError ||
											rsp.StatusCode == http.StatusServiceUnavailable) {
										errMsg = i18n.T(cI18n.MenuConfigWindowAddConfigUrlTimeout)
									} else if err == http.ErrNoLocation || err == http.ErrMissingFile ||
										(rsp != nil && rsp.StatusCode == http.StatusNotFound) {
										errMsg = i18n.T(cI18n.MenuConfigWindowAddConfigUrlCodeFail)
									} else {
										errMsg = i18n.T(cI18n.MenuConfigWindowAddConfigUrlDownloadFail)
									}
									walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle, errMsg, walk.MsgBoxIconError)
								}
								if rsp != nil && rsp.StatusCode == 200 {
									Reg, err := regexp.MatchString(`proxy-groups`, rspBody)
									if err != nil || !Reg {
										log.Errorln("%v: %v", i18n.T(cI18n.MenuConfigWindowAddConfigUrlNotClash), err)
										walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
											i18n.T(cI18n.MenuConfigWindowAddConfigUrlNotClash), walk.MsgBoxIconError)
										return
									}
									rspBodyReader := ioutil.NopCloser(strings.NewReader(rspBody))
									configDir := path.Join(constant.ConfigDir, oUrlName.Text()+constant.ConfigSuffix)
									f, err := os.Create(configDir)
									if err != nil {
										panic(err)
									}
									_, err = f.WriteString(fmt.Sprintf("# Clash.Mini : %s\n", oUrl.Text()))
									_, err = io.Copy(f, rspBodyReader)
									err = f.Close()
									go common.RefreshProfile()
									walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
										i18n.T(cI18n.MenuConfigWindowAddConfigUrlSuccess), walk.MsgBoxIconInformation)
									AddMenuConfig.Close()
								} else {
									walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
										i18n.T(cI18n.MenuConfigWindowAddConfigUrlFail), walk.MsgBoxIconError)
								}
							} else {
								walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
									i18n.T(cI18n.MenuConfigWindowAddConfigFail), walk.MsgBoxIconError)
							}
						},
					},
					PushButton{
						Text: i18n.T(cI18n.MenuConfigWindowAddConfigBottomCancel),
						OnClicked: func() {
							AddMenuConfig.Close()
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		return
	}
	StyleMenuRun(AddMenuConfig, 420, 120)
}
