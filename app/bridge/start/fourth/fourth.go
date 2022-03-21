package fourth

import (
	"fmt"
	"github.com/MetaCubeX/Clash.Mini/static"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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

	if _, err := os.Stat(AumIdDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(ProfileDir, 0666); err != nil {
				_ = fmt.Sprintf("cannot create config dir: %v", err)
				return
			}
			MakeLink(os.Args[0], path.Join(ProfileDir, MiniLnk))
		}
	}

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

func MakeLink(src, dst string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		return err
	}
	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", src)
	oleutil.CallMethod(idispatch, "Save")
	return nil
}
