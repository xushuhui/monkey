package global

import (
	"goal-layout/pkg/logger"
	"goal-layout/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting    *setting.AppSettings

	LogSetting *setting.LogSettings
	Logger     *logger.Logger
)
var (
	AppSignExpiry = "120"
)
