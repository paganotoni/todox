package internal

import (
	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

var (
	AssetsFolder = "./internal/assets"
	PublicFolder = "./public"

	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL = envor.Get("DATABASE_URL", "./todox.db")
	Environment = envor.Get("GO_ENV", "development")

	Port = envor.Get("PORT", "3000")
	Host = envor.Get("HOST", "0.0.0.0")

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []gloves.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(assets.Watcher(AssetsFolder, PublicFolder)),
		gloves.WatchExtension(".go"),
	}
)
