# Observer Pattern

In the observer pattern, an "observing" resource registers with a source "subject" resource, and by
doing so receives notifications of changes. Many front-end developers implicitly use this pattern
by registering callback functions with buttons and other visual elements.

This implementation deviates from the "Head First" implementation to take advantage of Golang 
concurrency primitives (goroutines and channels). In doing so, it adds further loose coupling
between the Subject and Observer.
