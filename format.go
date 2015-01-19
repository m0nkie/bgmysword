package main

import (
	"bytes"
	"unicode"
)

var (
	buffer bytes.Buffer
)

func str(slice []string) string {
	for _, s := range slice {
		buffer.WriteString(s)
	}
	concatenated := buffer.String()
	buffer.Reset()
	return concatenated
}

func slc(args ...string) []string {
	return args
}

func combine(args ...string) string {
	return str(args)
}

func appendCombine(slice []string, args ...string) (combined string) {
	appendedSlice := append(slice, args...)
	combined = str(appendedSlice)
	return
}

func strLast(s string) string {
	last := len(s) - 1
	return string([]rune(s)[last])
}

func suffixSpace(s string) string {
	for i, r := range s {
		if unicode.IsSpace(r) {
			i++
			return s[i:]
		}
	}
	return s
}

func suffixHyphenSecond(s string) string {
	var foundFirst bool
	ss := []byte(s)
	for i, b := range ss {
		if b == 45 {
			switch foundFirst {
			case false:
				foundFirst = true
			case true:
				i++
				return string(ss[i:])
			}
		}
	}
	return s
}

func prefixHyphenSecond(s string) string {
	var foundFirst bool
	ss := []byte(s)
	for i, b := range ss {
		if b == 45 {
			switch foundFirst {
			case false:
				foundFirst = true
			case true:
				i++
				return string(ss[:i])
			}
		}
	}
	return s
}
