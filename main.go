// Copyright (C) 2025 Adam B Kaplan
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"

	"github.com/adambkaplan/go-design-patterns/cmd/strategy"
)

func main() {
	// This is a simple Go program that prints "Hello, World!" to the console.
	// It will be updated as more design patterns are added.
	fmt.Println("Hello, Head First Design Patterns!")
	fmt.Println("== Pattern 1: Strategy Pattern ==")
	fmt.Println("---------------------------------")
	strategy.RunDuckSimulator()
	fmt.Println("---------------------------------")
}
