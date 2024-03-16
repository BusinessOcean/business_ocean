package becore

import "becore/bebase"

type BeCore interface {
	bebase.ConfigsFacade
	bebase.BootstrapFacade
	bebase.SettingAndBackupFacade
	bebase.AdminFacade
}

// BeCoreFacade: BeCoreFacade

// type BeRandom struct {
// }

// // New BeRandom
// func NewBeRandom() *BeRandom {
// 	return &BeRandom{}
// }
