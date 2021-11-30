module github.com/Clash-Mini/Clash.Mini

go 1.17

require (
	github.com/Dreamacro/clash v1.8.0
	github.com/JyCyunMe/go-i18n v0.0.2
	github.com/MakeNowJust/hotkey v0.0.0-20200628032113-41fa0caa507a
	github.com/beevik/etree v1.1.0
	github.com/bugsnag/bugsnag-go/v2 v2.1.2
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/fsnotify/fsnotify v1.5.1
	github.com/getlantern/systray v1.1.0
	github.com/go-toast/toast v0.0.0-20190211030409-01e6764cf0a4
	github.com/imdario/mergo v0.3.12
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lxn/walk v0.0.0-20210112085537-c389da54e794
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e
	github.com/mitchellh/mapstructure v1.4.2
	github.com/robfig/cron/v3 v3.0.1
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spf13/viper v1.9.0
	github.com/zserge/lorca v0.1.10
	golang.org/x/sys v0.0.0-20211107104306-e0b2ad06fe42
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Dreamacro/go-shadowsocks2 v0.1.7 // indirect
	github.com/bugsnag/panicwrap v1.3.4 // indirect
	github.com/getlantern/context v0.0.0-20190109183933-c447772a6520 // indirect
	github.com/getlantern/errors v1.0.1 // indirect
	github.com/getlantern/golog v0.0.0-20201105130739-9586b8bde3a9 // indirect
	github.com/getlantern/hex v0.0.0-20190417191902-c6586a6fe0b7 // indirect
	github.com/getlantern/hidden v0.0.0-20201229170000-e66e7f878730 // indirect
	github.com/getlantern/ops v0.0.0-20200403153110-8476b16edcd6 // indirect
	github.com/go-chi/chi/v5 v5.0.5 // indirect
	github.com/go-chi/cors v1.2.0 // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gofrs/uuid v4.1.0+incompatible // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/insomniacslk/dhcp v0.0.0-20211026125128-ad197bcd36fd // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kr328/tun2socket v0.0.0-20210412191540-3d56c47e2d99 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.1.2 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/oschwald/geoip2-golang v1.5.0 // indirect
	github.com/oschwald/maxminddb-golang v1.8.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/u-root/uio v0.0.0-20210528114334-82958018845c // indirect
	github.com/xtls/go v0.0.0-20201118062508-3632bf3b7499 // indirect
	github.com/yaling888/go-lwip v0.0.0-20211103185822-c9d650538091 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20211105192438-b53810dc28af // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.zx2c4.com/wireguard/windows v0.5.1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gvisor.dev/gvisor v0.0.0-20211104052249-2de3450f76d6 // indirect
)

replace (
	github.com/Dreamacro/clash v1.8.0 => github.com/Clash-Mini/clash v1.6.1-0.20211130100019-c65835d9e4c7
	//github.com/Shopify/logrus-bugsnag/v2 v2.0.0 => github.com/JyCyunMe/Clash.Mini-Vendor/Shopify/logrus-bugsnag v0.0.0-20210610225813-69b2b3cedbfe
	github.com/getlantern/systray v1.1.0 => github.com/JyCyunMe/Clash.Mini-Vendor/getlantern/systray v0.0.0-20211112095307-ac090dd0663d
)
