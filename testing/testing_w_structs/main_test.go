package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	type TestCase struct {
		Name           string
		XValue         int
		YValue         int
		ExpectedResult int
		ExpectedError  error
	}

	testCases := []TestCase{
		{
			"Basic Add - tens",
			1,
			1,
			2,
			nil,
		},
		{
			"Basic Add - hundreds",
			100,
			1,
			101,
			nil,
		},
	}

	for _, testCase := range testCases {
		log.Printf("Running Test='%s'", testCase.Name)
		result, err := Add(testCase.XValue, testCase.YValue)
		if err != nil {
			assert.ErrorIs(t, err, testCase.ExpectedError)
		}

		assert.Equal(t, result, testCase.ExpectedResult)
	}
}
