![](https://github.com/irisida/golang/blob/master/assets/freegopher.png)

# Learning Go

How do you start learning a programming language? ..or a new programming language? That's pretty much what I'm going to try to cover in this repository. It will be a mix of code snippets, little trivial mini-projects and readmes. I'll probably tart some of it out as blog posts with snatched and butchered ideas twisted to my own amusement.

So, I'm going to covering Go. or Golang as its sometimes called, mistakenly. Why go, you might ask? well to be honest because I'm trying to learn it myself and because in this content obsessed world I thought I'd have a pop at writing about how I'm going about it. Background wise I've been programming since before the internet. Once upon a time I did a lot of it, these days I mostly have other people working on my teams doing it and that relegates me to the evening hobbyist. The interested party and casual participant. So, of you're reading along you should manage your expectations that this is not going to be some revolutionary programming guide with a hot spin on learning techniques, at best you might get a sense of how an old dog _onboards_ new (Are they new?) tricks.

Right, let's go then... Go, a product of lofty minds at Google who fancied tarting up C/C++ of some of the issues as well as addressing the need for a modern language that deals with multicore processors and has a great concurrency model. Such promise!

## Installation and setup common best practice

- Install golang for your system from [here](https://golang.org/)
- Set the GOPATH to `Users/username/go`
- Within the GOPATH directory create three subdirectories, `bin`, `pkg` & `src`
- within the `src` directory create a `github.com` directory and your GH username as a child of that. From within that child directory you will create and store your Go projects.

## OK, hit me with the day 1 jazz...

So, we're going to take some assumptions on the basis that it's `strongly typed` and has a basis in C and C-like languages. By this the older cynic can deduce the intention of the project was always to move beyond an internal tool, to keep tried, tested and battle-hardened concepts from an absolute warhorse of a language. So we may also gather that they genuinely are looking to have ease at the core of the language and concession #1 was to use/re-use syntactic constructs from the most popular languages of the last 40+ years.

Let's see the hello world nonsense then....

![](https://github.com/irisida/golang/blob/master/assets/day1/day01.001.png)

OK, so the meat we're covering in that non-HelloWorld example is:

- we can see that a Go program follows a package/import/function model

Like every other successful language we have to think about statements and expressions and like every other language:

- Statements are operations and may contain one or many expressions. Instructions to do something. eg. fmt.Println()
- expressions are returned values, here it is "This is a go program!"

Let's see a cheeky pic of that in action.

![](https://github.com/irisida/golang/blob/master/assets/day1/day01.002.png)

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

## Basic formatting with fmt

fmt (short for format, often pronounced as "foomt") is the formatting package for Go that facilitates simple console IO as well as more complex formatting operations. A good early example of file level imports, where you have a multiple file program and multiple files with formatting demands you should import "fmt" in each.

![](https://github.com/irisida/golang/blob/master/assets/day1/day01.003.png)

## Variables in go

covers declaration and assignment forms.

- long form `var varname datatype = value`
- short form (walrus operator) `varname := value`
- inferred `var varname = value`

- `var value int` initialized with 0
- `var price float64` initialized with 0.0
- `var name string` initialized with empty string -> ""
- `var done bool` initialized with false

![](https://github.com/irisida/golang/blob/master/assets/day1/day01.004.png)

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
