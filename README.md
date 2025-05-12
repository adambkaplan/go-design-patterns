# Go Design Patterns

Design Patterns for Go - adapted from Head First Design Patterns (2nd Edition) by Eric Freeman and
Elisabeth Robson.

## Does Go Need Design Patterns?

This repository adapts the design patterns in the book to the Go programming language. Unlike Java
(which is used as the canoncial example in the book), Go is **not** an object-oriented programming
language. Some patterns translate well to Golang; others, not so much.

This repository will also take a [test-driven development](https://www.agilealliance.org/glossary/tdd/)
approach to writing code. Most structs and functions will be verified by tests using the
[Ginkgo](https://onsi.github.io/ginkgo/) test framework. This framework emphasizes testing
_behaviors_, which makes the test suite more resilient to change.

### Design Basics

- **Abstraction**: Only expose what clients/callers need to know. Hide what is not necessary.
  Present _what_ an entity does, not _how_ it does it.
- **Encapsulation**: Group together information and behaviors that logically belong together.
  Control access to information whose state needs to be tightly managed.
- **Composition**: Share behavior through combinations. Inheritance and polymorphisim do not exist
  in go (mostly).
- **Valuation**: Unlike Java and other object-oriented languages, errors and functions are values
  in Go and can be referenced just like a `string` or `int`.

### Principles and Proverbs

- Encapsulate what varies.
- Favor composition.
- Program to interfaces ~~, not implementations~~ when the interface is all you need.
- The bigger the interface, the weaker the abstraction [proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s).
- Strive for loose coupling between entites that interact.
- Make the zero value useful [proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s).
- A little copying is better than a little dependency [proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s).
- Clear is better than clever [proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s).
- Handle errors gracefully [proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=17m25s).
- Don't communicate by sharing memory, share memory by communicating [proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s).

## Patterns

1. [Strategy Pattern](./strategy/README.md)
2. [Observer Pattern](./observer/README.md)
3. [Decorator Pattern](./decorator/README.md)

## License

Copyright (C) 2025 Adam B Kaplan

Licensed under the [MIT License](./LICENSE)
