# bitfield
A basic variable length bit field implementation ported from [this stack overflow answer](https://stackoverflow.com/a/177092)

# Usage

```go
package main

import (
    "fmt"

    "github.com/purehyperbole/bitfield"
)

type Item int

const (
    Eggs Item = iota
    Milk
    Sugar
    Pasta
    Cheese
    Tomatoes
)

func (i Item) String() string {
    return [...]string{"Eggs", "Milk", "Sugar", "Pasta", "Cheese", "Tomatoes"}[i]
}

func main() {
    // creates a bitfield with 128 different possible flags
    shoppingList := bitfield.New(16)

    // set some flags
    shoppingList.Set(Eggs)
    shoppingList.Set(Pasta)
    shoppingList.Set(Cheese)

    for _, item := range []int{Eggs, Milk, Sugar, Pasta, Cheese, Tomatoes} {
        // check if the list has flags
        if shoppingList.Has(item) {
            fmt.Println("need:", item)
        }
    }
}
```
