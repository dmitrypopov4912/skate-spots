package main

import(
	"unicode/utf8"
)

func TruncateStringByBytes(s string, maxBytes int) string{
	bytes := len(s)
	if bytes <= maxBytes{
		return s
	}
	for maxBytes > 0 && !utf8.RuneStart(s[maxBytes]){
		maxBytes--
	}
	return s[:maxBytes]
}