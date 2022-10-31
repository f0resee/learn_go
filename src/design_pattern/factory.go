package design_pattern

import "fmt"

var _ Shape = (*Rectangle)(nil) // 检查一个结构体是否实现了Shape接口
var _ Shape = (*Square)(nil)
var _ Shape = (*Circle)(nil)

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::Draw() method.")
}

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("Inside Square::Draw() method.")
}

type Circle struct {
}

func (s *Circle) Draw() {
	fmt.Println("Inside Circle::Draw() method.")
}

type ShapeFactory struct {
}

func (s *ShapeFactory) GetShape(shapeType string) Shape {
	switch shapeType {
	case "CIRCLE":
		return &Circle{}
	case "RECTANGLE":
		return &Rectangle{}
	case "SQUARE":
		return &Square{}
	}
	return nil
}
