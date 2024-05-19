package _1_basic_testing

import (
	"testing"
)

func TestHelloReturnFunction(t *testing.T) {
	expectedFuncResult := "hello world"
	if actualFuncResult := HelloReturn(); actualFuncResult != expectedFuncResult {
		t.Errorf("expected %s; got: %s;", expectedFuncResult, actualFuncResult)
	}
}

// `go test` for run