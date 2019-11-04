# expector

Expector is a golang library for writing more concise tests.

It is heavily influenced by the matchers found in Jest.

## Usage

```go
package example

import (
  "testing"
  "github.com/allain/expector"
)

func TestExample1(t *testing.T) {
  expect := expector.New(t)

  // ToEqual using deep equality for comparison
  expect(10).ToEqual(10) 
  expect([]string{"a", "b"}).ToEqual([]string{"a", "b"})

  // Supports negation using Not()
  expect(10).Not().ToEqual(20) 

  // Supports matching string regexes
  expect("hello").ToMatch("^hell")
}
```