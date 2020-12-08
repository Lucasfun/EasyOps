package unit_install

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var fileBytes []byte

type UnitConfig struct {
	Config 	[]*Unit `yaml:"units"`
	Degree  map[*Unit]int
}

//初始化生成包含相应.yml文件信息的 struct
func GetUnitConfig() *UnitConfig {
	uc := &UnitConfig{
		Config: make([]*Unit,0),
		Degree: make(map[*Unit]int,0),
	}
	//----------- 读取文件 生成UnitConfig -----------
	var err error
	fileBytes,err = ioutil.ReadFile("./unit_install/unitConfig.yml")	//实际文件位置与形式需输入或定义
	if err != nil{
		log.Fatal("Get unit_dependence error!",err)
	}
	err = yaml.Unmarshal(fileBytes, uc)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ---------- 利用UnitConfig 生成InfoMap ---------
	for _, unit := range uc.Config {
		uc.Degree[unit] = 0          //初始化unit 入度都为0
		for _, unit2 := range uc.Config { //TODO 优化
			for _,name := range unit2.Next{
				if name == unit.Name{
					uc.Degree[unit] ++
				}
			}
		}
	}
	return uc
}



