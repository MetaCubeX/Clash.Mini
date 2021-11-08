package second

import (
	"fmt"
	"golang.org/x/text/language"

	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/first"

	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/static"
	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"

	. "github.com/JyCyunMe/go-i18n/i18n"
)

func init() {
	fmt.Println("[bridge] second")

	// 初始化语言
	initI18n()
}


// initI18n 初始化语言
func initI18n() {
	SetDefaultLang(English)
	langName := config.GetOrDefault("lang", DefaultLang.Name).(string)
	var lang *Lang
	tag, err := language.Parse(langName)
	if err != nil {
		lang = DefaultLang
		log.Errorln("[i18n] language \"%s\" is invalid, will use default: %s (%s)",
			langName, DefaultLang.Name, DefaultLang.Tag.String())
		config.Set("lang", DefaultLang.Tag.String())
	} else {
		lang = &Lang{Tag: tag}
	}

	packListFunc := func(options ...Option) ([]*Lang, error) {
		embedLangFiles, err := static.LoadEmbedLanguages(true)
		if err != nil {
			return nil, err
		}
		var preLanguages []string
		languageMap := make(map[string]*Lang)
		for _, embedLanguage := range embedLangFiles {
			embedLanguage := *embedLanguage
			data := embedLanguage.Sys().(*[]byte)
			embedLang := ReadLangFromBytes(data, embedLanguage.Name())
			if embedLang == nil {
				continue
			}
			embedLang.Data = data
			log.Infoln("[i18n] Found embed language: %s", embedLang.FullName())
			tagName := embedLang.Tag.String()
			preLanguages = append(preLanguages, tagName)
			languageMap[tagName] = embedLang
		}
		log.Infoln("[i18n] Found %d embed language(s)", len(preLanguages))
		externalLanguages, err := PackageListByPatternFunc(NewOptionWithData(PackagePattern,  commonUtils.GetExecutablePath("lang", "*.lang")))
		if err != nil {
			return nil, err
		}
		// 是否存在覆写解锁文件
		// 存在时才允许使用外置语言包覆盖内嵌语言包
		var overrideLanguageUnlock bool
		overrideLanguageUnlock, err = fileUtils.IsExists(commonUtils.GetExecutablePath("lang", ".unlock"))
		if err != nil {
			return nil, err
		}
		if overrideLanguageUnlock {
			log.Warnln("[i18n] external language override permission is unlocked")
		}
		for _, externalLanguage := range externalLanguages {
			tagName := externalLanguage.Tag.String()
			langName := externalLanguage.FullName()
			_, exists := languageMap[tagName]
			if exists {
				if overrideLanguageUnlock {
					log.Warnln("[i18n] found external language conflicts with embed, overwritten: %s", langName)
					languageMap[tagName] = externalLanguage
				} else {
					log.Warnln("[i18n] found external language conflicts with embed, skipped: %s", langName)
				}
			} else {
				preLanguages = append(preLanguages, tagName)
				languageMap[tagName] = externalLanguage
			}
		}
		var languages []*Lang
		for _, tagName := range preLanguages {
			languages = append(languages, languageMap[tagName])
		}
		log.Infoln("[i18n] Found %d embed and external language(s)", len(languages))
		return languages, err
	}
	//packListFunc := func(options ...Option) ([]*Lang, error) {
	//	return PackageListByPatternFunc(NewOptionWithData(PackagePattern, util.GetExecutablePath("lang", "*.lang")))
	//}
	//InitI18nWithLogFunc(lang, log.Infoln, log.Errorln)
	err = InitI18nWithAllFunc(lang, log.Infoln, log.Errorln, nil, packListFunc)
	if err != nil {
		panic(err)
	}
}
