package unitsInstall_1_0

import (
	"fmt"
	"testing"
)

func TestGetUnitsConfig(t *testing.T) {
	uc := GetUnitsConfig()
	fmt.Println(uc.GetInDegree())
	fmt.Println(uc.GetOutDegree())
}
