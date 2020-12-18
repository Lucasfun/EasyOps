package installController

import (
	_interface "github.com/Lucasfun/EasyOps/unitsInstall_2.0/interface"
	"github.com/Lucasfun/EasyOps/unitsInstall_2.0/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestInitService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	unitA := mocks.NewMockUnitInterface(ctl)
	unitB := mocks.NewMockUnitInterface(ctl)
	unitC := mocks.NewMockUnitInterface(ctl)

	inMap := map[string]int{"A": 0, "B": 0, "C": 2}
	outMap := map[string]_interface.UnitInterface{"A": unitA, "B": unitB, "C": unitC}

	want := &Service{
		In:  map[string]int{"A": 0, "B": 0, "C": 2},
		Out: map[string]_interface.UnitInterface{"A": unitA, "B": unitB, "C": unitC},
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

func TestService_Scheduler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	unitA := mocks.NewMockUnitInterface(ctl)
	unitB := mocks.NewMockUnitInterface(ctl)
	unitC := mocks.NewMockUnitInterface(ctl)

	unitA.EXPECT().GetNext().Return([]string{"B"})
	unitA.EXPECT().GetName().Return("A")
	unitB.EXPECT().InstallFunc(gomock.Any())
	unitB.EXPECT().GetNext().Return("C")
	unitB.EXPECT().GetName().Return("B")
	unitB.EXPECT().InstallFunc(gomock.Any())

	want := &Service{
		In: map[string]int{"A": 0, "B": 0, "C": 0},
		N:  3,
	}
	s := &Service{
		queue: make(chan methodRequest, 0),
		wait:  make(chan int),
		In:    map[string]int{"A": 0, "B": 1, "C": 1},
		Out:   map[string]_interface.UnitInterface{"A": unitA, "B": unitB, "C": unitC},
		N:     1,
	}
	go s.Scheduler()
	//test 1
	//s.Report(unitA)
	//s.Report(unitB)
	time.Sleep(time.Second * 7)
	//test 2
	s.wait <- 0

	if !(reflect.DeepEqual(want.In, s.In) && reflect.DeepEqual(want.N, s.N)) {
		t.Errorf("serviceDemo.In = %v,want.In = %v;"+
			"serviceDemo.N = %v,want.N = %v", s.In, want.In, s.N, want.N)
	}

}

func Test_installFunc(t *testing.T) {
	type args struct {
		unitInterface _interface.UnitInterface
		s             *Service
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
