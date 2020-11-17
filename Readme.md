# IceCream-Go

A Go port of Python's [IceCream](https://github.com/gruns/icecream).

## Usage

```go
package main

import (
	"fmt"
	ic "github.com/WAY29/icecream-go/icecream"
)
func foo(a int) int {
	return a + 333
}


func bar() {
	ic.Ic()
}

func main() {
ic.Ic(foo(123));
// Outputs:
// ic| foo(123): 456

ic.Ic(1 + 5);
// Outputs:
// ic| 1 + 5: 6

ic.Ic(foo(123), 1 + 5);
// Outputs:
// ic| foo(123): 456, 1 + 5: 6
bar();
}
// Outputs:
// ic| main.go:12 in main.bar()
```

## Depends
- [reflectsource][github.com/shurcooL/go/reflectsource]


## Installation

```
go get -u "github.com/WAY29/icecream-go/icecream"
```

## Configuration

If you want to change the prefix of the output, you can call `icecream.ConfigurePrefix("Hello| ")` (by default the prefix is `ic| `).

If you want to change how the result is outputted, you can call `icecream.ConfigureOutputFunction()`.
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
}
```

If you want to change how the value is outputted, you can call `icecream.ConfigureArgToStringFunction()`.
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

If you want to add call's filename, line number, and parent function to ic's output, you can call `icecream.ConfigureIncludeContext(true)`.
```go
func main() {
    ic.ConfigureIncludeContext(true)
    ic.Ic(1, "asd")
}
```