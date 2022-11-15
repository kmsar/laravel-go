package LocalFilesystem

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// IsExecutable tells whether the filename is executable
func IsExecutable(filename string) bool {
	if runtime.GOOS == "windows" {
		if !IsReadable(filename) {
			return false
		}
		return strings.ToUpper(filepath.Ext(filename)) == ".EXE"
	}
	return syscall.Access(filename, 0x1) == nil
}

// IsReadable tells whether a file exists and is readable
func IsReadable(filename string) bool {
	if runtime.GOOS == "windows" {
		_, err := os.Stat(filename)

		return err == nil
	}
	return syscall.Access(filename, 0x4) == nil
}

// IsWritable tells whether the filename is writable
func IsWritable(filename string) bool {
	if runtime.GOOS == "windows" {
		fi, err := os.Stat(filename)
		if err != nil {
			return false
		}

		return fi.Mode() == 0666
	}
	return syscall.Access(filename, 0x2) == nil
}
