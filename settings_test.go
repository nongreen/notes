package main

import (
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	if SECRETKEY == "" {
		log.Println("SECRETKEY is empty")
		t.Fail()
	}
	if GIN_MODE == "" {
		log.Println("GIN_MODE is empty")
		t.Fail()
	}
	if DOMAIN == "" {
		log.Println("DOMAIN is empty")
		t.Fail()
	}
	if PORT == "" {
		log.Println("PORT is empty")
		t.Fail()
	}
	if USERNAMELEN == 0 {
		log.Println("USERNAMELEN is empty")
		t.Fail()
	}
	if PASSLEN == 0 {
		log.Println("PASSLEN is empty")
		t.Fail()
	}
	if DBCONFIG.User == "" {
		log.Println("DBCONFIG.User is empty")
		t.Fail()
	}
	if DBCONFIG.Passwd == "" {
		log.Println("DBCONFIG.Passwd is empty")
		t.Fail()
	}
	if DBCONFIG.Net == "" {
		log.Println("DBCONFIG.Net is empty")
		t.Fail()
	}
	if DBCONFIG.Addr == "" {
		log.Println("DBCONFIG.Addr is empty")
		t.Fail()
	}
	if DBCONFIG.DBName == "" {
		log.Println("DBCONFIG.DBName is empty")
		t.Fail()
	}
}
