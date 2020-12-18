package installController

import (
	"fmt"
	"github.com/Lucasfun/EasyOps/unitsInstall_2.0/interface"
)

type methodRequest _interface.UnitInterface

type Service struct {
	queue chan methodRequest
	wait  chan int
	In    map[string]int
	Out   map[string]_interface.UnitInterface
	N     int //已安装unit数
	flag  bool
}

func InitService(in map[string]int, out map[string]_interface.UnitInterface) *Service {
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
			go unitInterface.InstallFunc(s.ReportFunc(unitInterface))
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
				s.Reduction(mr)
				fmt.Println(mr.GetName() + " install finished.")
				s.N++
				if s.N == len(s.In) {
					return
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

func (s *Service) Reduction(u _interface.UnitInterface) { //减度操作
	for _, delName := range u.GetNext() { //eg:unitA出队，依赖unitA的Next[]全部 in-1
		s.In[delName] -= 1
		if s.In[delName] == 0 { //没有依赖其他unit 开始安装
			unit := s.Out[delName]
			go unit.InstallFunc(s.ReportFunc(unit))
		}
	}
}

//func (s *Service) GetQueue() chan {
//	return s.queue
//}

func (s *Service) ReportFunc(u _interface.UnitInterface) func() {
	return func() {
		s.queue <- u
	}
}
