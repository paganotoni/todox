package internal

import (
	"embed"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed **/*.html *.html
	tmpls embed.FS

	Templates = mdfs.New(tmpls, "internal", Environment)
)
