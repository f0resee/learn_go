package abstract_factory

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r Rectangle) Draw() {
	fmt.Println("Inside Rectangle::Draw() method.")
}

type Square struct {
}

func (r Square) Draw() {
	fmt.Println("Inside Square::Draw() method.")
}

type Circle struct {
}

func (c Circle) Draw() {
	fmt.Println("Inside Circle::Draw() method.")
}

type Color interface {
	Fill()
}

type Red struct {
}

func (r Red) Fill() {
	fmt.Println("Inside Red::Fill() method.")
}

type Green struct {
}

func (r Green) Fill() {
	fmt.Println("Inside Red::Fill() method.")
}

type Blue struct {
}

func (r Blue) Fill() {
	fmt.Println("Inside Blue::Fill() method.")
}

type AbstractFactory interface {
	GetColor(color string) Color
	GetShape(shape string) Shape
}

type ColorFactory struct {
}

func (s ColorFactory) GetColor(color string) Color {
	if color == "" {
		return nil
	}
	switch color {
	case "RED":
		return Red{}
	case "GREEN":
		return Green{}
	case "BLUE":
		return Blue{}
	}
	return nil
}

func (s ColorFactory) GetShape(shape string) Shape {
	return nil
}

type ShapeFactory struct {
}

func (s ShapeFactory) GetColor(color string) Color {
	return nil
}

func (s ShapeFactory) GetShape(shape string) Shape {
	if shape == "" {
		return nil
	}
	switch shape {
	case "CIRCLE":
		return Circle{}
	case "RECTANGLE":
		return Rectangle{}
	case "SQUARE":
		return Square{}
	}
	return nil
}

type FactoryProducer struct {
}

func (f FactoryProducer) GetFactory(choice string) AbstractFactory {
	switch choice {
	case "SHAPE":
		return ShapeFactory{}
	case "COLOR":
		return ColorFactory{}
	}
	return nil
}
