package icecream

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/shurcooL/go/reflectsource"
)

var prefixString = "ic| "
var outputFunction = reflect.ValueOf(os.Stderr.WriteString)
var argToStringFunction reflect.Value
var includeContext = false

func printMsg(msg interface{}) {
	rf := outputFunction
	s := reflect.ValueOf(msg)
	if rf.Kind() == reflect.Func {
		rf.Call([]reflect.Value{s})
	}
}

func Ic(values ...interface{}) {
	var msg string
	line := 0
	pc, filename, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	pwd, _ := os.Getwd()
	relFilename, _ := filepath.Rel(pwd, filename)
	raf := argToStringFunction
	rafk := raf.Kind()
	if ok {
		lenOfValues := len(values)
		if lenOfValues > 0 { // ? print value
			if includeContext { // ? print prefix
				msg = fmt.Sprintf("%s%s:%d in %s()- ", prefixString, relFilename, line, funcName)
			} else {
				msg = prefixString
			}
			printMsg(msg)
			for i, v := range values { // ? print value
				vName := reflectsource.GetParentArgExprAsString(uint32(i))
				fmtV := fmt.Sprintf("%#v", v)
				if vName == fmtV { // ? vName the same as value, just print one
					printMsg(fmtV)
				} else {
					msg = fmt.Sprintf("%s: ", vName)
					printMsg(msg)
					if rafk == reflect.Invalid || rafk == reflect.Bool { // ? argToStringFunction default is fmt.Sprintf
						printMsg(fmt.Sprintf("%#v", v))
					} else if rafk == reflect.Func { // ? call custom argToStringFunction
						results := raf.Call([]reflect.Value{reflect.ValueOf(v)})
						printMsg(results[0].Interface())
					}
				}
				if i < lenOfValues-1 { // ? print split character
					printMsg(", ")
				}
			}
			printMsg("\n")
		} else { // ? print line if value is nil
			msg = fmt.Sprintf("%s%s:%d in %s()\n", prefixString, relFilename, line, funcName)
			printMsg(msg)
		}
	}
}

func ConfigurePrefix(prefix string) {
	prefixString = prefix
}

// f func(s interface{})
func ConfigureOutputFunction(f interface{}) bool {
	rf := reflect.ValueOf(f)
	if rf.Kind() == reflect.Func {
		outputFunction = rf
		return true
	}
	return false
}

// f func(v interface{}) interface{}
func ConfigureArgToStringFunction(f interface{}) bool {
	rf := reflect.ValueOf(f)
	if rf.Kind() == reflect.Func {
		argToStringFunction = rf
		return true
	}
	return false
}

func ConfigureIncludeContext(boolean bool) {
	includeContext = boolean
}

func ResetPrefix() {
	prefixString = "ic| "
}

func ResetOutputFunction() {
	outputFunction = reflect.ValueOf(os.Stderr.WriteString)
}

func ResetArgToStringFunction() {
	argToStringFunction = reflect.ValueOf(false)
}

func ResetIncludeContext() {
	includeContext = false
}
