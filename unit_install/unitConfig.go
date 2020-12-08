package unit_install

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var fileBytes []byte

type UnitConfig struct {
	Config 		[]*Unit `yaml:"units"`
	inDegree 	map[string]int
	outDegree 	map[string]*Unit
}

//初始化生成包含相应.yml文件信息的 struct
func GetUnitConfig() *UnitConfig {
	uc := &UnitConfig{
		Config: 	make([]*Unit,0),
		inDegree: 	make(map[string]int,0),
		outDegree: 	make(map[string]*Unit,0),
	}
	//----------- 读取文件 生成UnitConfig -----------
	var err error
	fileBytes,err = ioutil.ReadFile("./unit_install/unitConfig.yml")	//实际文件位置与形式需输入或定义
	if err != nil{
		log.Fatal("读取" + "unit_dependence.yml" + "失败!",err)
	}
	err = yaml.Unmarshal(fileBytes, uc)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ---------- 利用UnitConfig 生成InfoMap ---------
	for _,unit := range uc.Config{
		uc.outDegree[unit.Name] = unit
	}
	for _, unit := range uc.Config {
		uc.inDegree[unit.Name] = 0          //初始化unit 入度都为0
		for _, unit2 := range uc.Config { //TODO 优化
			for _,name := range unit2.Next{
				if name == unit.Name{
					uc.inDegree[unit.Name] ++
				}
			}
		}
	}
	return uc
}



