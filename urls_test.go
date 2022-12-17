package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEndPoints(t *testing.T) {
	go func() {
		router := gin.Default()
		router.LoadHTMLGlob("html/*.html")

		bindAllEndPoints(router)

		router.Run(fmt.Sprintf("%s:%s", DOMAIN, PORT))
	}()

	// get requests
	var getEndPointsForTest []string = []string{
		"/",
		"/register",
		"/login",
	}
	for _, endPoint := range getEndPointsForTest {
		testGetEndPoint(endPoint, t)
	}

	// post requests
	data := url.Values{
		"Username": {"Testusername"},
		"Password": {"testtest21"},
		"Email":    {"test@test.ru"},
	}
	baseUrl := fmt.Sprintf("http://%s:%s", DOMAIN, PORT)
	testingEndPoint := "/register"
	resp, err := http.PostForm(baseUrl+testingEndPoint, data)
	if err != nil {
		log.Fatal(err)
		t.Error()
	}
	if !((200 <= resp.StatusCode) && (resp.StatusCode < 300)) {
		t.Fail()
	}
}

// testGetEndPoint checks testingEndPoint on response with error status code by GET request
func testGetEndPoint(testingEndPoint string, t *testing.T) {
	baseUrl := fmt.Sprintf("http://%s:%s", DOMAIN, PORT)
	log.Println("testing url:" + baseUrl + testingEndPoint)

	resp, err := http.Get(baseUrl + testingEndPoint)
	if err != nil {
		log.Fatal(err)
		t.Error()
	}
	if !((200 <= resp.StatusCode) && (resp.StatusCode < 300)) {
		t.Fail()
	}
}
