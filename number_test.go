package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber_Value(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		d        int
		expected float64
	}{
		{1, 1, 1},
		{2, 1, 2},
		{1, 2, 0.5},
	}

	for _, test := range tests {
		n := number{numerator: test.n, denominator: test.d}
		assert.Equal(t, test.expected, n.Value(), "%+v", test)
	}
}

func TestNumber_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		d        int
		o        bool
		expected string
	}{
		{1, 1, false, "1"},
		{2, 1, false, "2"},
		{1, 2, false, "0.5"},
		{1, 3, false, "0.333333"},
		{2, 3, false, "0.666667"},
		{1, 1, true, "1st"},
		{2, 1, true, "2nd"},
		{3, 1, true, "3rd"},
		{4, 1, true, "4th"},
		{11, 1, true, "11th"},
		{12, 1, true, "12th"},
		{13, 1, true, "13th"},
		{21, 1, true, "21st"},
		{22, 1, true, "22nd"},
		{23, 1, true, "23rd"},
		{24, 1, true, "24th"},
		{1, 2, true, "1st"},
	}

	for _, test := range tests {
		n := number{
			numerator:   test.n,
			denominator: test.d,
			ordinal:     test.o,
		}
		assert.Equal(t, test.expected, n.String(), "%+v", test)
	}
}

func TestNumber_MaybeNumeric(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		val float64
		typ numberType
		ok  bool
	}{
		{in: "foo", ok: false},
		{in: "fist", ok: false},

		{"12", float64(12), numDirect, true},
		{"2", float64(2), numSingle, true},
		{"22", float64(22), numTens, true},
		{"222", float64(222), numBig, true},

		{"0.5", float64(.5), numFraction, true},

		{"12th", float64(12), numDirectOrdinal, true},
		{"2nd", float64(2), numSingleOrdinal, true},
		{"22nd", float64(22), numTensOrdinal, true},
		{"222nd", float64(222), numBigOrdinal, true},

		{"2,222,222", float64(2222222), numBig, true},
	}

	for _, test := range tests {
		n, ok := maybeNumeric(test.in)
		if assert.Equal(t, test.ok, ok) && ok {
			assert.Equal(t, test.val, n.Value())
			assert.Equal(t, test.typ, n.typ)
		}
	}
}
