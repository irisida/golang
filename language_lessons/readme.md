![](https://github.com/irisida/golang/blob/master/assets/freegopher.png)

# Go language coverage

lessons, tutorials, snippets and works in progress for the Go language.

- [Day 01 - Getting started](https://github.com/irisida/golang/tree/master/language_lessons/day01_gettingStarted)
- [Day 2](https://github.com/irisida/golang/tree/master/language_lessons/chapter02_typesAndDatatypes)
- [Chapter Three]()

# Simple pointers lesson

- Turn `address` into `value` with `*address`
- Turn `value` into `address` with `&value`

# Pointers shortcut in Go

In Go you can shortcircuit the process of:

- create variable holding address `ptr := &target`
- call function that receives a pointer with `ptr`

Instead you can pas the object and Go is smart enough to make that converion/leap for you.

```go
/* assume a struct that takes two values to make a person */
func main() {
	emp := person{
		firstName: "Bez",
		lastName:  "Mondo",
	}

    emp.print()
    /* call func that expects a pointer, on the actual object */
	emp.updateNameValue("Ed")
	emp.print()
}

/* note the func receiver */
func (p *person) updateNameValue(receivedNameValue string) {
	(*p).firstName = receivedNameValue
}
```

# Maps Vs Structs

| Maps                                                | Structs                                                       |
| --------------------------------------------------- | ------------------------------------------------------------- |
| All keys must be of the same type                   | keys don't support indexing                                   |
| All values must be of the same type                 | Values can be of different type                               |
| Use to represent a collection of related properties | Use to represent something with a lot of different properties |
| don't need to know keys at compile time             | Needs to know all the different filds at compile time         |
| reference type                                      | value type                                                    |

# interfaces

Interfaces are just a type that other types can have an honourary badge as _being of_. Other languages mention of `implementing an interface` and here we see it more as a loose `is also a` relationship. Therefore we will see in Go terms the notion of concrete types and interface types.

| Concrete types  | Interface types        |
| --------------- | ---------------------- |
| map             |                        |
| struct          |                        |
| int             |                        |
| string          |                        |
| userDefinedType | user defined interface |

- Go does not have generics.
- Interfaces are not generic types, as seen in other languages.
- Interfaces are implicit. We don't have to manually have to say that a custom type satisfies an interface.
- Interfaces are a contract to manage types.
- Interfaces are considered as advanced or tough. Requiring experience and some insight/forethought.
