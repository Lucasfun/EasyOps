package unit_install


func TopologicalSort(uc *UnitConfig,q *Queue)  {
	//首先将入度为0的unit入队
	for unit, in := range uc.Degree {
		if in == 0 {
			q.Enqueue(unit)
		}
	}

	s := InitService(q,uc.Degree) //拉起Scheduler goroutine

	for s.N < len(uc.Degree) {	//安装完成个数  < 待安装组件个数 ： 入度为0的Unit执行安装程序
		for q.Head != q.Tail { 	//队列有Unit可取出来安装，一直循环
			deU := q.Dequeue()
			go func() {
				isSuccess := deU.UnitInstall()
				if isSuccess{
					s.queue <-deU  //减度channel中添加任务
				}
			}()
		}
	}
}


func TestDemo() {
	uc := GetUnitConfig()
	q := GetQueue(len(uc.Degree))
	TopologicalSort(uc,q)
}

