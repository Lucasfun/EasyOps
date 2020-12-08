package unit_install

import (
	"fmt"
	"log"
)

type methodRequest *Unit

type Service struct { //作为ActiveObject通讯对象
	queue   chan methodRequest
	Q 		*Queue
	D 		map[*Unit]int
	N       int
}

func InitService(Q *Queue,D map[*Unit]int) *Service {  //处理所有unit对应的减度操作
	S := &Service{
		queue: 		make(chan methodRequest,0),
		Q: Q,
		D: D,
		N: 0,
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
	for _,delName := range u.Next{
		for unit,_ := range s.D{
			if delName == unit.Name{
				s.D[unit] -= 1
				if s.D[unit] == 0{
					s.Q.Enqueue(unit)
				}
			}
		}
	}
	return true
}


