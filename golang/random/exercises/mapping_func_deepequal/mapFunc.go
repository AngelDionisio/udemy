package main

import (
	"strings"
)

// MapFunc applies operation to each item in list
func MapFunc(list []string, operation func(string) string) []string {
	if list == nil || len(list) < 1 {
		return nil
	}

	var transformation []string
	for _, v := range list {
		transformation = append(transformation, operation(v))
	}

	return transformation
}

// UpperCase uppercases all runes in a string
func UpperCase(s string) string {
	return strings.ToUpper(s)
}
