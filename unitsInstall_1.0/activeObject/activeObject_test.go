package activeObject

import (
	"github.com/Lucasfun/EasyOps/unitsInstall_1.0/mocks/activeObject"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestService_Reduction(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	unitA := activeObject.NewMockUnitInterface(ctl)
	unitB := activeObject.NewMockUnitInterface(ctl)
	unitC := activeObject.NewMockUnitInterface(ctl)
	s := &Service{
		queue: make(chan methodRequest),
		wait:  nil,
		In:    map[string]int{"A": 0, "B": 0, "C": 2},
		Out:   map[string]UnitInterface{"A": unitA, "B": unitB, "C": unitC},
		N:     0,
	}

	gomock.InOrder(
		unitA.EXPECT().GetNext().Return([]string{"C"}).Times(1),
		unitB.EXPECT().GetNext().Return([]string{"C"}).Times(1),
		unitC.EXPECT().InstallFunc(gomock.Any()).Times(1),
	)

	s.Reduction(unitA)
	s.Reduction(unitB)
	time.Sleep(time.Second * 1)

	want := map[string]int{"A": 0, "B": 0, "C": 0}
	got := s.In
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Reduction() = %v, want %v", s.In, want)
	}
}

//InitService ：init一个*Service，完成后拉起一个 scheduler goroutine；for range s.in，若unitX.in == 0,拉起unitX的installFunc
func TestInitService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	unitA := activeObject.NewMockUnitInterface(ctl)
	unitB := activeObject.NewMockUnitInterface(ctl)
	unitC := activeObject.NewMockUnitInterface(ctl)

	inMap := map[string]int{"A": 0, "B": 0, "C": 2}
	outMap := map[string]UnitInterface{"A": unitA, "B": unitB, "C": unitC}

	want := &Service{
		In:  map[string]int{"A": 0, "B": 0, "C": 2},
		Out: map[string]UnitInterface{"A": unitA, "B": unitB, "C": unitC},
	}

	unitA.EXPECT().InstallFunc(gomock.Any()).Return(true).Times(1)
	unitB.EXPECT().InstallFunc(gomock.Any()).Return(true).Times(1)
	s := InitService(inMap, outMap)
	time.Sleep(time.Second * 7)
	// 测试s init后in、out一致、且拉起Unit的install次数与顺序正确即可
	if !(reflect.DeepEqual(s.In, want.In) && reflect.DeepEqual(s.Out, want.Out)) {
		t.Errorf("s.In,s.Out,s.N= %v,%v, want %v", s.In, s.N, want.In)
	}

}

//测试go s.Scheduler执行时，s.queue <- unitDemo是否有效
func TestService_Scheduler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	unitA := activeObject.NewMockUnitInterface(ctl)
	unitB := activeObject.NewMockUnitInterface(ctl)
	unitC := activeObject.NewMockUnitInterface(ctl)

	gomock.InOrder(
		unitA.EXPECT().GetNext().Return([]string{"B"}),
		unitA.EXPECT().GetName().Return("A"),
		unitB.EXPECT().InstallFunc(gomock.Any()),
	)
	want := &Service{
		In: map[string]int{"A": 0, "B": 0, "C": 1},
		N:  2,
	}
	s := &Service{
		queue: make(chan methodRequest, 0),
		wait:  make(chan int),
		In:    map[string]int{"A": 0, "B": 1, "C": 1},
		Out:   map[string]UnitInterface{"A": unitA, "B": unitB, "C": unitC},
		N:     1,
	}
	go s.Scheduler()
	//test 1
	s.Report(unitA)
	time.Sleep(time.Second * 3)
	//test 2
	s.wait <- 0

	if !(reflect.DeepEqual(want.In, s.In) && reflect.DeepEqual(want.N, s.N)) {
		t.Errorf("serviceDemo.In = %v,want.In = %v;"+
			"serviceDemo.N = %v,want.N = %v", s.In, want.In, s.N, want.N)
	}

}
