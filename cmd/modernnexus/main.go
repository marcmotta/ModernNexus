// cmd/modernnexus/main.go
package main

import (
	"flag"
	"log"
	"os"

	"modernnexus/internal/modernnexus"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := modernnexus.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
