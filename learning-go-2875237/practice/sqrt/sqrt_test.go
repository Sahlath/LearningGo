package sqrt

import (
	"fmt"
	"testing"
)

func almostEqual(val1, val2 float64) bool {
	return Abs(val1-val2) <= 0.001
}

func TestSqrt(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value %f", val)
	}
}

type testcase struct {
	input    float64
	expected float64
}

func TestMany(t *testing.T) {
	testcases := []testcase{
		{0.0, 0.0},
		{2.0, 1.414210},
		{9.0, 3.0},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%f", tc.input), func(t *testing.T) {
			out, err := Sqrt(tc.input)
			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}
