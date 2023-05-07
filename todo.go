package todox

import (
	"embed"
	"paganotoni/todox/internal/fs"

	"github.com/gofrs/uuid"
)

var (
	List = []Todo{
		// TODO: this should be a database
		{ID: uuid.Must(uuid.NewV4()), Content: "Task A", Completed: false},
		{ID: uuid.Must(uuid.NewV4()), Content: "Task B", Completed: false},
		{ID: uuid.Must(uuid.NewV4()), Content: "Task C", Completed: false},
	}

	//go:embed *.html */*.html
	templateFS embed.FS
	Templates  = fs.NewFallback(templateFS, ".")
)

type Todo struct {
	ID        uuid.UUID
	Content   string
	Completed bool
}
