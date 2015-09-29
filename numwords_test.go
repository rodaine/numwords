package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWords_ParseFloat(t *testing.T) {
	t.Parallel()

	_, err := ParseFloat("foobar")
	assert.Equal(t, ErrNonNumber, err)

	tests := []struct {
		in  string
		out float64
	}{
		{"one half", 0.5},
		{"one quarter", 0.25},
		{"three and a quarter", 3.25},
		{"one fifth", 0.2},
		{"nineteen eighty eight", 1988},
	}

	for _, test := range tests {
		f, err := ParseFloat(test.in)
		if assert.NoError(t, err) {
			assert.Equal(t, test.out, f, test.in)
		}
	}
}

func TestNumWords_ParseInt(t *testing.T) {
	t.Parallel()

	_, err := ParseInt("foobar")
	assert.Equal(t, ErrNonNumber, err)

	tests := []struct {
		in  string
		out int
	}{
		{"twelve", 12},
		{"twelve and a half", 12},
		{"zero", 0},
		{"nineteen eighty eight", 1988},
	}

	for _, test := range tests {
		i, err := ParseInt(test.in)
		if assert.NoError(t, err) {
			assert.Equal(t, test.out, i, test.in)
		}
	}
}

func TestNumWords_ParseString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		out string
	}{
		{"foo", "foo"},
		{"foo bar baz", "foo bar baz"},
		{"foo eleven", "foo 11"},
		{"a foo", "1 foo"},
		{"zero bar three foo", "0 bar 3 foo"},
		{"fifteen", "15"},
		{"ten eleven", "10 11"},
		{"one", "1"},
		{"two three", "2 3"},
		{"fifteen three", "15 3"},
		{"seven eighteen", "7 18"},
		{"one eighteen seven thirteen three three", "1 18 7 13 3 3"},
		{"twenty", "20"},
		{"twenty five", "25"},
		{"twenty zero", "20 0"},
		{"twenty five twenty", "25 20"},
		{"zero twenty thirty", "0 20 30"},
		{"three twenty three", "3 23"},
		{"ninety nine red balloons", "99 red balloons"},
		{"hundred", "100"},
		{"eleven hundred", "1100"},
		{"hundred eleven", "111"},
		{"four thousand", "4000"},
		{"thousand four", "1004"},
		{"three hundred twenty five", "325"},
		{"three hundred thousand", "300000"},
		{"twenty five hundred", "2500"},
		{"one hundred twenty one", "121"},
		{"thousand one hundred", "1100"},
		{"four hundred thirty one", "431"},
		{"fourteen hundred sixty seven", "1467"},
		{"one thousand four hundred sixty seven", "1467"},
		{"four thousand three hundred twenty one", "4321"},
		{"one million three hundred thousand", "1300000"},
		{"nineteen eighty eight", "1988"},
		{"twenty ten", "2010"},
		{"one half", "0.5"},
		{"three halves", "1.5"},
		{"one ninth", "0.111111"},
		{"one twentieth", "0.05"},
		{"one sixteenth", "0.0625"},
		{"one hundredth", "0.01"},
		{"seven o'clock", "7 o'clock"},
		{"two thirds", "0.666667"},
		{"one quarter of americans were born before nineteen eighty", "0.25 of americans were born before 1980"},
		{"ten fourtieths", "0.25"},
		{"nine hundred and ninety nine", "999"},
		{"zeroth", "0th"},
		{"one", "1"},
		{"five", "5"},
		{"ten", "10"},
		{"twenty seven", "27"},
		{"forty one", "41"},
		{"fourty two", "42"},
		{"a hundred", "100"},
		{"one hundred", "100"},
		{"one hundred and fifty", "150"},
		{"5 hundred", "500"},
		{"one thousand", "1000"},
		{"one thousand two hundred", "1200"},
		{"seventeen thousand", "17000"},
		{"twenty one thousand four hundred and seventy three", "21473"},
		{"seventy four thousand and two", "74002"},
		{"ninety nine thousand nine hundred ninety nine", "99999"},
		{"100 thousand", "100000"},
		{"two hundred fifty thousand", "250000"},
		{"one million two hundred fifty thousand and seven", "1250007"},
		{"the world population is seven billion two hundred seventy five million five hundred seventy eight thousand eight hundred eighty seven", "the world population is 7275578887"},
		{"two and a half", "2.5"},
		{"1 quarter", "0.25"},
		{"three quarters", "0.75"},
		{"one and a quarter", "1.25"},
		{"two & three eighths", "2.375"},
		{"1/2", "1/2"},
		{"07/10", "07/10"},
		{"three sixteenths", "0.1875"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, ParseString(test.in), test.in)
	}
}

func TestNumWords_ShouldIncludeAnd(t *testing.T) {
	t.Parallel()

	in := []string{"cat", "and"}
	buf := numbers{}
	ok := shouldIncludeAnd(in, buf, 1)
	assert.False(t, ok, "empty buffer, nothing to and")

	in = []string{"two", "and"}
	buf = numbers{number{}}
	ok = shouldIncludeAnd(in, buf, 1)
	assert.False(t, ok, "no more input strings available")

	in = []string{"2nd", "and", "three"}
	buf = numbers{number{ordinal: true}}
	ok = shouldIncludeAnd(in, buf, 1)
	assert.False(t, ok, "previous is ordinal")

	in = []string{"half", "and", "three"}
	buf = numbers{number{typ: numFraction}}
	ok = shouldIncludeAnd(in, buf, 1)
	assert.False(t, ok, "previous is a fraction")

	in = []string{"two", "and", "foo"}
	buf = numbers{number{}}
	ok = shouldIncludeAnd(in, buf, 1)
	assert.False(t, ok, "next is not a number")

	in = []string{"two", "and", "3"}
	buf = numbers{number{}}
	ok = shouldIncludeAnd(in, buf, 1)
	assert.True(t, ok, "numeric is ok")

	in = []string{"two", "and", "three"}
	buf = numbers{number{}}
	ok = shouldIncludeAnd(in, buf, 1)
	assert.True(t, ok, "the ideal case")
}
