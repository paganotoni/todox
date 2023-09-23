package main

import (
	"os"
	"os/exec"

	"github.com/paganotoni/tailo"
)

func main() {
	tailo.Build(
		tailo.UseInputPath("internal/web/assets/application.css"),
		tailo.UseConfigPath("internal/config/tailwind.config.js"),
		tailo.UseOutputPath("internal/web/public/application.css"),
	)

	cmd := exec.Command("go", "build")
	cmd.Args = append(
		cmd.Args,

		`--ldflags`, `-linkmode=external -extldflags="-static"`,
		`-tags`, `osusergo,netgo,musl`,
		`-buildvcs=false`,
		"-o", "bin/app",
		"cmd/app/main.go",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
