package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
type testData struct {
	p   People
	exp int
	err error
}

func TestLenData(t *testing.T) {
	tests := []testData{
		testData{p: People{}, exp: 0, err: nil},
		testData{p: People{Person{"A", "B", time.Now()}, Person{"B", "C", time.Now()}}, exp: 2, err: nil},
	}
	for _, test := range tests {
		got := test.p.Len()
		if got != test.exp {
			t.Error("Bad result!")
		}
	}
}
