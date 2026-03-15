package main

import (
	"os"
	"strings"
	"time"
)

func search(pathDirs []string, typeFiles []string, days int) []string {
	var result []string
	home, _ := os.UserHomeDir()
	cutoff := time.Now().AddDate(0, 0, -days)

	for _, dir := range pathDirs {
		cleanPath := strings.Replace(dir, "~", home, 1)

		found := scanDir(cleanPath, typeFiles, cutoff)
		result = append(result, found...)
	}

	return result
}

func scanDir(currentPath string, typeFiles []string, cutoff time.Time) []string {
	var foundFiles []string

	return foundFiles
}
