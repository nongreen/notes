package main

import (
	"fmt"
	"log"
	"testing"
)

func TestUsernameIsValid(t *testing.T) {
	usernames := ["username"]
	user := User{}
	usernameIsValid, err := user.usernameIsValid()
	if err != nil {
		t.Error(fmt.Errorf("UsernameIsValid"))
	}
	if !usernameIsValid {
		log.Println("True string, isn't true")
		t.Fail()
	}
}
