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

func TestPasswordIsValid(t *testing.T) {
	user := User{}
	expectedResults := []ExpectedResult{
		{
			testedStr:      "pass",
			expectedResult: false,
		},
		{
			testedStr:      "password",
			expectedResult: false,
		},
		{
			testedStr:      "12345678",
			expectedResult: false,
		},
		{
			testedStr:      "password1",
			expectedResult: true,
		},
		{
			testedStr:      "",
			expectedResult: false,
		},
	}

	for _, expectedResult := range expectedResults {
		user.Password = expectedResult.testedStr
		passwordIsValid, err := user.passwordIsValid()
		if err != nil {
			log.Println(fmt.Errorf("passwordIsValid"))
			t.Error(err)
		}
		if passwordIsValid != expectedResult.expectedResult {
			t.Fail()
		}
	}
}

func TestEmailIsValid(t *testing.T) {
	user := User{}
	expectedResults := []ExpectedResult{
		{
			testedStr:      "email@email.com@email",
			expectedResult: false,
		},
		{
			testedStr:      "emai\nl@mail.com",
			expectedResult: false,
		},
		{
			testedStr:      "email@email.com.com",
			expectedResult: false,
		},
		{
			testedStr:      "email@email",
			expectedResult: false,
		},
		{
			testedStr:      "email@email.com",
			expectedResult: true,
		},
		{
			testedStr:      "",
			expectedResult: false,
		},
	}

	for _, expectedResult := range expectedResults {
		user.Email = expectedResult.testedStr
		emailIsValid, err := user.emailIsValid()
		if err != nil {
			log.Println(fmt.Errorf("emailIsValid"))
			t.Error(err)
		}
		if emailIsValid != expectedResult.expectedResult {
			t.Fail()
		}
	}
}
