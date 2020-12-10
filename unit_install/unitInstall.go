package unit_install

func TopologicalSort(uc *UnitConfig)  {

	s := InitService(uc) //拉起Scheduler goroutine

	s.Wait()
}


func TestDemo() {
	uc := GetUnitConfig()
	TopologicalSort(uc)
}

