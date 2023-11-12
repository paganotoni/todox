// public package contains the public folder that will be served by the server
// it embeds the files that are part of it so that the application serves those
// files publicly.
package public

import (
	"embed"
	"path/filepath"
	"todox/internal/app/config"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed *
	files embed.FS

	// Folder is a mdfs instance that contains all the
	// files in the public folder.
	Folder = mdfs.New(files, filepath.Join("internal", "app", "public"), config.Environment)
)
