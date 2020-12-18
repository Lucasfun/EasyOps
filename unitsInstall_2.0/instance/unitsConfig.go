package instance

import (
	"github.com/Lucasfun/EasyOps/unitsInstall_2.0/instance/unit"
	_interface "github.com/Lucasfun/EasyOps/unitsInstall_2.0/interface"
	"gopkg.in/yaml.v2"
)

type UnitsConfig struct {
	Config []*unit.Unit `yaml:"units"`
}

func GetUnitsConfig(fileBytes []byte) (*UnitsConfig, error) {
	uc := &UnitsConfig{
		Config: make([]*unit.Unit, 0),
	}
	if err := yaml.Unmarshal(fileBytes, uc); err != nil {
		return nil, err
	}
	return uc, nil
}

func (uc *UnitsConfig) GetInDegree() map[string]int {
	inDegree := make(map[string]int, 0)
	for _, unit := range uc.Config {
		inDegree[unit.GetName()] = 0 //初始化unit 入度都为0
		for _, unit2 := range uc.Config {
			for _, name := range unit2.GetNext() {
				if name == unit.GetName() {
					inDegree[unit.GetName()]++
				}
			}
		}
	}
	return inDegree
}

func (uc *UnitsConfig) GetOutDegree() map[string]_interface.UnitInterface {
	outDegree := make(map[string]_interface.UnitInterface, 0)
	for _, unit := range uc.Config {
		outDegree[unit.GetName()] = unit
	}
	return outDegree
}
