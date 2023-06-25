package main

import (
	// "errors"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"database/sql"

	_ "github.com/lib/pq"

	andMapMemoriesAPI "aykelith/and-map-memories/src/pkg/api"
	andMapMemoriesDB "aykelith/and-map-memories/src/pkg/db"
)

func handleError(err error) {
	log.Print(err.Error())
	os.Exit(1)
}

func main() {
	var err error;
	var config Config;

	config, err = getConfig()
	if err != nil {
		handleError(err)
	}

	db, err := sql.Open("postgres", config.PostgreSQLDBURL)
	if err != nil {
		handleError(err)
	}
	defer db.Close()

	log.Println("GGG")

	mux := http.NewServeMux()

	if os.Getenv("ENV") == "production" {

	} else {
		workingDirectory, err := os.Getwd()
		if err != nil {
			log.Printf("[dev] Error getting the working directory: %s\n", err)
		}

		mux.Handle("/", http.FileServer(http.Dir(path.Join(workingDirectory, "..", "client", "static"))))
	}

	pinsTableHandler := andMapMemoriesDB.CreatePinsTableHandler(db)

	andMapMemoriesAPI.SetupApp(mux, pinsTableHandler)

	err = http.ListenAndServe(":15010", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	log.Printf("Started")
}