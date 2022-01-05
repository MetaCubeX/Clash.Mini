package fourth

import (
	"fmt"
	"github.com/MetaCubeX/Clash.Mini/static"
	"io"
	"io/ioutil"
	"os"
	"path"

	_ "github.com/MetaCubeX/Clash.Mini/app/bridge/start/third"

	"github.com/MetaCubeX/Clash.Mini/common"
	. "github.com/MetaCubeX/Clash.Mini/constant"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/notify"
)

func init() {
	log.Infoln("[bridge] Step Fourth: Checking...")

	//common.GetVarFlags()
	//common.InitVariablesAfterGetVarFlags()
	//common.InitFunctionsAfterGetVarFlags()

	if _, err := os.Stat(ProfileDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(ProfileDir, 0666); err != nil {
				errMsg := fmt.Sprintf("cannot create config dir: %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				common.DisabledCore = true
				return
			}
			if err = ioutil.WriteFile(path.Join(ProfileDir, ConfigFile), static.ExampleConfig, 0644); err != nil {
				err = fmt.Errorf("write default core config file error: %s", err.Error())
			}
		}
	}
	if _, err := os.Stat(CacheDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(CacheDir, 0666); err != nil {
				errMsg := fmt.Sprintf("cannot create cache dir: %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				common.DisabledCore = true
				return
			}
		}
	}
	if _, err := os.Stat(MixinDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(MixinDir, 0666); err != nil {
				errMsg := fmt.Sprintf("cannot create mixin dir: %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				common.DisabledCore = true
				return
			}
			mixinFiles, err := static.Mixin.ReadDir("mixin")
			if err != nil {
				errMsg := fmt.Sprintf("cannot find minxin.yaml(s): %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				return
			}
			for _, mixinFile := range mixinFiles {
				in, _ := static.Mixin.Open(path.Join("mixin", mixinFile.Name()))
				out, _ := os.Create(path.Join(MixinDir, mixinFile.Name()))
				if _, err := io.Copy(out, in); err != nil {
					errMsg := fmt.Sprintf("cannot create minxin.yaml(s): %v", err)
					log.Errorln(errMsg)
					notify.PushError("", errMsg)
					return
				}
				if err = out.Close(); err != nil {
					return
				}
				if err = in.Close(); err != nil {
					return
				}
			}
		}
	}
}
