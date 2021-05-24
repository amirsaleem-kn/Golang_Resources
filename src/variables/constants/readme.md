## Constants in Go

Constants have fixed values. Once assigned, the value of constants cannot be changed. In other words, constants are read-only and immutable.

__Some points to remember about constants__

- Constants are declared using a `const` keyword
- The declaration and assignment of constants must happen in the same line
- Once assigned, the value of constants cannot be modified.
- Constants cannot be declared using a short-variable syntax i.e. `:=`

__How to declare constants?__

Constants are declared using a `const` keyword. The declaration and assignment must happen in the same line. The syntax is as follows:

```go
const message = "Golang is a programming language" // string constant
const rating = 2.4 // floating constant
const percentage = 100 // integer constant

// re-assignment is not allowed. This line will throw the error
message = "Golang is not a programming language"
```
