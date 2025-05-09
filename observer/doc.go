// Package observer and its sub-packages contains a port of the "Head First Design Patterns"
// Observer pattern, originally written in Java. This implementation deviates from the "Head First"
// implementation to take advantage of Golang concurrency primitives (goroutines and channels).
// In doing so, it adds further loose coupling between the Subject and Observer.
package observer
