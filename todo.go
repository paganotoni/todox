package todox

import (
	"embed"
	"paganotoni/todox/internal/fs"

	"github.com/gofrs/uuid"
)

var (
	//go:embed *.html */*.html
	templateFS embed.FS
	Templates  = fs.NewFallback(templateFS, ".")
)

type Todo struct {
	ID        uuid.UUID
	Content   string
	Completed bool
}
