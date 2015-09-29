package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExploder_Explode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in       string
		expected []string
	}{
		{"1,000,000 dollars", []string{"1000000", "dollars"}},
		{"Foo Bar", []string{"Foo", "Bar"}},
		{"twenty-one", []string{"twenty", "one"}},
		{"Nintey-Nine Red Balloons, by Nena", []string{"Nintey", "Nine", "Red", "Balloons", "by", "Nena"}},
	}

	for _, test := range tests {
		assert.EqualValues(t, test.expected, explode(test.in), "%+v", test)
	}
}
