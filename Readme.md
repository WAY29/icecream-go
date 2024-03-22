# IceCream-Go

A Go port of Python's [IceCream](https://github.com/gruns/icecream).

## Usage

```go
package main

import (
	. "github.com/WAY29/icecream-go/icecream"
)

func foo(a int) int {
	return a + 333
}


func bar() {
	Ic()
}

func main() {
	Ic(foo(123))
	// Outputs:
	// ic| foo(123): 456

	Ic(1 + 5)
	// Outputs:
	// ic| 1 + 5: 6

	Ic(foo(123), 1 + 5)
	// Outputs:
	// ic| foo(123): 456, 1 + 5: 6
	bar()
}
// Outputs:
// ic| main.go:12 in main.bar()
```

## Depends
- [reflectsource](https://github.com/shurcooL/go/tree/master/reflectsource)


## Installation

```
go get -u "github.com/WAY29/icecream-go/icecream"
```

## Import

```go
import ic "github.com/WAY29/icecream-go/icecream"
// or just import . "github.com/WAY29/icecream-go/icecream"
```

## Configuration

If you want to change the prefix of the output, you can call `icecream.ConfigurePrefix("Hello| ")` (by default the prefix is `ic| `).

If you want to change how the result is outputted, you can call `icecream.ConfigureOutputFunction(f)`. func may be type of `func(s interface{})`.
For example, if you want to log your messages to a file:
```go
func logfile(s string) {
	filePath := "log.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(s)
	write.Flush()
}
func main() {
	ic.ConfigureOutputFunction(logfile)
	ic.Ic(1, 2, 3)

	// output to stringBuilder
	builder := strings.Builder{}
	ic.ConfigureOutputFunction(&builder)
	ic.Ic("hello")
}
```

If you want to change how the value is outputted, you can call `icecream.ConfigureArgToStringFunction(f)`, func may be type of `func(v interface{}) interface{}`.
For example, if you want to print more detail about a string:
```go
func toString(v interface{}) interface{} {
    rv := reflect.ValueOf(v)
    if rv.Kind() == reflect.String {
        return fmt.Sprintf("[!string %#v with length %d!]", v, len(v.(string)))
    }
    return fmt.Sprintf("%#v", v)
}
func main() {
    s := "string"
    ic.ConfigureArgToStringFunction(toString)
    ic.Ic(s)
    ic.Ic("test")
}
```

If you want to change the value name is outputted, you can call `icecream.ConfigureArgNameFormatterFunc`.
For example, if you want to print value name with ansi color:
```go
func main() {
    s := "string"
    ic.ConfigureArgNameFormatterFunc(func(name string) string {
        return "\u001B[36m" + name + "\u001B[0m"
    })
    ic.Ic(s)
    ic.Ic("test")
}
```

If you want to add call's filename, line number, and parent function to ic's output, you can call `icecream.ConfigureIncludeContext(true)`.
```go
func main() {
    ic.ConfigureIncludeContext(true)
    ic.Ic(1, "asd")
}
```

If you want to reset configuration,  you can call `icecream.ResetPrefix()`,`icecream.ResetOutputFunction()`, `icecream.ResetArgToStringFunction()`,`icecream.ResetIncludeContext()` .

## Return Value
`Ic()` returns its arguments, so `Ic()` can easily be inserted into pre-existing code.

```go
func half(i interface{}) int {
    if ii, ok := i.(int); ok {
        return ii / 2
    }

    return -1
}


func main() {
    a := 6
    b := half(ic.Ic(a)[0])
    ic.Ic(b)
}
```
Prints
```
ic| a: 6
ic| b: 3
```

## Miscellaneous

`Format(...interface{})` is like `ic()` but the output is returned as a string instead of written to stderr.

```go
func main() {
    result := ic.Format("sup")
    fmt.Printf("%s", result)
}
```

Additionally, `Ic()`'s output can be entirely disabled, and later re-enabled, with `Disable()` and `Enable()` respectively.
```go
func main() {
    ic.Ic(1)
    ic.Disable()
    ic.Ic(2)
    ic.Enable()
    ic.Ic(3)
}
```
Prints
```
ic| 1
ic| 3
```