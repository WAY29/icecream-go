package icecream

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/shurcooL/go/reflectsource"
)

var prefixString = "ic| "
var outputFunction = reflect.ValueOf(os.Stderr.WriteString)
var argToStringFunction reflect.Value
var argNameFormatterFunc reflect.Value
var includeContext = false
var disableOutput = false

func printMsg(msg interface{}) {
	rf := outputFunction
	s := reflect.ValueOf(msg)
	if rf.Kind() == reflect.Func && !disableOutput {
		rf.Call([]reflect.Value{s})
	}
}

func formatValue(v interface{}) interface{} {
	raf := argToStringFunction
	rafk := raf.Kind()
	if rafk == reflect.Invalid || rafk == reflect.Bool { // ? argToStringFunction default is fmt.Sprintf
		return fmt.Sprintf("%#v", v)
	} else if rafk == reflect.Func { // ? call custom argToStringFunction
		results := raf.Call([]reflect.Value{reflect.ValueOf(v)})
		return results[0].Interface()
	}

	return nil
}

func formatArgName(argName string) string {
	raf := argNameFormatterFunc
	if raf.Kind() == reflect.Func { // call custom argNameFormatterFunc
		results := raf.Call([]reflect.Value{reflect.ValueOf(argName)})
		if ret, ok := results[0].Interface().(string); ok {
			return ret
		}
	}

	return argName
}

func Ic(values ...interface{}) []interface{} {
	var (
		msg         string
		returnValue = make([]interface{}, 0, len(values))
	)

	line := 0
	pc, filename, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	pwd, _ := os.Getwd()
	relFilename, _ := filepath.Rel(pwd, filename)
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
					printMsg(formatValue(v))
				} else {
					msg = fmt.Sprintf("%s: ", formatArgName(vName))
					printMsg(msg)
					printMsg(formatValue(v))
				}
				if i < lenOfValues-1 { // ? print split character
					printMsg(", ")
				}

				returnValue = append(returnValue, v) // ? add value in returnValue
			}
			printMsg("\n")
		} else { // ? print line if value is nil
			msg = fmt.Sprintf("%s%s:%d in %s()\n", prefixString, relFilename, line, funcName)
			printMsg(msg)
		}
	}

	return returnValue
}

func Format(values ...interface{}) string {
	var (
		returnMsg string = ""
	)

	line := 0
	pc, filename, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	pwd, _ := os.Getwd()
	relFilename, _ := filepath.Rel(pwd, filename)
	if ok {
		lenOfValues := len(values)
		if lenOfValues > 0 { // ? print value
			if includeContext { // ? print prefix
				returnMsg += fmt.Sprintf("%s%s:%d in %s()- ", prefixString, relFilename, line, funcName)
			} else {
				returnMsg += prefixString
			}
			for i, v := range values { // ? print value
				vName := reflectsource.GetParentArgExprAsString(uint32(i))
				fmtV := fmt.Sprintf("%#v", v)
				if vName == fmtV { // ? vName the same as value, just print one
					returnMsg += fmt.Sprintf("%v", formatValue(v))
				} else {
					returnMsg += fmt.Sprintf("%s: ", vName)
					returnMsg += fmt.Sprintf("%v", formatValue(v))
				}
				if i < lenOfValues-1 { // ? print split character
					returnMsg += ", "
				}

			}
		} else { // ? print line if value is nil
			returnMsg = fmt.Sprintf("%s%s:%d in %s()\n", prefixString, relFilename, line, funcName)
		}
	}

	return returnMsg
}

func ConfigurePrefix(prefix string) {
	prefixString = prefix
}

// ConfigureOutputFunction modifies output that writes the final msg into w.
// w should be io.StringWriter or function with string arg,
// the return value indicates whether w is legal
func ConfigureOutputFunction(w interface{}) bool {
	if strWriter, ok := w.(io.StringWriter); ok {
		outputFunction = reflect.ValueOf(strWriter.WriteString)
		return true
	}
	rf := reflect.ValueOf(w)
	if rf.Kind() == reflect.Func {
		outputFunction = rf
		return true
	}
	return false
}

func ConfigureArgToStringFunction(f interface{}) bool {
	rf := reflect.ValueOf(f)
	if rf.Kind() == reflect.Func {
		argToStringFunction = rf
		return true
	}
	return false
}

// ConfigureArgNameFormatterFunc config arg name formatter.
// f should be `func(string) string`,
// the return value indicates whether w is legal.
func ConfigureArgNameFormatterFunc(f func(string) string) bool {
	rf := reflect.ValueOf(f)
	if rf.Kind() == reflect.Func {
		argNameFormatterFunc = rf
		return true
	}
	return false
}

func Disable() {
	disableOutput = true
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

func Enable() {
	disableOutput = false
}
