package handler

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"

	_ "embed"
)

var (
	ConfigDir string
	dbname    string = "/todo.db"
)

//go:embed sqlc/schema.sql
var ddl string

func Config() {
	configDir()

}

func makeDir(d string, p fs.FileMode) {
	if _, err := os.Stat(d); err != nil {
		log.Printf("Make Directory: %v, Perm: %v\n", d, p)
		if err := os.Mkdir(d, p); err != nil {
			fmt.Println(err)
		}
	}
}

func configDir() {
	homeDir := os.Getenv("HOME")
	if _, err := os.Stat(homeDir); err != nil {
		panic(err)
	}

	ConfigDir = homeDir + "/.config"
	makeDir(ConfigDir, 0766)

	ConfigDir = ConfigDir + "/go-todo"
	makeDir(ConfigDir, 0644)
}

func DbInit() {
	db, err := sql.Open("sqlite3", ConfigDir+dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	// queries := database.New(db)

	return nil
}
