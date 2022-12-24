//go:build mail_tests
// +build mail_tests

package main

import (
	"log"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := sendMail("newspaw1111@mail.ru", "Test", "Some big text")
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
}
