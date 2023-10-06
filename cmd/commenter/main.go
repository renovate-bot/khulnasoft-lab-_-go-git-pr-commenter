package main

import (
	"log"
	"os"

	"github.com/khulnasoft-lab/go-git-pr-commenter/pkg/app"
)

func main() {
	app := app.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
