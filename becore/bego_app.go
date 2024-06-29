package becore

import "becore/bebase"

type IBegoApp interface {
	bebase.IConfigs
	bebase.IBootstrap
	bebase.ISettingAndBackup
	bebase.IAdmin
}
