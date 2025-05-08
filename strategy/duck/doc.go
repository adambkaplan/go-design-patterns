// duck defines the Duck interface and provides implementations for various kinds of Ducks.
package duck

// The Strategy pattern example in "Head First Design Patterns" relies on Java's polymorphism
// model to provide an "abstract" implementation of a Duck. This contains all the method
// definitions of a "Duck" supertype and shared behavior, but requires concrete implementations
// to provide the remaining gaps.

// Golang eschews object-oriented polymorphism altogheter, favoring composition instead. This
// implementation uses a true interface to define the behavior of all ducks, and programs to it in
// the unit test suite. Code reuse is encouraged through a "GenericDuck" struct type, which
// provides a working interface implementation. The other duck types embed this struct, overriding
// functions as needed to alter behavior.
