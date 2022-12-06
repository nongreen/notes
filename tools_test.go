package main

import (
	"log"
	"testing"
)

func TestNonwordSymbolsInString(t *testing.T) {
	var expectedResults []ExpectedResult = []ExpectedResult{
		{
			testedStr:      "simplestring",
			expectedResult: false,
		},
		{
			testedStr:      "String with spaces",
			expectedResult: true,
		},
		{
			testedStr:      "String with nonword \n symbol ",
			expectedResult: true,
		},
		{
			testedStr:      "StringWithDigit5",
			expectedResult: false,
		},
		{
			testedStr:      "AllWellBut\b",
			expectedResult: true,
		},
	}
	testExpectedResults(t, expectedResults, nonwordSymbolsInString)
}

func TestPasswordHaveMinStrong(t *testing.T) {
	var expectedResults []ExpectedResult = []ExpectedResult{
		{
			testedStr:      "password",
			expectedResult: false,
		},
		{
			testedStr:      "11",
			expectedResult: false,
		},
		{
			testedStr:      "password11",
			expectedResult: true,
		},
		{
			testedStr:      "daffgafdg11n.,asdf",
			expectedResult: true,
		},
		{
			testedStr:      "342523542345",
			expectedResult: false,
		},
	}
	testExpectedResults(t, expectedResults, passwordHaveMinStrong)
}

func TestStringIsEmail(t *testing.T) {
	var expectedResults []ExpectedResult = []ExpectedResult{
		{
			testedStr:      "password",
			expectedResult: false,
		},
		{
			testedStr:      "email@",
			expectedResult: false,
		},
		{
			testedStr:      "email@email",
			expectedResult: false,
		},
		{
			testedStr:      "email.email@email.com",
			expectedResult: false,
		},
		{
			testedStr:      "email@email\b.com",
			expectedResult: false,
		},
		{
			testedStr:      "email@email.com",
			expectedResult: true,
		},
	}
	testExpectedResults(t, expectedResults, stringIsEmail)
}

func testExpectedResults(t *testing.T, expectedResults []ExpectedResult, tested_func func(string) (bool, error)) {
	for _, element := range expectedResults {
		result, err := tested_func(element.testedStr)
		if err != nil {
			t.Error(err)
		}
		if result != element.expectedResult {
			log.Println("tested string:", element.testedStr)
			t.Fail()
		}
	}
}
