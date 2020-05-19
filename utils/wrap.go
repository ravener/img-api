package utils

import (
	"regexp"
	"strconv"
)

// Literally all the existing word wrap implementations suck.
// They split by spaces so "aaa aaa" works but "aaaaaa" wouldn't get word wrapped.
// That is not how real word wrapping works. It should still put them on a newline.
// But we should also respect words so as to not cut words.
// E.g if "hello world" exceeded then "world" should be in a newline but not trimmed like "wo\nrld"
func WordWrap(s string, limit int) []string {
	if len(s) < limit {
		return []string{s}
	}

	// Ok maybe regex is not the best solution but I can't think right now.
	// Plus I already did this in JavaScript before so I know it works.
	regex := regexp.MustCompile(".{1," + strconv.Itoa(limit) + "}(\\s|$)" + "|.{" + strconv.Itoa(limit) + "}|.+$")
	return regex.FindAllString(s, -1)
}
