package main

import (
	"os"
	"os/exec"
	"todox/internal/app/config"

	"github.com/paganotoni/tailo"
)

func main() {
	tailo.Build(config.TailoOptions...)

	cmd := exec.Command("go", "build")
	cmd.Args = append(
		cmd.Args,

		`--ldflags`, `-linkmode=external -extldflags="-static"`,
		`-tags`, `osusergo,netgo,sqlite_omit_load_extension`,
		`-buildvcs=false`,
		"-o", "bin/app",
		"cmd/app/main.go",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
