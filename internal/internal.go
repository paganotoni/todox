package internal

import (
	"embed"
	"paganotoni/todox/internal/config"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed **/*.html
	tmpls embed.FS

	Templates = mdfs.New(
		tmpls,
		"internal",
		config.Environment,
	)
)
