package fs

import (
	"io"
	"io/fs"
	"os"
	"strings"
)

// FS wraps a directory and an embed FS that are expected to have the same contents.
// it prioritizes the directory FS and falls back to the embedded FS if the file cannot
// be found on disk. This is useful during development or when deploying with
// assets not embedded in the binary.
type Fallback struct {
	dir string

	embed fs.FS
	dirFs fs.FS
}

// NewFallbackFS returns a new FS that wraps the given directory and embedded FS.
// the embed.FS is expected to embed the same files as the directory FS.
func NewFallback(embed fs.FS, dir string) Fallback {
	// If the directory is empty, use the current working directory.
	if dir == "" {
		pwd, _ := os.Getwd()
		dir = pwd
	}

	return Fallback{
		embed: embed,
		dirFs: os.DirFS(dir),
		dir:   dir,
	}
}

// Open opens the named file.
//
// When Open returns an error, it should be of type *PathError with the Op
// field set to "open", the Path field set to name, and the Err field
// describing the problem.
//
// Open should reject attempts to open names that do not satisfy
// ValidPath(name), returning a *PathError with Err set to ErrInvalid or
// ErrNotExist.
func (f Fallback) Open(name string) (file fs.File, err error) {
	env := "development"
	if e := os.Getenv("GO_ENV"); e != "" {
		env = e
	}

	switch env {
	case "development":
		nname := strings.TrimLeft(name, f.dir)
		file, err = f.dirFs.Open(nname)
		if err == nil {
			return
		}

		fallthrough
	case "production":
		file, err = f.embed.Open(name)
	}

	return file, err
}

// ReadFile reads the named file from the file system fs and returns its contents.
// It uses the custom Open method to open the file.
func (f Fallback) ReadFile(name string) ([]byte, error) {
	x, err := f.Open(name)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(x)
}
