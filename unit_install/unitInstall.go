package unit_install


func TopologicalSort(uc *UnitConfig,q *Queue)  {
	//首先将入度为0的unit入队
	for unit, in := range uc.inDegree {
		if in == 0 {
			unit := uc.outDegree[unit]
			q.Enqueue(unit)
		}
	}

	s := InitService(q,uc) //拉起Scheduler goroutine

	for s.N < len(uc.inDegree) {	//安装完成个数  < 待安装组件个数 ： 入度为0的Unit执行安装程序
		for q.Head != q.Tail { 	//队列有Unit可取出来安装，一直循环
			deU := q.Dequeue()
			go deU.UnitInstall(s.queue)
		}
	}
}


func TestDemo() {
	uc := GetUnitConfig()
	q := GetQueue(len(uc.inDegree))
	TopologicalSort(uc,q)
}

