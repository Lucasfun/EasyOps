package unit

import (
	"fmt"
	"math/rand"
	"time"
)

type Unit struct {
	Name    string   `yaml:"name"`
	Install string   `yaml:"install"` //安装脚本，testDemo
	Next    []string `yaml:"next"`    //依赖this unit的其他unit[]
}

func (u *Unit) InstallFunc(callback func()) bool {
	fmt.Println(u.Name + " start installing...")
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second) //模拟安装func
	callback()
	return true
}
func (u *Unit) GetName() string {
	return u.Name
}
func (u *Unit) GetNext() []string {
	return u.Next
}
