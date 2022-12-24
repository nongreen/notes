package main

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func connectToTestDB(config *mysql.Config) (*sql.DB, error) {
	config = config.Clone()
	config.DBName += "_tests"
	db, err := connectToDB(config)
	return db, err
}

func clearDB(db sql.DB) {
	db.Query("DELETE FROM User;")
}

func TestConnectToDB(t *testing.T) {
	connectToDB(&DBCONFIG)

	// Test connect with wrong config
	config := DBCONFIG.Clone()
	config.User = ""
	_, err := connectToDB(config)
	if err == nil {
		t.Error(err)
	}
}

func TestAddUser(t *testing.T) {
	var err error
	test_user := User{
		Username: "test_username",
		Password: "password12",
		Email:    "testmail@test.org",
	}
	test_user.ID, err = addUser(&test_user, db)
	if err != nil {
		t.Error(err)
	}

	isExist, err := userIsExist(&test_user, db)
	if err != nil {
		t.Error(err)
	}
	if !isExist {
		t.Fail()
	}
}

func TestUserIsExist(t *testing.T) {
	var err error
	test_user := User{
		ID:       32,
		Username: "test_username",
		Password: "password12",
		Email:    "testmail@test.org",
	}

	// Check wrong case
	subtestUserExist(t, &test_user, db, false)

	// Check true case
	test_user.ID, err = addUser(&test_user, db)
	if err != nil {
		t.Error(err)
	}
	subtestUserExist(t, &test_user, db, true)

	// Check valid id, but invalid username
	validUsername := test_user.Username
	test_user.Username = "another username"
	subtestUserExist(t, &test_user, db, false)
	test_user.Username = validUsername

	// Check valid id, but invalid password
	validPassword := test_user.Password
	test_user.Password = "another password"
	subtestUserExist(t, &test_user, db, false)
	test_user.Password = validPassword

	// Check valid id, but invalid email
	validEmail := test_user.Email
	test_user.Email = "another email"
	subtestUserExist(t, &test_user, db, false)
	test_user.Email = validEmail
}

func subtestUserExist(t *testing.T, test_user *User, db *sql.DB, expected bool) {
	var isExist bool
	var err error
	isExist, err = userIsExist(test_user, db)
	if err != nil {
		t.Error(err)
	}
	if isExist != expected {
		t.Fail()
	}
}
