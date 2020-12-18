package activeObject

import (
	"fmt"
	"log"
)

//怎么过的Unit的结构并在本包中使用呢？  -- 使用interface解决
type UnitInterface interface {
	InstallFunc(callback func()) bool
	GetName() string
	GetNext() []string
}

type methodRequest UnitInterface

type Service struct { //作为ActiveObject通讯对象
	queue chan methodRequest
	wait  chan int
	In    map[string]int //同张入、出度表
	Out   map[string]UnitInterface
	N     int //已安装unit数
}

func InitService(in map[string]int, out map[string]UnitInterface) *Service {
	s := &Service{
		queue: make(chan methodRequest, 0),
		wait:  make(chan int),
		In:    in,
		Out:   out,
		N:     0,
	}
	go s.Scheduler()
	for unit, n := range s.In {
		if n == 0 {
			unitInterface := s.Out[unit]
			go unitInterface.InstallFunc(func() {
				s.Report(unitInterface)
			})
		}
	}
	return s
}

func (s *Service) Scheduler() {
	defer func() {
		close(s.wait)
	}()

	for {
		select {
		case mr := <-s.queue:
			{
				isSuccess := s.Reduction(mr)
				if isSuccess {
					fmt.Println(mr.GetName() + "安装完成！")
					s.N++
					if s.N == len(s.In) {
						return
					}
				} else {
					log.Fatal(mr.GetName() + "安装失败，请重启安装程序！")
				}
			}
		case <-s.wait:
			return
		}
	}
}

func (s *Service) Wait() {
	<-s.wait
}

//减度操作
func (s *Service) Reduction(u UnitInterface) bool {
	for _, delName := range u.GetNext() { //eg:unitA出队，unitA的next[...]全部入度 - 1
		s.In[delName] -= 1
		if s.In[delName] == 0 { //入度为 0 开始安装
			unit := s.Out[delName]
			go unit.InstallFunc(func() {
				s.Report(unit)
			})
		}
	}
	return true
}

func (s *Service) Report(u UnitInterface) bool {
	s.queue <- u
	return true
}
