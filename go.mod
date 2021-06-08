module github.com/Clash-Mini/Clash.Mini

go 1.16

require (
	github.com/Dreamacro/clash v1.6.0
	github.com/MakeNowJust/hotkey v0.0.0-20200628032113-41fa0caa507a
	github.com/beevik/etree v1.1.0
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/getlantern/golog v0.0.0-20201105130739-9586b8bde3a9 // indirect
	github.com/getlantern/hidden v0.0.0-20201229170000-e66e7f878730 // indirect
	github.com/getlantern/ops v0.0.0-20200403153110-8476b16edcd6 // indirect
	github.com/getlantern/systray v1.1.0
	github.com/go-toast/toast v0.0.0-20190211030409-01e6764cf0a4
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lxn/walk v0.0.0-20210112085537-c389da54e794
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/robfig/cron/v3 v3.0.1
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/zserge/lorca v0.1.10
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015
	golang.org/x/text v0.3.6
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace (
	github.com/Dreamacro/clash v1.6.0 => github.com/JyCyunMe/Clash.Mini-Vendor/Dreamacro/clash v0.0.0-20210605163326-2639e89c1934
	github.com/getlantern/systray v1.1.0 => github.com/JyCyunMe/Clash.Mini-Vendor/getlantern/systray v0.0.0-20210605164140-1b8bd62f33a7
)
