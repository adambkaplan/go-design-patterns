package strategy

import (
	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
	"github.com/adambkaplan/go-design-patterns/strategy/duck"
)

func RunDuckSimulator() {

	mallard := &duck.MallardDuck{}
	mallard.PerformQuack()
	mallard.PerformFly()

	modelDuck := &duck.ModelDuck{}
	modelDuck.PerformFly()
	modelDuck.FlyBehavior = &behavior.FlyRocketPowered{}
	modelDuck.PerformFly()

}
