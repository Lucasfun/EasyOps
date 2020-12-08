package activeObjectDemo

// 封装操作方法【那个被并发操作可能会出data race问题的相关变量相关】的上下文
type MethodRequest struct {
	//anything
}

type Service struct {
	queue chan MethodRequest	//相应mr的通道，用于线程减的通信
	v     int					//该引发问题的变量
}
func New(buffer int) *Service {	//初始化Service的时候同时拉起调度器Scheduler goroutine
	s := &Service{
		queue: make(chan MethodRequest, buffer),
	}
	go s.schedule()
	return s
}
func (s *Service) schedule() {
	//对s.queue相关调度

	//一直监听？
	//	for{
	//		select {
	//		case <- MethodRequest
	//		default:
	//
	//		}
	//	}

	//循环处理
	//	for mr := range s.queue {
	//
	//	}
}
