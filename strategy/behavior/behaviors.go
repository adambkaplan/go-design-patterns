// Copyright (C) 2025 Adam Kaplan
//
// SPDX-License-Identifier: MIT

package behavior

import "fmt"

// Flyable represents anything that has flying behavior.
type Flyable interface {
	Fly()
}

// FlyWithWings flies with wings.
type FlyWithWings struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings!")
}

// FlyNoWay doesn't fly.
type FlyNoWay struct{}

func (f FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

// FlyRocketPowered flies with rockets.
type FlyRocketPowered struct{}

func (f FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}

// Quackable represents anything that has quacking behavior.
type Quackable interface {
	Quack()
}

// Quack does your typical duck quack.
type Quack struct{}

func (q Quack) Quack() {
	fmt.Println("Quack!")
}

// MuteQuack is silent.
type MuteQuack struct{}

func (m MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

// Squeak is the quack that sounds like a squeak, perhaps from a rubber duck.
type Squeak struct{}

func (s Squeak) Quack() {
	fmt.Println("Squeak!")
}
