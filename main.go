// Copyright (C) 2025 Adam Kaplan
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"

	"github.com/adambkaplan/go-design-patterns/cmd"
)

func main() {
	// This is simple Go program that prints the outputs of the various pattern "test harnesses".
	fmt.Println("Head First Design Patterns in Go!")

	fmt.Println("Pattern 1: Strategy Pattern with the Duck Simulator")
	cmd.RunDuckSimulator()
}
