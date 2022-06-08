package coverage

import (
	"errors"
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
	p    People
	exp  int
	err  error
	i, j int
	flag bool
}

func TestLenData(t *testing.T) {
	tests := []testData{
		testData{p: People{}, exp: 0, err: nil},
		testData{p: People{Person{"A", "B", time.Now()}, Person{"B", "C", time.Now()}}, exp: 2, err: nil},
	}
	for _, test := range tests {
		got := test.p.Len()
		if got != test.exp {
			errors.New("Error")
		}
	}
}

func TestLessData(t *testing.T) {
	tests := []testData{
		testData{
			p:    People{Person{"A", "B", time.Now().Add(1000 * time.Millisecond)}, Person{"A", "B", time.Now()}},
			err:  nil,
			i:    0,
			j:    1,
			flag: true,
		},
		testData{
			p:    People{Person{"A", "B", time.Now()}, Person{"A", "B", time.Now().Add(1000 * time.Millisecond)}},
			err:  nil,
			i:    0,
			j:    1,
			flag: false,
		},
		testData{
			p:    People{Person{"A", "B", time.Now().Add(1000 * time.Millisecond)}, Person{"A", "B", time.Now()}},
			err:  nil,
			i:    1,
			j:    0,
			flag: false,
		},
		testData{
			p:    People{Person{"A", "B", time.Now()}, Person{"A", "B", time.Now().Add(1000 * time.Millisecond)}},
			err:  nil,
			i:    1,
			j:    0,
			flag: true,
		},
		testData{
			p:    People{Person{"A", "B", time.Now()}, Person{"A", "B", time.Now()}},
			err:  nil,
			i:    0,
			j:    0,
			flag: false,
		},
		testData{
			p:    People{Person{"A", "B", time.Now()}, Person{"A", "B", time.Now()}},
			err:  nil,
			i:    0,
			j:    1,
			flag: false,
		},
		testData{
			p:    People{Person{"A", "B", time.Now()}, Person{"B", "C", time.Now()}},
			err:  nil,
			i:    1,
			j:    0,
			flag: false,
		},
	}

	for _, test := range tests {
		got := test.p.Less(test.i, test.j)
		if got != test.flag {
			errors.New("Error")
		}
	}
}

func TestSwapData(t *testing.T) {
	tests := []testData{
		testData{
			p: People{Person{"A", "B", time.Now()}, Person{"B", "C", time.Now()}},
			i: 0,
			j: 1,
		},
	}
	for _, test := range tests {
		test.p.Swap(test.i, test.j)
		if test.p[0].firstName != "B" && test.p[0].lastName != "C" {
			errors.New("Error")
		}
	}
}

/*

 */
