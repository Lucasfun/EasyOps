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

type UnitConfigInterface interface {
	GetInDegree() map[string]int
	GetOutDegree() map[string]UnitInterface
}

type methodRequest UnitInterface

type Service struct { //作为ActiveObject通讯对象
	queue chan methodRequest
	wait  chan int
	In    map[string]int //同张入、出度表
	Out   map[string]UnitInterface
	N     int //已安装unit数
}

type Server interface {
	Scheduler()
	Wait()
	Reduction(unit UnitInterface) bool
	Report(unit UnitInterface) bool
}

func InitService(config UnitConfigInterface) *Service {
	s := &Service{
		queue: make(chan methodRequest, 0),
		wait:  make(chan int),
		In:    config.GetInDegree(),
		Out:   config.GetOutDegree(),
		N:     0,
	}
	go s.Scheduler()
	for unit, in := range config.GetInDegree() {
		if in == 0 {
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
		if s.In[delName] == 0 { //入度为 0 入队
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
