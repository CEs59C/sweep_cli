package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func search(pathDirs []string, typeFiles []string, days int) []string {
	var result []string
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cutoff := time.Now().AddDate(0, 0, -days)

	for _, dir := range pathDirs {
		var cleanPath string
		if strings.HasPrefix(dir, "~/") {
			cleanPath = filepath.Join(home, dir[2:])
		} else {
			cleanPath = dir
		}

		found := scanDir(cleanPath, typeFiles, cutoff)
		result = append(result, found...)
	}

	return result
}

func scanDir(currentPath string, typeFiles []string, cutoff time.Time) []string {
	var foundFiles []string

	contents, err := os.ReadDir(currentPath)
	if err != nil {
		return nil
	}

	for _, entry := range contents {
		fullPath := filepath.Join(currentPath, entry.Name())

		info, err := entry.Info()
		if err != nil {
			fmt.Println("[Error] Ошибка взятия информации:", err.Error())
			continue
		}
		//fmt.Println(fullPath, info.ModTime())

		// Проверка даты
		if info.ModTime().Before(cutoff) {
			// Проверка расширения
			for _, pattern := range typeFiles {
				matched, err := filepath.Match(strings.ToLower(pattern), strings.ToLower(entry.Name()))
				if err != nil {
					continue
				}
				if matched {
					foundFiles = append(foundFiles, fullPath)
					break
				}
			}
		}
	}

	return foundFiles
}
