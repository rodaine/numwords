package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary_IncludeSecond(t *testing.T) {
	t.Parallel()

	n, ok := dictionary.m["second"]
	assert.True(t, ok)

	IncludeSecond(false)
	n, ok = dictionary.m["second"]
	assert.False(t, ok)

	IncludeSecond(true)
	n, ok = dictionary.m["second"]
	assert.True(t, ok)
	assert.EqualValues(t, second, n)
}

func TestDictionary_LookupNumber(t *testing.T) {
	t.Parallel()

	_, ok := lookupNumber("one")
	assert.True(t, ok)

	_, ok = lookupNumber("ONE")
	assert.True(t, ok)

	_, ok = lookupNumber("foobar")
	assert.False(t, ok)
}
