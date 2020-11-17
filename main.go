package main

import (
	"bufio"
	"fmt"
	"os"
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
	thisIsFUnny := 1
	funny := "qwe"
	ic.ConfigureArgToStringFunction(toString)
	ic.ConfigureIncludeContext(true)
	ic.Ic(thisIsFUnny, funny)
	ic.Ic(1)
	ic.Ic(add(123))
	test()
}
