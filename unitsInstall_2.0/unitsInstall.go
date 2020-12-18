package unitsInstall_2_0

import (
	"github.com/Lucasfun/EasyOps/unitsInstall_2.0/instance"
	"github.com/Lucasfun/EasyOps/unitsInstall_2.0/instance/installController"
	"io/ioutil"
	"log"
)

func InstallTest() {
	filePath := "/Users/lucasmac/go/src/github.com/Lucasfun/EasyOps/unitsInstall_2.0/unitsConfig.yml"
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to get info from "+filePath, err)
	}
	uc, _ := instance.GetUnitsConfig(fileBytes)
	s := installController.InitService(uc.GetInDegree(), uc.GetOutDegree())
	s.Wait()
}
