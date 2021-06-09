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

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"

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
		Title:    util.GetSubTitle("添加配置"),
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
						Text: "订阅名称:",
					},
					LineEdit{
						AssignTo: &oUrlName,
					},
					Label{
						Text: "订阅链接:",
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
						Text: "添加",
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
										errMsg = "无法访问到订阅链接！"
									} else if err == http.ErrNoLocation || err == http.ErrMissingFile ||
										(rsp != nil && rsp.StatusCode == http.StatusNotFound) {
										errMsg = "请检查订阅链接是否正确且未失效！"
									} else {
										errMsg = "下载失败！"
									}
									walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle, errMsg, walk.MsgBoxIconError)
								}
								if rsp != nil && rsp.StatusCode == 200 {
									Reg, err := regexp.MatchString(`proxy-groups`, rspBody)
									if err != nil || !Reg {
										log.Errorln("配置内容有误: %v", err)
										walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
											"检测为非Clash配置，添加配置失败！", walk.MsgBoxIconError)
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
									walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
										"添加配置成功！", walk.MsgBoxIconInformation)
									AddMenuConfig.Close()
								} else {
									walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
										"请检查订阅链接是否正确！", walk.MsgBoxIconError)
								}
							} else {
								walk.MsgBox(AddMenuConfig, constant.UIConfigMsgTitle,
									"请输入并检查订阅名称和链接！", walk.MsgBoxIconError)
							}
						},
					},
					PushButton{
						Text: "取消",
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
