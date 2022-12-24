// main_test.go contains TestMain for other tests

package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	db, err = connectToTestDB(&DBCONFIG)
	if err != nil {
		log.Fatal("Cannot connect to test db")
		os.Exit(1)
	}
	defer db.Close()

	code := m.Run()
	clearDB(*db)
	os.Exit(code)
}
