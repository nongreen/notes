package main

import (
	"testing"
)

func TestConnectToDB(t *testing.T) {
	connectToDB(&DBCONFIG)
}
