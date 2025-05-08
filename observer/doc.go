// Package observer and its sub-packages contains a direct port of the "Head First Design Patterns"
// Observer pattern, originally written in Java. This implementation claims to demonstrate "loose
// couplinng," however each display requires a direct reference to the main Subject object.
//
// This implementation is not recommened for golang as it does not take advantage of its native
// concurrency features (goroutines, channels, and context.Context objects). The Observer pattern
// is otherwise ideal for concurrent workloads.
package observer
