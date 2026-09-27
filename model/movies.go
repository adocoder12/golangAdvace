package model

import "strings"

type Movie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
