package unitsInstall_1_0

import (
	"github.com/Lucasfun/EasyOps/unitsInstall_1.0/activeObject"
)

func UnitsInstall() {
	uc := GetUnitsConfig()
	s := activeObject.InitService(uc.inDegree, uc.outDegree)
	s.Wait()
}
