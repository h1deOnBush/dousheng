package global

import (
	"github/h1deOnBush/dousheng/pkg/logger"
	"github/h1deOnBush/dousheng/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting    *setting.JWTSettingS
)
