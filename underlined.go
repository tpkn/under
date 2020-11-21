// Underlined, https://tpkn.me
package under

import (
	"strings"
	"unicode/utf8"
)

func LinedSolid(text string, with_text bool) string {
	return Lined(text, "solid", with_text)
}

func LinedDashed(text string, with_text bool) string {
	return Lined(text, "dashed", with_text)
}

func LinedDotted(text string, with_text bool) string {
	return Lined(text, "dotted", with_text)
}

func Lined(text, symbol string, with_text bool) string {
	
	// Predefined line styles
	switch symbol {
	case "solid": symbol = "─"
	case "dashed": symbol = "-"
	case "dotted": symbol = "."
	}
	
	under := make([]string, utf8.RuneCountInString(text))
	for i := range under {
		under[i] = symbol
	}
	
	result := strings.Join(under, "")
	
	if with_text {
		result = text + "\n" + result
	}
	
	return result
}
