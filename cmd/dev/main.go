package main

import (
	"fmt"
	"todox/internal"

	"github.com/leapkit/core/tools/rebuilder"
)

func main() {
	err := rebuilder.Start(
		"cmd/app/main.go",

		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		rebuilder.WithRunner(internal.TailoWatcher),
		rebuilder.WithRunner(internal.Assets.Watch),
		rebuilder.WatchExtension(".go", ".css", ".js"),
	)

	if err != nil {
		fmt.Println(err)
	}
}
