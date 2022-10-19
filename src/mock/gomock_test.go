package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

//  mockgen -source=foo.go -package=mock -destination mock_foo.go

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	m.EXPECT().Bar(gomock.Eq(99)).Return(101)

	SUT(m)
}
