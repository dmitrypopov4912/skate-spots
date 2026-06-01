package main

import(
	"fmt"
	"strings"
)

func CleanSpotName (name string) string{
	googName := strings.TrimSpace(name)
	goodName = strings.ToLower(goodName)
	goodName = strings.ReplaceAll(goodName, "#", "-")
	return goodName
}