package main

import (
	"fmt"

	"paganotoni/todox/internal/config"

	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		gloves.WithRunner(func() {
			// Run the tailo watcher so when changes are made to
			// the html code it rebuilds css.

			tailo.Watch(
				tailo.UseInputPath("internal/web/assets/application.css"),
				tailo.UseConfigPath("internal/config/tailwind.config.js"),
				tailo.UseOutputPath("internal/web/public/application.css"),
			)
		}),

		// // Extensions to watch
		gloves.WatchExtension(config.GlovesExtensionsToWatch...),

		// Exclude paths from code reloading.
		gloves.ExcludePaths(config.GlovesExcludePaths...),
	)

	if err != nil {
		fmt.Println(err)
	}
}
