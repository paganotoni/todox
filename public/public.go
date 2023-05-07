package public

import "embed"

// Folder is the embedded filesystem for the static files in this folder.
//
//go:embed *.css
var Folder embed.FS
