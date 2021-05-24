package main

import (
	"fmt"
	"strings"
)

func TrimStr(value string) string {
	// the original value is not updated
	// instead a new memory is allocated
	// as strings in go are readonly
	value = strings.TrimSpace(value)
	return value
}

func main() {
	name := "   Amir Saleem   "
	trimmedName := TrimStr(name)
	fmt.Println(name)
	fmt.Println(trimmedName)
}
