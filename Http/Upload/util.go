package Upload

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var contentTypeByExtension = map[string]string{
	".jsonld": "application/ld+json",
	".json":   "application/json",
	".js":     "text/javascript",
	".mjs":    "text/javascript",
	".css":    "text/css",
}

// GetFileExtension returns the last part of a file name.
// If the file doesn't have an extension, returns an empty string.
func GetFileExtension(file string) string {
	index := strings.LastIndex(file, ".")
	if index == -1 {
		return ""
	}
	return file[index+1:]
}

// GetMIMEType get the mime type and size of the given file.
//
// If the file cannot be opened, panics. You should check if the
// file exists, using "fsutil.FileExists()"", before calling this function.
func GetMIMEType(file string) (string, int64) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buffer := make([]byte, 512)

	_, errRead := f.Read(buffer)
	if errRead != nil {
		panic(errRead)
	}

	stat, errStat := f.Stat()
	if errStat != nil {
		panic(errStat)
	}

	contentType := http.DetectContentType(buffer)

	if strings.HasPrefix(contentType, "application/octet-stream") || strings.HasPrefix(contentType, "text/plain") {
		for ext, t := range contentTypeByExtension {
			if strings.HasSuffix(file, ext) {
				tmp := t
				if i := strings.Index(contentType, ";"); i != -1 {
					tmp = t + contentType[i:]
				}
				contentType = tmp
				break
			}
		}
	}

	return contentType, stat.Size()
}

// FileExists returns true if the file at the given path exists and is readable.
// Returns false if the given file is a directory.
func FileExists(file string) bool {
	if stats, err := os.Stat(file); err == nil {
		return !stats.IsDir()
	}
	return false
}

// IsDirectory returns true if the file at the given path exists, is a directory and is readable.
func IsDirectory(path string) bool {
	if stats, err := os.Stat(path); err == nil {
		return stats.IsDir()
	}
	return false
}

// Delete the file at the given path.
//
// To avoid panics, you should check if the file exists.
func Delete(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func timestampFileName(name string) string {
	var prefix string
	var extension string
	index := strings.LastIndex(name, ".")
	if index == -1 {
		prefix = name
		extension = ""
	} else {
		prefix = name[:index]
		extension = name[index:]
	}
	return prefix + "-" + strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10) + extension
}
