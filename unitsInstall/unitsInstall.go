package unitsInstall

import (
	"github.com/Lucasfun/EasyOps/unitsInstall/activeObject"
)

func UnitsInstall() {
	uc := GetUnitsConfig()
	s := activeObject.InitService(uc)
	s.Wait()
}
