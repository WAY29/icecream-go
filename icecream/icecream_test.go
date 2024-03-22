package icecream

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConfigureOutputFunction(t *testing.T) {
	builder := strings.Builder{}
	ok := ConfigureOutputFunction(&builder)
	assert.Equal(t, true, ok)

	expected := "111"
	printMsg(expected)
	actual := builder.String()
	assert.Equal(t, expected, actual)
}

func TestIc(t *testing.T) {
	builder := strings.Builder{}
	ok := ConfigureOutputFunction(&builder)
	assert.Equal(t, true, ok)
	ConfigureArgNameFormatterFunc(func(name string) string {
		return "\u001B[36m" + name + "\u001B[0m"
	})

	str := "hello"
	Ic(str)
	actual := builder.String()
	assert.Equal(t, "ic| \u001B[36mstr\u001B[0m: \"hello\"\n", actual)
}
