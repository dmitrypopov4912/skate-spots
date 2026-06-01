package main

import(
	"strings"
)

func CleanSpotName (name string) string{
	goodName := strings.TrimSpace(name)
	goodName = strings.ToLower(goodName)
	goodName = strings.ReplaceAll(goodName, "#", "-")
	return goodName
}