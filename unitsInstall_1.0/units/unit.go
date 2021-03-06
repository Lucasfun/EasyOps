package units

import (
	"fmt"
	"math/rand"
	"time"
)

type Unit struct {
	Name    string   `yaml:"name"`
	Install string   `yaml:"install"`
	Next    []string `yaml:"next"`
}

func (u *Unit) InstallFunc(callback func()) bool {
	fmt.Println(u.Name + " 开始安装...")
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	callback() //报告给Service本unit安装完成
	return true
}

func (u Unit) GetName() string {
	return u.Name
}
func (u Unit) GetNext() []string {
	return u.Next
}
