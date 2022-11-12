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
	/* mock本质上是对测试过程中所依赖的一些函数/接口进行接管，
	即便真实调用尚未实现/不可直接调用也可以返回预期值，例如此处SUT函数依赖Foo接口的Bar方法，
	但是尚未有一个实现了Foo接口的结构体，但是可以利用mock来，接管SUT函数内的Bar调用，达到测试
	SUT函数的目的。
	*/
	m.EXPECT().Bar(gomock.Eq(99)).Return(99)

	SUT(m)
}
