package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const PasswordSalt = "AQI"

func MD5Password(password string) string {
	return MD5(PasswordSalt + password)
}

func MD5(input string) string {
	h := md5.Sum([]byte(input))
	return hex.EncodeToString(h[:])
}

var suffixToType = map[string]string{
	"jpg": "IMG", "jpeg": "IMG", "png": "IMG", "gif": "IMG", "bmp": "IMG",
	"tiff": "IMG", "svg": "IMG", "ico": "IMG", "webp": "IMG", "heic": "IMG",
	"psd": "IMG", "ai": "IMG",
	"mp4": "VIDEO", "avi": "VIDEO", "mkv": "VIDEO", "flv": "VIDEO", "mov": "VIDEO",
	"wmv": "VIDEO", "mpeg": "VIDEO", "rmvb": "VIDEO", "3gp": "VIDEO", "webm": "VIDEO",
	"m4v": "VIDEO", "ts": "VIDEO",
	"mp3": "AUDIO", "wav": "AUDIO", "flac": "AUDIO", "aac": "AUDIO", "ogg": "AUDIO",
	"wma": "AUDIO", "m4a": "AUDIO",
	"doc": "DOC", "docx": "DOC", "pdf": "PDF", "txt": "TXT", "ppt": "PPT",
	"pptx": "PPT", "xls": "EXCEL", "xlsx": "EXCEL", "odt": "DOC", "rtf": "DOC",
	"csv": "CSV", "md": "TXT", "epub": "DOC", "mobi": "DOC", "tex": "DOC",
	"zip": "COMPRESS", "rar": "COMPRESS", "7z": "COMPRESS", "tar": "COMPRESS",
	"gz": "COMPRESS", "bz2": "COMPRESS", "xz": "COMPRESS", "iso": "COMPRESS",
	"exe": "CODE", "bat": "CODE", "sh": "CODE", "apk": "CODE",
}

var imageExts = map[string]bool{
	"jpg": true, "jpeg": true, "png": true, "gif": true, "bmp": true,
	"tiff": true, "svg": true, "ico": true, "webp": true, "heic": true,
}

var extToMIME = map[string]string{
	"jpg":  "image/jpeg", "jpeg": "image/jpeg", "png": "image/png",
	"gif":  "image/gif",  "bmp":  "image/bmp",  "tiff": "image/tiff",
	"svg":  "image/svg+xml", "ico": "image/x-icon", "webp": "image/webp",
	"heic": "image/heic", "pdf":  "application/pdf",
	"mp4":  "video/mp4", "webm": "video/webm", "avi": "video/x-msvideo",
	"mp3":  "audio/mpeg", "wav":  "audio/wav", "ogg": "audio/ogg",
}

func GetMIMEType(suffix string) string {
	suffix = strings.ToLower(strings.TrimPrefix(suffix, "."))
	if mime, ok := extToMIME[suffix]; ok {
		return mime
	}
	return "application/octet-stream"
}

func IsImage(suffix string) bool {
	suffix = strings.ToLower(strings.TrimPrefix(suffix, "."))
	return imageExts[suffix]
}

func DetectFileType(suffix string) string {
	suffix = strings.ToLower(strings.TrimPrefix(suffix, "."))
	if t, ok := suffixToType[suffix]; ok {
		return t
	}
	return "COMMON"
}

func RandomString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[utilRand()%len(letters)]
	}
	return string(b)
}

var seed int64

func utilRand() int {
	seed++
	return int(seed)
}
