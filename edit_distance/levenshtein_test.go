package editdistance

import (
	"testing"
)

var testCases = []struct {
	source   string
	target   string
	distance int
}{
	{"", "a", 1},
	{"a", "aa", 1},
	{"a", "aaa", 2},
	{"", "", 0},
	{"a", "b", 1},
	{"aaa", "aba", 1},
	{"aaa", "ab", 2},
	{"a", "a", 0},
	{"ab", "ab", 0},
	{"a", "", 1},
	{"aa", "a", 1},
	{"aaa", "a", 2},
}

func TestEditDistance2words(t *testing.T) {
	for _, testcase := range testCases {
		distance := EditDistance2words(testcase.source, testcase.target)
		if distance != testcase.distance {
			t.Fatalf(
				"Distance between",
				testcase.source,
				"and",
				testcase.target,
				"computed as",
				distance,
				", should be",
				testcase.distance)
		}
	}
}
