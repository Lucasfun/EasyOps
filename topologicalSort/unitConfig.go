package topologicalSort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"
)

//从相关的yaml文件中获取Unit的安装依赖信息

var configFile []byte

//----UnitType附属操作----
type UnitType struct {
	Name    string   `yaml:"name"`
	Install string   `yaml:"install"`
	Next    []string `yaml:"next"`
}

type Install interface { //TODO 安装接口 怎么定义？
	UnitInstall() error
}

func (u *UnitType) UnitInstall(info *InfoMap, q *Queue) error {
	//fmt.Println(u.Name + "安装开始 ： " + u.InstallFunc)
	fmt.Print(u.Name)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	reduction(u, info, q)
	fmt.Print(u.Name)
	//fmt.Println(u.Name + "安装结束！")
	return nil
}

//----UnitConfig附属操作----
type UnitConfig struct {
	UnitConfig []*UnitType `yaml:"unit"`
}

//初始化生成包含相应.yml文件信息的 struct
func (unitConfig *UnitConfig) InitUnitConfig() {
	var err error
	//实际文件位置与形式需输入或定义
	configFile, err = ioutil.ReadFile("./topologicalSort/unit_dependence.yml")
	if err != nil {
		log.Fatal("Get unit_dependence error!", err)
	}
	err = yaml.Unmarshal(configFile, unitConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
}

//利用 struct 生成想要的 map
type InfoMap struct {
	Mux    sync.Mutex        //互斥锁，协程并发中保证数据安全
	Degree map[*UnitType]int //记录 unit.next.len -- 出度数；int -- 入度数
}

func GetInfoMap(u *UnitConfig) *InfoMap {
	infoMap := InfoMap{
		Mux:    sync.Mutex{},
		Degree: make(map[*UnitType]int, 0),
	}

	for _, unit := range u.UnitConfig { //初始化unit 入度都为0
		infoMap.Degree[unit] = 0
		for _, unit2 := range u.UnitConfig { //TODO 优化
			for _, name := range unit2.Next {
				if name == unit.Name {
					infoMap.Degree[unit]++
				}
			}
		}
	}
	return &infoMap
}
