package main

import (
	"fmt"
	"todox/internal"

	"github.com/leapkit/core/tools/rebuilder"
)

func main() {
	err := rebuilder.Start(
		"cmd/app/main.go",

		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}
