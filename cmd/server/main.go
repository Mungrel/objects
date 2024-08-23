package main

import (
	"log"
	"os"

	"github.com/Mungrel/objects/internal/db"
)

func main() {
	mode := "server"

	if len(os.Args) > 1 {
		mode = os.Args[1] // Arg 0 is always the binary name.
	}

	switch mode {
	case "server":
		runServer()
	case "migrate":
		runMigrate()
	default:
		log.Fatalf("Unknown mode %q", mode)
	}
}

func runServer() {
	// TODO, start up all gRPC services.
}

func runMigrate() {
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations up-to-date")
}
