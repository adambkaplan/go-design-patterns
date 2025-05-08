package strategy

import (
	"fmt"
	"os"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
	"github.com/adambkaplan/go-design-patterns/strategy/duck"
)

func RunDuckSimulator() {

	mallard := &duck.MallardDuck{}
	Must(mallard.PerformQuack())
	Must(mallard.PerformFly())

	modelDuck := &duck.ModelDuck{}
	Must(modelDuck.PerformFly())
	modelDuck.FlyBehavior = &behavior.FlyRocketPowered{}
	Must(modelDuck.PerformFly())

}

func Must(err error) {
	if err != nil {
		fmt.Printf("An unexpected error occurred: %v", err)
		os.Exit(1)
	}
}
