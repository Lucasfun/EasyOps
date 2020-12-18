package user

import (
	"github.com/Lucasfun/EasyOps/gomockDemo/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUser_GetUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(nil),
	)
	user := NewUser(mockMale)
	err := user.GetUser(id)
	if err != nil {
		t.Errorf("user.GetUser err: %v", err)
	}
}
