package topologicalSort

import "sync"

//拓扑排序得到Unit安装顺序，并执行并行安装
//inDegree  -- map[UnitType.Name] int
//outDegree -- map[UnitType.Next] UnitType

type num struct {
	mux    sync.Mutex
	insNum int
}

func TopologicalSort(infoMap *InfoMap) {

	q := Queue{}
	q.Init(len(infoMap.Degree)) //以所有unit个数作为队列容量初始化，确保不会出现队满

	//首先将入度为0的unit入队
	for unit, in := range infoMap.Degree { //O(N) = len(inDegree),实际执行 m <= N次
		if in == 0 {
			q.Enqueue(unit) //入度为0的string 对应的UnitType入队
		}
	}

	//安装完成个数  < 待安装组件个数 ： 入度为0的 Unit 执行安装程序
	n := num{sync.Mutex{}, 0}
	for n.insNum < len(infoMap.Degree) {
		for q.Head != q.Tail { //队列有Unit可取出来安装，一直循环
			headU := q.Dequeue()
			//install
			go func() {
				headU.UnitInstall(infoMap, &q)
				n.mux.Lock()
				defer n.mux.Unlock()
				n.insNum++
			}()
		}
	}
}

//减度 -- 同时删去对应uint.next[]中的name
func reduction(u *UnitType, info *InfoMap, q *Queue) {
	//上锁
	info.Mux.Lock()
	q.mux.Lock()
	//解锁
	defer info.Mux.Unlock()
	defer q.mux.Unlock()

	for _, delUnit := range u.Next { //遍历u.next,取出要减度的 unit.name
		for unit, _ := range info.Degree { //找出某name对应的 *Unit
			if unit.Name == delUnit {
				info.Degree[unit] -= 1      //减度
				if info.Degree[unit] == 0 { //入队
					q.Enqueue(unit)
				}
			}
		}
	}
}

func TestInstall() {
	unitConfig := UnitConfig{}
	unitConfig.InitUnitConfig()
	//for _,unitType := range unitConfig.Config{
	//	fmt.Println(unitType.Name)
	//	fmt.Println(unitType.Next)
	//	fmt.Println(unitType.Install)
	//}
	degree := GetInfoMap(&unitConfig)
	//fmt.Println(degree)
	TopologicalSort(degree)

}
