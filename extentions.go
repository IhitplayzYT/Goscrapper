package main

// Video file extensions
var VideoExt = []string{
	".mp4", ".mkv", ".avi", ".mov", ".webm",
	".flv", ".wmv", ".mpeg", ".mpg", ".3gp",
}

// Image file extensions
var ImageExt = []string{
	".jpg", ".jpeg", ".png", ".gif", ".bmp",
	".webp", ".svg", ".tiff", ".ico", ".avif",
}

// Audio file extensions
var AudioExt = []string{
	".mp3", ".wav", ".aac", ".ogg", ".flac",
	".m4a", ".wma", ".opus", ".aiff",
}

// Document / file extensions
var DocExt = []string{
	".pdf", ".doc", ".docx", ".xls", ".xlsx",
	".ppt", ".pptx", ".txt", ".csv", ".rtf",
	".odt", ".ods", ".odp", ".md",
}

// Archive / compressed files
var ArchiveExt = []string{
	".zip", ".rar", ".7z", ".tar", ".gz",
	".bz2", ".xz", ".tgz",
}

// Web-related files (often scraped)
var WebExt = []string{
	".html", ".htm", ".css", ".js", ".json",
	".xml", ".wasm",
}

// Fonts (useful for asset scraping)
var FontExt = []string{
	".woff", ".woff2", ".ttf", ".otf", ".eot",
}

// Misc useful extensions
var MiscExt = []string{
	".apk", ".exe", ".dmg", ".iso",
	".torrent", ".log",
}

var assetExt = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
	".mp4": true, ".webm": true, ".mp3": true, ".wav": true, ".ogg": true,
	".pdf": true, ".zip": true, ".tar": true, ".gz": true,
}
