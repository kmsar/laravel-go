package IFilesystem

import (
	"bufio"
	"io/fs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"time"
)

type FileVisibility int

// FileSystemProvider Get the filesystem with the given name and configuration information.
type FileSystemProvider func(name string, config Support.Fields) FileSystem

type File interface {
	fs.FileInfo

	// Read read a file.
	Read() []byte

	// ReadString read a file and return a string.
	ReadString() string

	// Disk get filesystem name.
	Disk() string
}

type FileSystemFactory interface {

	// Disk Get a filesystem implementation.
	Disk(disk string) FileSystem

	// Extend  the filesystem factory with the given name and filesystem provider.
	Extend(driver string, provider FileSystemProvider)

	FileSystem
}

type FileSystem interface {

	// Name Extract the file name from a file path.
	Name() string

	// Exists Determine if a file or directory exists.
	Exists(path string) bool

	// Get  the contents of a file.
	Get(path string) (string, error)

	// Read  a file.
	Read(path string) ([]byte, error)

	// ReadStream Retrieves a read-stream for a path.
	ReadStream(path string) (*bufio.Reader, error)

	// Put Create a file or update if exists.
	Put(path, contents string) error

	// WriteStream Write a new file using a stream.
	WriteStream(path string, contents string) error

	// GetVisibility get a file's visibility.
	GetVisibility(path string) FileVisibility

	// SetVisibility Set the visibility for a file.
	SetVisibility(path string, perm fs.FileMode) error

	// Prepend  to a file.
	Prepend(path, contents string) error

	// Append  to a file.
	Append(path, contents string) error

	// Delete the file at a given path.
	Delete(path string) error

	// Copy  a file to a new location.
	Copy(from, to string) error

	// Move  a file to a new location.
	Move(from, to string) error

	// Size get the file size of a given file.
	Size(path string) (int64, error)

	// LastModified Get the file's last modification time.
	LastModified(path string) (time.Time, error)

	// Files get an array of all files in a directory.
	Files(directory string) []File

	// AllFiles get all of the files from the given directory (recursive).
	AllFiles(directory string) []File

	// Directories get all of the directories within a given directory.
	Directories(directory string) []string

	// AllDirectories get all directories from a given directory (recursive).
	AllDirectories(directory string) []string

	// MakeDirectory Create a directory.
	MakeDirectory(path string) error

	// DeleteDirectory Recursively delete a directory.
	DeleteDirectory(directory string) error
}
