package main

import (
	"fmt"
	"log"
	"testing"
)

func TestUsernameIsValid(t *testing.T) {
	user := User{}
	expectedResults := []ExpectedResult{
		{
			testedStr:      "username",
			expectedResult: true,
		},
		{
			testedStr:      "user name",
			expectedResult: false,
		},
		{
			testedStr:      "user\nanme",
			expectedResult: false,
		},
		{
			testedStr:      "username5",
			expectedResult: true,
		},
		{
			testedStr:      "",
			expectedResult: false,
		},
	}

	for _, expectedResult := range expectedResults {
		user.Username = expectedResult.testedStr
		usernameIsValid, err := user.usernameIsValid()
		if err != nil {
			log.Println(fmt.Errorf("usernameIsValid"))
			t.Error(err)
		}
		if usernameIsValid != expectedResult.expectedResult {
			t.Fail()
		}
	}
}
