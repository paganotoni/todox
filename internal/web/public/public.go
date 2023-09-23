package public

import (
	"embed"
	"path/filepath"

	"paganotoni/todox/internal/config"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed *.css
	files embed.FS

	// Folder is a mdfs instance that contains all the
	// files in the public folder.
	Folder = mdfs.New(
		files,
		filepath.Join("internal", "web", "public"),
		config.Environment,
	)
)
