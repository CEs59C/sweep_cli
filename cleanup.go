package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func cleanup(files []string) {
	fmt.Println("--- Начинаю перемещение файлов ---")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	trashPath := filepath.Join(home, ".Trash")

	for _, oldPath := range files {
		fileName := filepath.Base(oldPath)

		newPath := filepath.Join(trashPath, fileName)

		err := os.Rename(oldPath, newPath)

		if err != nil {
			fmt.Printf("Не удалось переместить %s: %v\n", fileName, err)
		} else {
			fmt.Printf("Успешно: %s -> %s\n", fileName, trashPath)
		}
	}
}
