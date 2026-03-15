package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestScanDir_FindsMatchingFiles(t *testing.T) {
	tmpDir := t.TempDir()

	files := []string{
		filepath.Join(tmpDir, "test1.txt"),
		filepath.Join(tmpDir, "test2.txt"),
		filepath.Join(tmpDir, "test3.jpg"),
	}

	for _, f := range files {
		err := os.WriteFile(f, []byte("test"), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	oldTime := time.Now().Add(-48 * time.Hour)
	for _, f := range files[:2] { // только .txt файлы
		os.Chtimes(f, oldTime, oldTime)
	}

	result := scanDir(tmpDir, []string{"*.txt"}, time.Now().Add(-24*time.Hour))

	if len(result) != 2 {
		t.Errorf("expected 2 files, got %d", len(result))
	}
}

func TestScanDir_DateFilter(t *testing.T) {
	tmp := t.TempDir()

	oldFile := filepath.Join(tmp, "old.txt")
	newFile := filepath.Join(tmp, "new.txt")

	os.WriteFile(oldFile, []byte("test"), 0644)
	os.WriteFile(newFile, []byte("test"), 0644)

	oldTime := time.Now().AddDate(0, 0, -5)

	os.Chtimes(oldFile, oldTime, oldTime)

	patterns := []string{"*.txt"}

	cutoff := time.Now().AddDate(0, 0, -2)

	result := scanDir(tmp, patterns, cutoff)

	if len(result) != 1 {
		t.Fatalf("expected 1 file, got %d", len(result))
	}

	if filepath.Base(result[0]) != "old.txt" {
		t.Fatal("wrong file selected")
	}
}

func TestCleanup_MovesFile(t *testing.T) {
	tmp := t.TempDir()

	trash := filepath.Join(tmp, ".Trash")
	os.Mkdir(trash, 0755)

	file := filepath.Join(tmp, "test.txt")
	os.WriteFile(file, []byte("data"), 0644)

	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tmp)
	defer os.Setenv("HOME", oldHome)

	cleanup([]string{file})

	_, err := os.Stat(filepath.Join(trash, "test.txt"))
	if err != nil {
		t.Fatal("file not moved to trash")
	}
}

func TestSearch_MultipleDirs(t *testing.T) {
	home := t.TempDir()

	downloads := filepath.Join(home, "Downloads")
	desktop := filepath.Join(home, "Desktop")

	os.Mkdir(downloads, 0755)
	os.Mkdir(desktop, 0755)

	os.WriteFile(filepath.Join(downloads, "file.txt"), []byte("x"), 0644)
	os.WriteFile(filepath.Join(desktop, "file.jpg"), []byte("x"), 0644)

	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", home)
	defer os.Setenv("HOME", oldHome)

	result := search([]string{"~/Downloads", "~/Desktop"}, []string{"*.txt"}, 0)

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
}
