package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

type testCase struct {
	name        string
	a, b        float64
	want        float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "2+2=4", a: 2, b: 2, want: 4, errExpected: false},
		{name: "1+1=2", a: 1, b: 1, want: 2, errExpected: false},
		{name: "5+0=0", a: 5, b: 0, want: 5, errExpected: false},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%s): want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestRandomAdd(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)

		if want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "4-2=2", a: 4, b: 2, want: 2, errExpected: false},
		{name: "1-1=0", a: 1, b: 1, want: 0, errExpected: false},
		{name: "5-0=5", a: 5, b: 0, want: 5, errExpected: false},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%s): want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "4x2=8", a: 4, b: 2, want: 8, errExpected: false},
		{name: "1x1=1", a: 1, b: 1, want: 1, errExpected: false},
		{name: "5x0=0", a: 5, b: 0, want: 0, errExpected: false},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%s): want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "8/4=2", a: 8, b: 4, want: 2, errExpected: false},
		{name: "5/1=5", a: 5, b: 1, want: 5, errExpected: false},
		{name: "3/0=0", a: 3, b: 0, want: 0, errExpected: true},
		{name: "3/0=0", a: 0, b: 3, want: 0, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if tc.errExpected != (err != nil) {
			t.Fatalf("unexpected error - %v", err)
		} else {
			if tc.want != got {
				t.Errorf("Divide(%s): want %f, got %f", tc.name, tc.want, got)
			}
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "sqrt 16", a: 16, b: 0, want: 4, errExpected: false},
		{name: "sqrt 25", a: 25, b: 0, want: 5, errExpected: false},
		{name: "sqrt 1", a: 1, b: 0, want: 1, errExpected: false},
		{name: "sqrt 0", a: 0, b: 0, want: 0, errExpected: false},
		{name: "sqrt -1", a: -1, b: 0, want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		if tc.errExpected != (err != nil) {
			t.Fatalf("unexpected error - %v", err)
		} else {
			if tc.want != got {
				t.Errorf("Sqrt(%s): want %f, got %f", tc.name, tc.want, got)
			}
		}
	}

}
