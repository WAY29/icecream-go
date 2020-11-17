package main

import (
	"fmt"
	"reflect"

	ic "github.com/WAY29/icecream-go/icecream"
)

func add(a int) int {
	return a + 333
}

func test() {
	ic.Ic()
}

func toString(v interface{}) interface{} {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.String {
		return fmt.Sprintf("[!string %#v with length %d!]", v, len(v.(string)))
	}
	return fmt.Sprintf("%#v", v)
}

func main() {
	thisIsFUnny := 1
	funny := "qwe"

	ic.ConfigureIncludeContext(true)
	ic.Ic(thisIsFUnny, funny)
	ic.Ic(1)
	ic.Ic(add(123))
	test()
}
