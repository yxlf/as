package global

import (
	"as/pkg/logger"
	"as/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
)
var (
	Logger *logger.Logger
)
