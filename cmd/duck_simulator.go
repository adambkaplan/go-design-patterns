// Copyright (C) 2025 Adam Kaplan
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
	"github.com/adambkaplan/go-design-patterns/strategy/duck"
)

func RunDuckSimulator() {
	mallardDuck := duck.NewMallardDuck()
	mallardDuck.PerformQuack()
	mallardDuck.PerformFly()

	modelDuck := duck.NewModelDuck()
	modelDuck.PerformFly()
	modelDuck.FlyBehavior = behavior.FlyRocketPowered{}
	modelDuck.PerformFly()
}
