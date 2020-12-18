package unit

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
		//do nothing
		return
	})
	fmt.Printf("Name = %v,Next[] = %v", unitDemo.GetName(), unitDemo.GetNext())
}
