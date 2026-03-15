package main

func main() {
	pathDirs := []string{"~/Downloads", "~/Desktop"}
	typeFiles := []string{
		"*.dmg",
		"*.web*",
		"*.json",
		"*.srt",
		"*.png",
		"*.jpg",
		"*.csv",
		"*.mp3",
		"*.txt",
		"*.md",
		"*.lua",
		"*.apkg",
		"*.HEIC",
		"*.heic",
		"*.MOV",
		"*.mkv",
		"*.epub",
		"*.mobi",
		"*.html",
		"*.otg",
		"*.obg",
	}

	search(pathDirs, typeFiles)
	cleanup()
}
