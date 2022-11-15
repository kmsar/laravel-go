package Upload

type IFile interface {
	// Path Get the fully qualified path to the file.
	Path() string

	// GetExtension  the file's extension.
	GetExtension() string

	// HashName   a filename for the file.
	HashName() string
}
