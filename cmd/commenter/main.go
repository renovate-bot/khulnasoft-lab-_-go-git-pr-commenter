package main

import (
	"github.com/khulnasoft-lab/go-git-pr-commenter/pkg/app"
	"log"
	"os"
)

func main() {
	app := app.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}