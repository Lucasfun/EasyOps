package unit_install

//定义组件Unit的结构、安装方法
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

type Install interface {
	UnitInstall() (bool, string)
}

type Scheduler interface {
	Report2(*Unit) bool
}

func (u *Unit) UnitInstall2(s Scheduler) { //TODO 实际安装出错形式，判断返回
	fmt.Println(u.Name)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	//fmt.Println(u.Name + "finish install.")
	s.Report2(u)
}
