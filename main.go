package main

func main() {
	pathDirs := []string{"~/Downloads", "~/Desktop"}
	filePatterns := []string{"*.torrent", "*.dmg", "*.web*", "*.json", "*.srt", "*.png", "*.jpg", "*.csv", "*.mp3", "*.txt", "*.md",
		"*.lua", "*.apkg", "*.HEIC", "*.heic", "*.MOV", "*.mkv", "*.epub", "*.mobi", "*.html", "*.otg", "*.obg",
	}
	//days := -1
	days := -10

	cleanup(search(pathDirs, filePatterns, days))
}
