package design_pattern

import "testing"

func TestShapeFactory_GetShape(t *testing.T) {
	s := ShapeFactory{}
	circle := s.GetShape("CIRCLE")
	circle.Draw()

	rectangle := s.GetShape("RECTANGLE")
	rectangle.Draw()

	square := s.GetShape("SQUARE")
	square.Draw()
}
