package i18n

const (
	Enabled  = "ENABLED"
	Disabled = "DISABLED"
	Running  = "RUNNING"
	Stopped  = "STOPPED"

	TrayMenuCoreEnabled  = "TRAY_MENU.CORE.ENABLED"
	TrayMenuCoreDisabled = "TRAY_MENU.CORE.DISABLED"
	TrayMenuCoreRunning  = "TRAY_MENU.CORE.RUNNING"
	TrayMenuCoreStopped  = "TRAY_MENU.CORE.STOPPED"

	// 核心开关
	TrayMenuCoreTurnOn  = "TRAY_MENU.CORE.TURN_ON"
	TrayMenuCoreTurnOff = "TRAY_MENU.CORE.TURN_OFF"

	// 代理模式
	TrayMenuGlobalProxy = "TRAY_MENU.GLOBAL_PROXY"
	TrayMenuRuleProxy   = "TRAY_MENU.RULE_PROXY"
	TrayMenuDirectProxy = "TRAY_MENU.DIRECT_PROXY"

	// 其他菜单
	TrayMenuSystemProxy         = "TRAY_MENU.SYSTEM_PROXY"
	TrayMenuSwitchProxy         = "TRAY_MENU.SWITCH_PROXY"
	TrayMenuSwitchProfile       = "TRAY_MENU.SWITCH_PROFILE"
	TrayMenuSwitchProfileUpdate = "TRAY_MENU.SWITCH_PROFILE_UPDATE"
	TrayMenuDashboard           = "TRAY_MENU.DASHBOARD"
	TrayMenuConfigManagement    = "TRAY_MENU.CONFIG_MANAGEMENT"
	TrayMenuShowLog             = "TRAY_MENU.SHOW_LOG"

	// 延迟测速
	TrayMenuPingTest            = "TRAY_MENU.PING_TEST"
	TrayMenuPingTestDoNow       = "TRAY_MENU.PING_TEST.DO_NOW"
	TrayMenuPingTestFastProxy   = "TRAY_MENU.PING_TEST.FAST_PROXY"
	TrayMenuPingTestLowestDelay = "TRAY_MENU.PING_TEST.LOWEST_DELAY"
	TrayMenuPingTestLastUpdate  = "TRAY_MENU.PING_TEST.LAST_UPDATE"

	// 其他设置
	TrayMenuOtherSettings                  = "TRAY_MENU.OTHER_SETTINGS"
	TrayMenuOtherSettingsMixin             = "TRAY_MENU.OTHER_SETTINGS.MIXIN"
	TrayMenuOtherSettingsMixinDir          = "TRAY_MENU.OTHER_SETTINGS.MIXIN_DIR"
	TrayMenuOtherSettingsMixinGeneral      = "TRAY_MENU.OTHER_SETTINGS.MIXIN_GENERAL"
	TrayMenuOtherSettingsMixinTun          = "TRAY_MENU.OTHER_SETTINGS.MIXIN_TUN"
	TrayMenuOtherSettingsMixinDns          = "TRAY_MENU.OTHER_SETTINGS.MIXIN_DNS"
	TrayMenuOtherSettingsSwitchLanguage    = "TRAY_MENU.OTHER_SETTINGS.SWITCH_LANGUAGE"
	TrayMenuOtherSettingsSystemAutorun     = "TRAY_MENU.OTHER_SETTINGS.SYSTEM_AUTORUN"
	TrayMenuOtherSettingsSystemAutoProxy   = "TRAY_MENU.OTHER_SETTINGS.SYSTEM_AUTO_PROXY"
	TrayMenuOtherSettingsCronUpdateConfigs = "TRAY_MENU.OTHER_SETTINGS.CRON_UPDATE_CONFIGS"
	TrayMenuOtherSettingsHotkey            = "TRAY_MENU.OTHER_SETTINGS.HOTKEY"
	TrayMenuOtherSettingsRegisterProtocol  = "TRAY_MENU.OTHER_SETTINGS.REGISTER_PROTOCOL"
	TrayMenuOtherSettingsUwpLoopback       = "TRAY_MENU.OTHER_SETTINGS.UWP_LOOPBACK"
	TrayMenuOtherSettingsSetMMDB           = "TRAY_MENU.OTHER_SETTINGS.SET_MMDB"
	TrayMenuOtherSettingsSetMMDBMaxmind    = "TRAY_MENU.OTHER_SETTINGS.SET_MMDB.MAXMIND"
	TrayMenuOtherSettingsSetMMDBHackl0Us   = "TRAY_MENU.OTHER_SETTINGS.SET_MMDB.HACKL0US"

	TrayMenuQuit = "TRAY_MENU.QUIT"

	ProxyTestTimeout = "PROXY.TEST.TIMEOUT"

	MenuConfigWindowEnableConfig             = "MENU_CONFIG.WINDOW.ENABLE_CONFIG"
	MenuConfigWindowEditConfig               = "MENU_CONFIG.WINDOW.EDIT_CONFIG"
	MenuConfigWindowUpdateConfig             = "MENU_CONFIG.WINDOW.UPDATE_CONFIG"
	MenuConfigWindowUpdateAll                = "MENU_CONFIG.WINDOW.UPDATE_ALL"
	MenuConfigWindowUpdating                 = "MENU_CONFIG.WINDOW.UPDATING"
	MenuConfigWindowUpdateFinished           = "MENU_CONFIG.WINDOW.UPDATE_FINISHED"
	MenuConfigWindowDeleteConfig             = "MENU_CONFIG.WINDOW.DELETE_CONFIG"
	MenuConfigWindowCurrentConfig            = "MENU_CONFIG.WINDOW.CURRENT_CONFIG"
	MenuConfigWindowConvertSubscription      = "MENU_CONFIG.WINDOW.CONVERT_SUBSCRIPTION"
	MenuConfigWindowOpenConfigDir            = "MENU_CONFIG.WINDOW.OPEN_CONFIG_DIR"
	MenuConfigWindowCloseWindow              = "MENU_CONFIG.WINDOW.CLOSE_WINDOW"
	MenuConfigWindowAddConfig                = "MENU_CONFIG.WINDOW.ADD_CONFIG"
	MenuConfigWindowAddConfigBottomAdd       = "MENU_CONFIG.WINDOW.ADD_CONFIG.ADD_BOTTOM"
	MenuConfigWindowAddConfigBottomCancel    = "MENU_CONFIG.WINDOW.ADD_CONFIG.CANCEL_BOTTOM"
	MenuConfigWindowAddConfigFail            = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.FAIL"
	MenuConfigWindowAddConfigUrlTimeout      = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.URL.TIMEOUT"
	MenuConfigWindowAddConfigUrlCodeFail     = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.URL.CODE_FAIL"
	MenuConfigWindowAddConfigUrlDownloadFail = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.URL.DOWNLOAD_FAIL"
	MenuConfigWindowAddConfigUrlFail         = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.URL.FAIL"
	MenuConfigWindowAddConfigUrlNotClash     = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.URL.NOT_CLASH"
	MenuConfigWindowAddConfigUrlSuccess      = "MENU_CONFIG.WINDOW.ADD_CONFIG.MESSAGEBOX.URL.SUCCESS"

	AddConfigWindowTitle = "ADD_CONFIG.WINDOW.TITLE"

	EditConfigWindowTitle                 = "EDIT_CONFIG.WINDOW.TITLE"
	EditConfigWindowSubscriptionName      = "EDIT_CONFIG.WINDOW.SUBSCRIPTION_NAME"
	EditConfigWindowSubscriptionUrl       = "EDIT_CONFIG.WINDOW.SUBSCRIPTION_URL"
	EditConfigMessageEditConfigConfirmMsg = "EDIT_CONFIG.MESSAGE.EDIT_CONFIG.CONFIRM_MESSAGE"
	EditConfigMessageEditConfigSuccess    = "EDIT_CONFIG.MESSAGE.EDIT_CONFIG.SUCCESS"
	EditConfigMessageEditConfigFailure    = "EDIT_CONFIG.MESSAGE.EDIT_CONFIG.FAILURE"
	EditConfigMessageEditConfigNothing    = "EDIT_CONFIG.MESSAGE.EDIT_CONFIG.NOTHING"

	MsgBoxTitleTips                         = "MESSAGEBOX.TITLE.TIPS"
	MenuConfigMessageEnableConfigSuccess    = "MENU_CONFIG.MESSAGE.ENABLE_CONFIG.SUCCESS"
	MenuConfigMessageEnableConfigFailure    = "MENU_CONFIG.MESSAGE.ENABLE_CONFIG.FAILURE"
	MenuConfigMessageUpdateConfigSuccess    = "MENU_CONFIG.MESSAGE.UPDATE_CONFIG.SUCCESS"
	MenuConfigMessageUpdateConfigFailure    = "MENU_CONFIG.MESSAGE.UPDATE_CONFIG.FAILURE"
	MenuConfigMessageDeleteConfigConfirmMsg = "MENU_CONFIG.MESSAGE.DELETE_CONFIG.CONFIRM_MESSAGE"
	MenuConfigMessageDeleteConfigSuccess    = "MENU_CONFIG.MESSAGE.DELETE_CONFIG.SUCCESS"
	MenuConfigMessageDeleteConfigFailure    = "MENU_CONFIG.MESSAGE.DELETE_CONFIG.FAILURE"
	MenuConfigWindowConfigName              = "MENU_CONFIG.WINDOW.CONFIG_NAME"
	MenuConfigWindowFileSize                = "MENU_CONFIG.WINDOW.FILE_SIZE"
	MenuConfigWindowUpdateDatetime          = "MENU_CONFIG.WINDOW.UPDATE_DATETIME"
	MenuConfigWindowSubscriptionUrl         = "MENU_CONFIG.WINDOW.SUBSCRIPTION_URL"

	MenuConfigWindowConfigManagement = "MENU_CONFIG.WINDOW.CONFIG_MANAGEMENT"

	MenuConfigCronUpdateSuccessful = "MENU_CONFIG.CRON.UPDATE.SUCCESSFUL"
	MenuConfigCronUpdateFailed     = "MENU_CONFIG.CRON.UPDATE.FAILED"

	UacMsgBoxTitle                                    = "UAC.MESSAGEBOX.TITLE"
	UacMsgBoxProtocolInstallConfigConfirmMsg          = "UAC.MESSAGEBOX.PROTOCOL.INSTALL_CONFIG.CONFIRM.MESSAGE"
	UacMsgBoxProtocolInstallConfigOverwriteMsg        = "UAC.MESSAGEBOX.PROTOCOL.INSTALL_CONFIG.OVERWRITE.MESSAGE"
	UacMsgBoxProtocolInstallConfigSuccessfulConfigMsg = "UAC.MESSAGEBOX.PROTOCOL.INSTALL_CONFIG.SUCCESSFUL.MESSAGE"
	UacMsgBoxProtocolInstallConfigFailedConfigMsg     = "UAC.MESSAGEBOX.PROTOCOL.INSTALL_CONFIG.FAILED.MESSAGE"

	UacMsgBoxTunFailedMsg = "UAC.MESSAGEBOX.TUN.MESSAGE"

	ButtonSubmit = "BUTTON.SUBMIT"
	ButtonCancel = "BUTTON.CANCEL"

	UtilDatetimeAgo               = "UTIL.DATETIME.AGO"
	UtilDatetimeLater             = "UTIL.DATETIME.LATER"
	UtilDatetimeHalf              = "UTIL.DATETIME.HALF"
	UtilDatetimeMoments           = "UTIL.DATETIME.MOMENTS"
	UtilDatetimeShortMilliSeconds = "UTIL.DATETIME.SHORT.MILLI_SECONDS"
	UtilDatetimeShortSeconds      = "UTIL.DATETIME.SHORT.SECONDS"
	UtilDatetimeShortMinutes      = "UTIL.DATETIME.SHORT.MINUTES"
	UtilDatetimeShortHours        = "UTIL.DATETIME.SHORT.HOURS"
	UtilDatetimeMilliSeconds      = "UTIL.DATETIME.MILLI_SECONDS"
	UtilDatetimeSeconds           = "UTIL.DATETIME.SECONDS"
	UtilDatetimeMinutes           = "UTIL.DATETIME.MINUTES"
	UtilDatetimeHours             = "UTIL.DATETIME.HOURS"
	UtilDatetimeDays              = "UTIL.DATETIME.DAYS"
	UtilDatetimeWeeks             = "UTIL.DATETIME.WEEKS"
	UtilDatetimeMonths            = "UTIL.DATETIME.MONTHS"

	NotifyMessageMixinGeneralOn  = "NOTIFY.MESSAGE.MIXIN.GENERAL.ON"
	NotifyMessageMixinGeneralOff = "NOTIFY.MESSAGE.MIXIN.GENERAL.OFF"
	NotifyMessageMixinTunOn      = "NOTIFY.MESSAGE.MIXIN.TUN.ON"
	NotifyMessageMixinTunOff     = "NOTIFY.MESSAGE.MIXIN.TUN.OFF"
	NotifyMessageMixinDnsOn      = "NOTIFY.MESSAGE.MIXIN.DNS.ON"
	NotifyMessageMixinDnsOff     = "NOTIFY.MESSAGE.MIXIN.DNS.OFF"

	NotifyMessageSysOn          = "NOTIFY.MESSAGE.SYS.ON"
	NotifyMessageSysOff         = "NOTIFY.MESSAGE.SYS.OFF"
	NotifyMessageModeDirect     = "NOTIFY.MESSAGE.MODE.DIRECT"
	NotifyMessageModeRULE       = "NOTIFY.MESSAGE.MODE.RULE"
	NotifyMessageModeGLOBAL     = "NOTIFY.MESSAGE.MODE.GLOBAL"
	NotifyMessageStartupOn      = "NOTIFY.MESSAGE.STARTUP.ON"
	NotifyMessageStartupOff     = "NOTIFY.MESSAGE.STARTUP.OFF"
	NotifyMessageAutosysOn      = "NOTIFY.MESSAGE.AUTOSYS.ON"
	NotifyMessageAutosysOff     = "NOTIFY.MESSAGE.AUTOSYS.OFF"
	NotifyMessageMmdbMax        = "NOTIFY.MESSAGE.MMDB.MAX"
	NotifyMessageMmdbLite       = "NOTIFY.MESSAGE.MMDB.LITE"
	NotifyMessageCronOn         = "NOTIFY.MESSAGE.CRON.ON"
	NotifyMessageCronOff        = "NOTIFY.MESSAGE.CRON.OFF"
	NotifyMessageTitle          = "NOTIFY.MESSAGE.TITLE"
	NotifyMessageErrorTitle     = "NOTIFY.MESSAGE.ERROR_TITLE"
	NotifyMessageFlowTitle      = "NOTIFY.MESSAGE.FLOW.TITLE"
	NotifyMessageFlowUsed       = "NOTIFY.MESSAGE.FLOW.USED"
	NotifyMessageFlowUnused     = "NOTIFY.MESSAGE.FLOW.UNUSED"
	NotifyMessageFlowExpiration = "NOTIFY.MESSAGE.FLOW.EXPIRATION"

	NotifyMessageUpdateFinish = "NOTIFY.MESSAGE.UPDATE.FINISH"

	NotifyMessageCronTitle      = "NOTIFY.MESSAGE.CRON.TITLE"
	NotifyMessageCronFinish     = "NOTIFY.MESSAGE.CRON.FINISH"
	NotifyMessageCronNumSuccess = "NOTIFY.MESSAGE.CRON.NUM.SUCCESS"
	NotifyMessageCronNumFail    = "NOTIFY.MESSAGE.CRON.NUM.FAIL"
	NotifyMessageCronFinishAll  = "NOTIFY.MESSAGE.CRON.FINISH.ALL"
	TaskSchedulerDescription    = "TASKSCHEDULER.DESCRIPTION"
)
