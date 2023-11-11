package public

import (
	"embed"
	"paganotoni/todox/internal/config"
	"path/filepath"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed *
	files embed.FS

	// Folder is a mdfs instance that contains all the
	// files in the public folder.
	Folder = mdfs.New(
		files,
		filepath.Join("internal", "public"),
		config.Environment,
	)
)
