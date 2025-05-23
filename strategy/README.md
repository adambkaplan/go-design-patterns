# Strategy Pattern

The Strategy pattern is used to encapsulate object behavior that can change at runtime. In the
book's example, a hypothetical "duck simulator" presents an assortment of ducks on screen, each
type with their own "flying" and "quacking" behavior. Each duck type has a default flying and
quacking behavior, however this can change at runtime.

The Strategy pattern example in "Head First Design Patterns" relies on Java's polymorphism
model to provide an "abstract" implementation of a Duck. This contains all the method
definitions of a "Duck" supertype and shared behavior, but requires concrete implementations
to provide the remaining gaps.

Golang eschews object-oriented polymorphism altogheter, favoring composition instead. This
implementation uses a true interface to define the behavior of all ducks, and programs to it in
the unit test suite. Code reuse is encouraged through a "GenericDuck" struct type, which
provides a working interface implementation. The other duck types embed this struct, overriding
functions as needed to alter behavior.
