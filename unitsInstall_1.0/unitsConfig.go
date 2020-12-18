package unitsInstall_1_0

import (
	"github.com/Lucasfun/EasyOps/unitsInstall_1.0/activeObject"
	"github.com/Lucasfun/EasyOps/unitsInstall_1.0/units"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var fileBytes []byte

type UnitsConfig struct {
	Config    []*units.Unit                         `yaml:"units"`
	inDegree  map[string]int                        //unit.name: unit.inDegree
	outDegree map[string]activeObject.UnitInterface //unit.name: *unit
}

func (u UnitsConfig) GetInDegree() map[string]int {
	return u.inDegree
}

func (u UnitsConfig) GetOutDegree() map[string]activeObject.UnitInterface {
	return u.outDegree
}

func GetUnitsConfig() *UnitsConfig {
	uc := &UnitsConfig{
		Config:    make([]*units.Unit, 0),
		inDegree:  make(map[string]int, 0),
		outDegree: make(map[string]activeObject.UnitInterface, 0),
	}
	//从YAML文件读取各组件依赖关系
	var err error
	fileBytes, err = ioutil.ReadFile("./unitsInstall_1.0/units/unitsConfig.yml")
	if err != nil {
		log.Fatal("无法成功从"+"unitsConfig.yml"+"中读取相关信息！", err)
	}
	err = yaml.Unmarshal(fileBytes, uc)
	if err != nil {
		log.Fatal(err.Error())
	}
	//初始化UnitsConfig.inDegree outDegree
	for _, unit := range uc.Config {
		uc.outDegree[unit.GetName()] = unit
	}
	for _, unit := range uc.Config {
		uc.inDegree[unit.GetName()] = 0 //初始化unit 入度都为0
		for _, unit2 := range uc.Config {
			for _, name := range unit2.GetNext() {
				if name == unit.GetName() {
					uc.inDegree[unit.GetName()]++
				}
			}
		}
	}
	return uc
}
