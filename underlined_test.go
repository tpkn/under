package under

import (
	"fmt"
	"testing"
)

type data struct {
	title string
	symbol string
	with_text bool
}

var input = []data{
	data {"Solid", "solid", true},
	data {"Dashed", "dashed", true},
	data {"Dotted", "dotted", true},
	data {"Custom line style", "+", true},
	data {"Just a line", "solid", false},
	data {"", "solid", true},
}

func TestUnderlined(t *testing.T) {
	for _, d := range input {
		result := Lined(d.title, d.symbol, d.with_text)
		fmt.Println(result)
		fmt.Println("")
	}
}
