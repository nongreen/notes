package main

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
)

var test_db *sql.DB

/*func TestMain(m *testing.M) {
	connectToTestDB(&DBCONFIG)
	code := m.Run()
	os.Exit(code)
} */

func connectToTestDB(config *mysql.Config) error {
	config = config.Clone()

	config.DBName += "_tests"

	var err error
	test_db, err = connectToDB(config)

	return err
}

func TestConnectToDB(t *testing.T) {
	connectToDB(&DBCONFIG)

	config := DBCONFIG.Clone()
	config.User = ""
	_, err := connectToDB(config)
	if err == nil {
		t.Fail()
	}
}

func TestAddUser(t *testing.T) {

}
