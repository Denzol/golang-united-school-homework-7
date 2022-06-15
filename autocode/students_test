package coverage

import (
	"errors"
	"fmt"
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
	p        People
	m        *Matrix
	exp      int
	err      error
	i, j     int
	flag     bool
	str      string
	expValue [][]int
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

func TestSwapNew(t *testing.T) {
	tests := []testData{
		testData{
			m:   &Matrix{},
			str: "",
			err: nil,
		},
		testData{
			m:   &Matrix{2, 1, []int{11, 12}},
			str: "11 \n 12",
			err: nil,
		},
		testData{
			m:   &Matrix{},
			str: "Hello /n World!",
			err: fmt.Errorf("strconv.Atoi: parsing \"Hello\": invalid syntax"),
		},
		testData{
			m:   &Matrix{},
			str: "11 \n 12 14",
			err: nil,
		},
		testData{
			m:   &Matrix{3, 2, []int{11, 12}},
			str: "11 12 \n 13 14 \n 15 16",
			err: nil,
		},
	}
	for _, test := range tests {
		got, err := New(test.str)
		if got != test.m || err != nil {
			errors.New("Error")
		}
	}
}

func TestRows(t *testing.T) {
	tests := []testData{
		testData{
			m:        &Matrix{2, 2, []int{2, 4, 6, 8}},
			expValue: [][]int{[]int{2, 4}, []int{6, 8}},
		},
		testData{
			m:        &Matrix{0, 0, []int{}},
			expValue: [][]int{},
		},
	}
	for _, test := range tests {
		got := test.m.Rows()
		for i, k := range got {
			if k[i] != test.expValue[i][i] {
				errors.New("Error")
			}
		}

	}
}
func TestCols(t *testing.T) {
	tests := []testData{
		testData{
			m:        &Matrix{2, 2, []int{2, 4, 6, 8}},
			expValue: [][]int{[]int{2, 4}, []int{6, 8}},
		},
		testData{
			m:        &Matrix{0, 0, []int{}},
			expValue: [][]int{},
		},
	}
	for _, test := range tests {
		got := test.m.Cols()
		for i, k := range got {
			if k[i] != test.expValue[i][i] {
				errors.New("Error")
			}
		}

	}
}

func TestSet(t *testing.T) {
	tests := []testData{
		testData{
			m:    &Matrix{2, 2, []int{2, 4, 6, 8}},
			flag: true,
		},
		testData{
			m:    &Matrix{0, 0, []int{}},
			flag: false,
		},
	}
	for _, test := range tests {
		got := test.m.Set(1, 1, 3)
		if got != test.flag {
			errors.New("Error")
		}
	}

}
