package abstract_factory

import "testing"

func TestAbstractFactoryPattern(t *testing.T) {
	factoryProducer := FactoryProducer{}

	//
	shapeFactory := factoryProducer.GetFactory("SHAPE")
	shape1 := shapeFactory.GetShape("CIRCLE")
	shape1.Draw()

	shape2 := shapeFactory.GetShape("RECTANGLE")
	shape2.Draw()

	shape3 := shapeFactory.GetShape("SQUARE")
	shape3.Draw()

	//
	colorFactory := factoryProducer.GetFactory("COLOR")
	color1 := colorFactory.GetColor("RED")
	color1.Fill()

	color2 := colorFactory.GetColor("GREEN")
	color2.Fill()

	color3 := colorFactory.GetColor("BLUE")
	color3.Fill()
}
