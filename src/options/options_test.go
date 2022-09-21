package options

import (
	"log"
	"testing"
)

type apple interface {
	Size() int
}

type tapple struct {
	color  string
	size   int
	weight float32
}

func (t *tapple) Size() int {
	return t.size
}

type appBuilder func(*options) (apple, error)

type options struct {
	color  string
	size   int
	weight float32
}

type option func(*options)

func color(c string) option {
	return func(o *options) {
		o.color = c
	}
}

func size(s int) option {
	return func(o *options) {
		o.size = s
	}
}

func weight(w float32) option {
	return func(o *options) {
		o.weight = w
	}
}

var aab tapple

func getab() (appBuilder, error) {
	return mp["aaa"], nil
}

var mp = make(map[string]appBuilder)

func insertmp(a string, b appBuilder) {
	mp[a] = b
}

func newab(opts *options) (apple, error) {
	return &tapple{
		color:  opts.color,
		size:   opts.size,
		weight: opts.weight,
	}, nil
}

func newApple(opts ...option) (apple, error) {
	var oo options
	for _, opt := range opts {
		opt(&oo)
	}
	//var ab appBuilder
	insertmp("aaa", newab)
	ab, err := getab() // 获取想要的builder
	if err == nil {

	}
	log.Println(oo)
	return ab(&oo) //使用builder生成想要的接口
}

func TestApple(t *testing.T) {
	oc := color("red")
	os := size(1)
	ow := weight(1.2)
	ap, err := newApple(oc, os, ow)
	if err != nil {
	}
	t.Log(ap.Size())
}
