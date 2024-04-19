# Assert
Assert is a very simple, dependency free assertion package for go.

# Motivation
Generally, assertions are discouraged in Go as they provide a mechanism for lazy error handling. While that is true, they also allow to check your programs for correctness and shutdown if something completely unexpected happens.

This package was motivated by the [TigerStyle](https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md) programming practices. They also have an awesome [talk](https://youtu.be/w3WYdYyjek4?si=EqGZFovJyIyl6VqK&t=1527) that illustrates why assertions are good.

# Usage
Call `assert.Assert()` to check for invariants in your code and make sure it shutsdown if they are violated.
Optionally you can configure the panic message and add a writer to which the message and a trace will be written before panicking (e.g. to a Log file or distributed logging service in the cloud). 

```go
package main

import (
    "assert"
    "log"
)

func main() {
    x := 2
    logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

    // this will panic with a custom message and will log the panic before crashing with a trace to [main.go:12].
    mustBeGreaterThanFive(x, logger)
}


func mustBeGreaterThanFive(x int, logger log.Logger) {
    // make sure the invariant of the function is met
    assert.Assert(x > 5, assert.WithMessage("x must always be greater than 5"), assert.WithWriter(logger.Writer()))

    // do stuff that only work if x > 5

}
```
