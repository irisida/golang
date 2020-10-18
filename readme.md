![](https://github.com/irisida/golang/blob/master/assets/gopher.jpeg)

# Repo contents

- [Go language lessons and tutorials](https://github.com/irisida/golang/tree/master/language_lessons)
- [projects in Go](https://github.com/irisida/golang/tree/master/projects/)

# Additional info

In building this repository several golang Udemy resources were used and are refereced.

- [Go - Developers guide](https://www.udemy.com/course/go-the-complete-developers-guide)
- [Master Go programming](https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp).
- [Go - complete bootcamp](https://www.udemy.com/course/learn-go-the-complete-bootcamp-course-golang)

# Go resources

- [Official Go website](https://golang.org/)
- [Language reference documentation](https://golang.org/doc/)
- [Go playground](https://play.golang.org/)

# Installation and setup common best practice

- Install golang for your system from [here](https://golang.org/)
- Set the GOPATH to `Users/username/go`
- Within the GOPATH directory create three subdirectories, `bin`, `pkg` & `src`
- within the `src` directory create a `github.com` directory and your GH username as a child of that. From within that child directory you will create and store your Go projects.

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
