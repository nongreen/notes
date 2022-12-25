// tools.go contains some useful simple functions for other functions

package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// nonwordSymbolsInString returns true, if string storages nonword symbols
func nonwordSymbolsInString(str string) (bool, error) {
	re, err := regexp.Compile(`[\w-]`)
	if err != nil {
		return false, err
	}

	result := re.FindAll([]byte(str), -1)
	if len(result) != len(str) {
		return true, nil
	}
	return false, nil
}

// passwordHaveMinStrong returns true, if lenght of string > 8, and string don't storages word symbols only
func passwordHaveMinStrong(password string) (bool, error) {
	if len(password) < 8 {
		return false, nil
	}

	re, err := regexp.Compile(`[A-z]`)
	if err != nil {
		return false, err
	}

	// Check password on contain word symbols only
	result := re.FindAll([]byte(password), -1)
	if len(result) == len(password) {
		return false, nil
	}

	re, err = regexp.Compile(`[0-9]`)
	if err != nil {
		return false, err
	}

	// Check password on contain digit symbols only
	result = re.FindAll([]byte(password), -1)
	if len(result) == len(password) {
		return false, nil
	}

	return true, nil
}

// stringIsEmail returns true, if string have someText@domain2.domain1 and haven't nonword symbols
func stringIsEmail(str string) (bool, error) {
	splitByAt := regexp.MustCompile(`@`).Split(str, 2)
	if len(splitByAt) != 2 {
		return false, nil
	}

	nonwordSymbolsInStr, err := nonwordSymbolsInString(splitByAt[0])
	if err != nil {
		return false, fmt.Errorf("string storage nonword symbols")
	}
	if nonwordSymbolsInStr {
		return false, nil
	}

	re, err := regexp.Compile(`[\w-]+\.[\w]+`)
	if err != nil {
		return false, fmt.Errorf("string don't storage domain like gmail.com")
	}
	result := re.FindAll([]byte(splitByAt[1]), -1)
	if len(result) != 1 {
		return false, nil
	}
	if len(result[0]) != len(splitByAt[1]) {
		return false, nil
	}

	return true, nil
}

// todo: Add security
func generateEmailVerificationLink(user *User) string {
	var result string
	if GIN_MODE == "debug" {
		result = "https://" + DOMAIN + ":" + "/email_verification?email="
	} else {
		result = "https://" + DOMAIN + "/email_verification?email="
	}
	result += hex.EncodeToString([]byte(user.Email))

	return result
}

// Handlers

// handleMailVerification gets uesr from context, checks it on valid and saves it. If have error
// returns bad request
func handleMailVerification(context *gin.Context) {
	user := &User{}
	if err := context.Bind(user); err != nil {
		if GIN_MODE == "debug" {
			log.Println("Error bind")
			log.Fatal(err)
		}
		return
	}

	userIsValid, err := user.isValid()
	if err != nil {
		if GIN_MODE == "debug" {
			log.Println("Error userIsValid")
			log.Fatal(err)
		}
		return
	}
	if !userIsValid {
		if GIN_MODE == "debug" {
			log.Println("User isn't valid")
		}
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = sendMail(user.Email, "Activation mail", "Your link: ")
	if err != nil {
		if GIN_MODE == "debug" {
			log.Println("Error by sending email")
		}
		context.AbortWithStatus(http.StatusBadRequest)
	}

	err = user.save()
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	/* todo: redirect to page with message */
	context.Redirect(http.StatusFound, "/")
}
