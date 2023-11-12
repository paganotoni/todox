package main

import (
	"fmt"
	"todox/internal/app/config"

	"github.com/leapkit/core/gloves"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		config.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}
