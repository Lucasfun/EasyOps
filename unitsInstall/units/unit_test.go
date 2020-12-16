package units //TODO 接口根目录下 interface.go集中，instance package下各种实现

import (
	"fmt"
	"testing"
)

func TestUnit_InstallFunc(t *testing.T) {

	unitDemo := &Unit{
		Name:    "Demo",
		Install: "DemoInstalling",
		Next:    []string{"SecDemo"},
	}
	unitDemo.InstallFunc(func() {
	})
	fmt.Printf("Name = %v,Next[] = %v", unitDemo.GetName(), unitDemo.GetNext())
}
