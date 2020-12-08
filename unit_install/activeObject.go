package unit_install

import (
	"fmt"
	"log"
)

type methodRequest *Unit

type Service struct { //作为ActiveObject通讯对象
	queue   chan methodRequest
	Q 		*Queue				//topologicalSort所用队列引用
	In      map[string]int		//同张入、出度表
	Out 	map[string]*Unit
	N       int					//已安装unit数
}

func InitService(Q *Queue,uc *UnitConfig) *Service {  //处理所有unit对应的减度操作
	S := &Service{
		queue: 		make(chan methodRequest,0),
		Q: 			Q,
		In: 		uc.inDegree,
		Out: 		uc.outDegree,
		N: 			0,
	}
	go S.Scheduler()
	return S
}

func (s *Service) Scheduler()  {
	for{
		select {
		case unit := <- s.queue:{
			isSuccess := s.Reduction(unit)
			if isSuccess{
				fmt.Println(unit.Name + "finish install.")
				s.N ++
			}else{
				log.Fatal(unit.Name + "安装失败，请重启安装程序！")
			}
		}
		}
	}
}

//减度操作
func (s *Service) Reduction(u *Unit) bool {	//TODO 检查出错？
	for _,delName := range u.Next{	//eg:unitA出队，unitA的next[...]全部入度 - 1
		s.In[delName] -= 1
		if s.In[delName] == 0{	//入度为 0 入队
			unit := s.Out[delName]
			unit.UnitInstall(s.queue)
		}
	}
	return true
}


