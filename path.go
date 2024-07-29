package main

import "strings"

func normalizePath(path string) string {
	return strings.TrimPrefix(path, "./")
}
