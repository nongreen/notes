package main

import (
	"log"
	"regexp"
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

func TestGenerateCode(t *testing.T) {
	code := generateCode()
	log.Println("Code: " + code)
	if len(code) != VERIFICATION_CODE_LEN {
		log.Println("Len of code isn't equal VERIFICATION_CODE_LEN")
		t.Fail()
	}

	re, err := regexp.Compile(`\D+`)
	if err != nil {
		t.Error(err)
	}
	if regexpResult := re.FindAll([]byte(code), -1); len(regexpResult) != 0 {
		log.Println("Code holds non digit symbols")
		t.Fail()
	}

	lastDigitInCode := code[0]
	var i int
	for i = 1; i < VERIFICATION_CODE_LEN; i += 1 {
		if code[i] != lastDigitInCode {
			break
		}
	}
	if i == VERIFICATION_CODE_LEN-1 && code[i] == lastDigitInCode {
		t.Fail()
	}
}
