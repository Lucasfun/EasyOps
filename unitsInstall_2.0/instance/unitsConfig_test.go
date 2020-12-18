package instance

import (
	"fmt"
	"reflect"
	"testing"
)

var bytes string = `# 组件
units:
  - name: A  # 组件名
    install: A_installDemo  #安装方法
    next: [D,E]  # 依赖A的其他组件
  - name: B
    install: B_installDemo
    next: [E]
  - name: C
    install: C_installDemo
    next: [F]
  - name: D
    install: D_installDemo
    next: [G]
  - name: E
    install: E_installDemo
    next: [G]
  - name: F
    install: F_installDemo
    next: [H]
  - name: G
    install: G_installDemo
    next: [H]
  - name: H
    install: H_installDemo
    next:`

func TestGetUnitsConfig(t *testing.T) {
	dataset := []byte(bytes)
	uc, _ := GetUnitsConfig(dataset)
	for name, unit := range uc.GetOutDegree() {
		fmt.Printf("%v: ", name)
		fmt.Printf("{%v}\n", unit)
	}
	want := map[string]int{"A": 0, "B": 0, "C": 0, "D": 1, "E": 2, "F": 1, "G": 2, "H": 2}
	if !reflect.DeepEqual(uc.GetInDegree(), want) {
		t.Errorf("test.GetInDegree() = %v,want = %v", uc.GetInDegree(), want)
	}
}
