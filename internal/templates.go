package internal

import (
	"embed"
	"todox/internal/app/config"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed **/*.html **/**/*.html
	tmpls embed.FS

	Templates = mdfs.New(tmpls, "internal", config.Environment)
)
