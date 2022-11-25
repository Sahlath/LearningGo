package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqrtSimple(t *testing.T) {
	val, err := Sqrt(2)
	require.NoError(t, err)
	require.InDelta(t, 1.414214, val, 0.001)
}

func loadTest(t *testing.T) []testcase {
	file, err := os.Open("sqrt_testcases.csv")
	require.NoError(t, err)

	defer file.Close()
	var cases []testcase
	rdr := csv.NewReader(file)
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		input, err := strconv.ParseFloat(record[0], 64)
		require.NoError(t, err)
		expected, err := strconv.ParseFloat(record[1], 64)
		require.NoError(t, err)
		tc := testcase{input, expected}
		cases = append(cases, tc)
	}
	return cases
}

func TestMany1(t *testing.T) {
	for _, tc := range loadTest(t) {
		t.Run(fmt.Sprintf("%f", tc.input), func(t *testing.T) {
			out, err := Sqrt(tc.input)
			require.NoError(t, err)
			require.InDelta(t, tc.expected, out, 0.001)
		})
	}
}

//benchamrk and profiling can be done
