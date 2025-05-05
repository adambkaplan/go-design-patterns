// Copyright (C) 2025 Adam Kaplan
//
// SPDX-License-Identifier: MIT

package duck

import (
	"fmt"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
)

// duck is the equivalent of an "abstract class" in Go. It is internal to the `duck` package and
// can't be instantiated outside of the package. "Sub-classes" inherit the `duck` behavior through
// type embedding.
type duck struct {
	FlyBehavior   behavior.Flyable
	QuackBehavior behavior.Quackable
}

// PerformFly is a method that performs the fly behavior of a Duck.
// If the Duck has a FlyBehavior set, it will be called.
// Otherwise, the method will return without performing any action.
//
// Assisted by watsonx Code Assistant
func (d *duck) PerformFly() {
	if d.FlyBehavior != nil {
		d.FlyBehavior.Fly()
	}
	return
}

// PerformQuack is a method that performs the quack behavior.
// It takes no input parameters and returns nothing.
//
// Assisted by watsonx Code Assistant
func (d *duck) PerformQuack() {
	if d.QuackBehavior != nil {
		d.QuackBehavior.Quack()
	}
	return
}

// Swim is a method that makes the duck swim.
// It prints a message to the standard output and returns nothing.
//
// Assisted by watsonx Code Assistant
func (d *duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
	return
}

// MallardDuck represents a real Mallard duck.
type MallardDuck struct {
	duck
}

// NewMallardDuck creates a new MallardDuck instance.
//
// Assisted by watsonx Code Assistant
func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		duck: duck{
			FlyBehavior:   behavior.FlyWithWings{},
			QuackBehavior: behavior.Quack{},
		},
	}
}

// Display shows the duck to the console.
func (m *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}

// ModelDuck represents a model duck that can be adapted.
type ModelDuck struct {
	duck
}

// NewModelDuck creates a new ModelDuck instance.
func NewModelDuck() *ModelDuck {
	return &ModelDuck{
		duck: duck{
			FlyBehavior:   behavior.FlyNoWay{},
			QuackBehavior: behavior.MuteQuack{},
		},
	}
}

// Display shows the duck to the console.
func (m *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}
