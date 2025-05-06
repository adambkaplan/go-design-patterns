# Go Design Patterns

Design Patterns for Go - adapted from Head First Design Patterns (2nd Edition) by Eric Freeman and
Elisabeth Robson.

## Notes

This repository adapts the design patterns in the book to the Go programming language. Unlike Java
(which is used as the canoncial example in the book), Go is **not** an object-oriented programming
language. Patterns that rely on polymorphism/class inheritance will thus be modified to rely less
on these techniques. Some "inherited" code may be copied, and that is okay;
[a little copying is better than a little dependency](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s).

This repository will also take a [test-driven development](https://www.agilealliance.org/glossary/tdd/)
approach to writing code. Most structs and functions will be verified by tests using the
[Ginkgo](https://onsi.github.io/ginkgo/) test framework. This framework emphasizes testing
_behaviors_, which makes the test suite more resilient to change.

## License

Copyright (C) 2025 Adam B Kaplan

Licensed under the [MIT License](./LICENSE)
