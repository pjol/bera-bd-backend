package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pjol/bera-bd-backend/db"
	"github.com/pjol/bera-bd-backend/db/auth"
	"github.com/pjol/bera-bd-backend/db/cfi"
	"github.com/pjol/bera-bd-backend/db/ramps"
	"github.com/pjol/bera-bd-backend/logger"
	"github.com/pjol/bera-bd-backend/router"

	authHandlers "github.com/pjol/bera-bd-backend/handlers/auth"
	cfiHandlers "github.com/pjol/bera-bd-backend/handlers/cfi"
	rampsHandlers "github.com/pjol/bera-bd-backend/handlers/ramps"
)

func main() {
	godotenv.Load()

	adb, err := db.PgxDB("auth")
	if err != nil {
		log.Fatalf("error initializing auth db: %s", err)
	}

	cdb, err := db.PgxDB("cfi")
	if err != nil {
		log.Fatalf("error initializing cfi db: %s", err)
	}

	rdb, err := db.PgxDB("ramps")
	if err != nil {
		log.Fatalf("error initializing ramps db: %s", err)
	}

	l, err := logger.New("./logs/prod/app.log", "APP: ")
	if err != nil {
		log.Fatalf("error initializing app logger: %s\n", err)
	}
	defer l.Close()

	authDb := auth.NewDb(adb, l)
	err = authDb.CreateTables()
	if err != nil {
		log.Fatalf("error creating authdb tables: %s", err)
	}

	rampsDb := ramps.NewDb(rdb, l)
	err = rampsDb.CreateTables()
	if err != nil {
		log.Fatalf("error creating rampsdb tables: %s", err)
	}

	cfiDb := cfi.NewDb(cdb, l)
	err = cfiDb.CreateTables()
	if err != nil {
		log.Fatalf("error creating cfidb tables: %s", err)
	}

	a := authHandlers.New(authDb, l)
	r := rampsHandlers.New(rampsDb, l)
	c := cfiHandlers.New(cfiDb, l)

	router := router.New(a, r, c)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("now listening on port %s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	fmt.Println(err)
}
