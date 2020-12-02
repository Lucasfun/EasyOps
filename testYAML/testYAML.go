package testYAML

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configFile []byte

type VertexType struct {
	Name string `yaml:"name"`
	Next []string `yaml:"next"`
}

type GraphType struct {
	Vertex []*VertexType `yaml:"vertex"`
}

type GraphConfig struct {
	Graph *GraphType `yaml:"graph"`
}

//初始化生成包含相应.yml文件信息的 struct
func (graph *GraphConfig) Init() {
	var err error
	//实际文件位置与形式需输入或定义
	configFile,err = ioutil.ReadFile("./testYAML/unit_dependence.yml")
	//fmt.Println(configFile)
	if err != nil{
		log.Fatal("unit_dependence file get error! %v",err)
	}
	//
	err = yaml.Unmarshal(configFile,graph)
	if err != nil {
		log.Fatal(err.Error())
	}

}

//利用 struct 生成想要的 map
func (graph *GraphConfig) New_outDegree() (outDegree *map[string][]string) {
	out :=make(map[string][]string)
	vertex := graph.Graph.Vertex
	for _,demo := range vertex{
		out[demo.Name] = demo.Next
	}
	return &out
}