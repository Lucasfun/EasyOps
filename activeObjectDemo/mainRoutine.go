package activeObjectDemo

//原本的设计，需求就是并发对Service中的v进行读写，造成date race问题

//	type Service struct {
//		v int
//	}
//	func (s *Service) Incr() {
//		s.v++
//	}
//	func (s *Service) Decr() {
//		s.v--
//	}

//	s := newService()
//	{
//		go goroutineDemo() {
//			s.queue <- 某个methodRequestDemo
//		}
//	}