![](https://github.com/irisida/golang/blob/master/src/assets/freegopher.png)

# Lesson 1

## Installation & Setup common best practices

Up until 1.11 Go was quite strict on the package management structure and expected all work to be in a single repository located at `$HOME/go/src` and more typically at `$HOME/go/src/github.com/username/repository_name`. Since go 1.12 go modules allow for a different location and packaging system. This repo will adhere to the original structure because:

- That's how my machine is setup currently
- The majority of learning resources you will stumble upon while learning Go will still refer to this
- migrating to Go modules once I know what I'm doing will give a useful reference point of working with both flavours of Go structure.

- Install golang for your system from [here](https://golang.org/)
- Set the GOPATH to `Users/username/go`
- Within the GOPATH directory create three subdirectories, `bin`, `pkg` & `src`
- within the `src` directory create a `github.com` directory and your GH username as a child of that. From within that child directory you will create and store your Go projects.

## OK, hit me with the day 1 jazz...

So, we're going to take some assumptions on the basis that it's `strongly typed` and has a basis in C and C-like languages. By this the older cynic can deduce the intention of the project was always to move beyond an internal tool, to keep tried, tested and battle-hardened concepts from an absolute warhorse of a language. So we may also gather that they genuinely are looking to have ease at the core of the language and concession #1 was to use/re-use syntactic constructs from the most popular languages of the last 40+ years.

## Hello World

Let's see the hello world nonsense then....

```go
package main

import "fmt"

func main() {
	fmt.Println("hello, World")
}
```

## Statements and Expressions

Like every other successful language we have to think about statements and expressions and like every other language:

- Statements are operations and may contain one or many expressions. Instructions to do something. eg. fmt.Println()
- expressions are returned values, here it is "This is a go program!"

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*
		1. runtime.NumCPU() is an expression, it returns a value.
		2. fmt.Printlin() is a statement. It tells Go to perform an action, to do something.
	*/

	fmt.Println(runtime.NumCPU())
}
```

## Makefiles

Now, if you're thinking "how do we run this?" well, me too! In short we now have two possibilities:

- Execute with `go run <filename>` which will compile and execute the program immediately.
- Build with `go build <filename>` which starts the build mechanism that if compilation is successful then the build tool will produce an executable file that can be run, re-run later.

To understand the steps involved in a `go run` request behind the scenes:

- below in the out we can see the tool linking in packages
- Internal system calls etc...

```
WORK=/tmp/go-build267267774
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/home/ed/.cache/go-build/98/98abe4a0683ebb0296e3931266e333ae8ac23cce798fc1222addf5c713f66f3c-d
packagefile fmt=/usr/local/go/pkg/linux_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/linux_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/linux_amd64/errors.a
packagefile internal/fmtsort=/usr/local/go/pkg/linux_amd64/internal/fmtsort.a
packagefile io=/usr/local/go/pkg/linux_amd64/io.a
packagefile math=/usr/local/go/pkg/linux_amd64/math.a
packagefile os=/usr/local/go/pkg/linux_amd64/os.a
packagefile reflect=/usr/local/go/pkg/linux_amd64/reflect.a
packagefile strconv=/usr/local/go/pkg/linux_amd64/strconv.a
packagefile sync=/usr/local/go/pkg/linux_amd64/sync.a
packagefile unicode/utf8=/usr/local/go/pkg/linux_amd64/unicode/utf8.a
packagefile internal/bytealg=/usr/local/go/pkg/linux_amd64/internal/bytealg.a
packagefile internal/cpu=/usr/local/go/pkg/linux_amd64/internal/cpu.a
packagefile runtime/internal/atomic=/usr/local/go/pkg/linux_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/local/go/pkg/linux_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/local/go/pkg/linux_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/usr/local/go/pkg/linux_amd64/internal/reflectlite.a
packagefile sort=/usr/local/go/pkg/linux_amd64/sort.a
packagefile math/bits=/usr/local/go/pkg/linux_amd64/math/bits.a
packagefile internal/oserror=/usr/local/go/pkg/linux_amd64/internal/oserror.a
packagefile internal/poll=/usr/local/go/pkg/linux_amd64/internal/poll.a
packagefile internal/syscall/execenv=/usr/local/go/pkg/linux_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/usr/local/go/pkg/linux_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/local/go/pkg/linux_amd64/internal/testlog.a
packagefile sync/atomic=/usr/local/go/pkg/linux_amd64/sync/atomic.a
packagefile syscall=/usr/local/go/pkg/linux_amd64/syscall.a
packagefile time=/usr/local/go/pkg/linux_amd64/time.a
packagefile internal/unsafeheader=/usr/local/go/pkg/linux_amd64/internal/unsafeheader.a
packagefile unicode=/usr/local/go/pkg/linux_amd64/unicode.a
packagefile internal/race=/usr/local/go/pkg/linux_amd64/internal/race.a
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/b001/exe/main -importcfg $WORK/b001/importcfg.link -s -w -buildmode=exe -buildid=Tco5MYLftP_hRMqPbBlP/8a1GjSoYKC5J7O4qeuoq/aCesrJqnn4SnXX_xE7w3/Tco5MYLftP_hRMqPbBlP -extld=gcc /home/ed/.cache/go-build/98/98abe4a0683ebb0296e3931266e333ae8ac23cce798fc1222addf5c713f66f3c-d
$WORK/b001/exe/main
hello Go
```

However, in this repo we're going to use `Makefile` to be able to use the make tool to tell Go what to do and how to run or build. If you are in the directory ou intend to execute you can simply type out `make run` and if you

## variables, output and basic formatting with fmt

fmt (short for format, often pronounced as "foomt") is the formatting package for Go that facilitates simple console IO as well as more complex formatting operations. A good early example of file level imports, where you have a multiple file program and multiple files with formatting demands you should import "fmt" in each.

```go
package main

import "fmt"

/*
	Package fmt:
	1) implements formatted I/O with functions analogous to the C language
	printf and scanf functions. It's used mainly to print to the stdout.
	2) fmt.Println() writes to standard output. Spaces are always added
	between operands and a newline is appended.
	3) fmt.Printf() prints out to stdout according to a format specifier
	called verb. It doesn't add a newline (\n)
		VERBS:
			%d -> decimal
			%f -> float
			%s -> string
			%q -> double-quoted string
			%v -> value (any)
			%#v -> a Go-syntax representation of the value
			%T -> value Type
			%t -> bool (true or false)
			%p -> pointer (address in base 16, with leading 0x)
			%c -> char (rune) represented by the corresponding Unicode code point
*/

func main() {
	var message string
	message = "Hello from Go"

	fmt.Println(message)

	x, y, z := 100, -100, .100

	fmt.Printf("x: %v  %T", x, x)
	fmt.Printf("y: %v  %T", y, y)
	fmt.Printf("z: %v  %T", z, z)
}
```

Above we see a couple of different variable declaration techniques. We're going to see more of that in the dedicated code sample [here](). It's entirely possible to see multiple types of declaration in a single large program where many authors are present and where code review does not forbid specific declaration types.

- long form `var varname datatype = value`
- short form (walrus operator) `varname := value`
- inferred `var varname = value`

- `var value int` initialized with 0
- `var price float64` initialized with 0.0
- `var name string` initialized with empty string -> ""
- `var done bool` initialized with false

## Rules of assignment

- long form: `var varname type = value` or `var score float64 = 6.5`
- short form: `varname := value` or `score := 6.5`.
- inferred form: `var varname = value` or `var x = 6.5`
- All literals are untyped constants, if numerically outside the binds of the intended type you will have an overflow complication error.

## Rules of constants

- You cannot change a constant.
- Constants belong to compile time, they cannot be initlaised at runtime.
  - One caveat being built-in functions such as len()
- You cannot use a variable to initialise a constant.

## Operators in go

- Arithmetic operators

  - `+` sum
  - `-` difference
  - `*` product
  - `/` quotient
  - `%` remainder

- Assignment operators

  - `=` (simple assignment)
  - `+=` (increment assignment)
  - `-=` (decrement assignment)
  - `*=` (multiplication assignment)
  - `/=` (division assignment)
  - `%=` (modulus assignment)

- Comparison operators

  - `==` equal values
  - `!=` not equal
  - `>` left operand is greater than right operand
  - `<` left operand is less than right operand
  - `>=` left operand is greater than or equal to right operand
  - `<=` left operand is less than or equal to right operand

- Logical operators
  - `&&` logical and
  - `||` logical or
  - `!` logical negation
