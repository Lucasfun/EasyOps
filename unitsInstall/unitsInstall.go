package unitsInstall

import (
	"github.com/Lucasfun/EasyOps/unitsInstall/activeObject"
)

func UnitsInstall() {
	uc := GetUnitsConfig()
	s := activeObject.InitService(uc.inDegree, uc.outDegree)
	s.Wait()
}
