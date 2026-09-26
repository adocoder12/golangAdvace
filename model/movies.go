package model

import "strings"

type Movie struct {
	ID          int
	Name        string
	Year        int
	Description string
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
