// Copyright (C) 2025 Adam B Kaplan
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"

	"github.com/adambkaplan/go-design-patterns/cmd/decorator"
	"github.com/adambkaplan/go-design-patterns/cmd/factory"
	"github.com/adambkaplan/go-design-patterns/cmd/observer"
	"github.com/adambkaplan/go-design-patterns/cmd/strategy"
)

func main() {
	// This is a simple Go program that prints "Hello, World!" to the console.
	// It will be updated as more design patterns are added.
	fmt.Println("Hello, Head First Design Patterns!")
	fmt.Println("== Pattern 1: Strategy Pattern  ==")
	fmt.Println("----------------------------------")
	strategy.RunDuckSimulator()
	fmt.Println("----------------------------------")
	fmt.Println("== Pattern 2: Observer Pattern  ==")
	observer.RunWeatherStation()
	fmt.Println("----------------------------------")
	fmt.Println("== Pattern 3: Decorator Pattern ==")
	decorator.RunCoffeeShop()
	fmt.Println("----------------------------------")
	fmt.Println("====== Pattern 4: Factories ======")
	factory.RunPizzaStores()
	fmt.Println("----------------------------------")
}
