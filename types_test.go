package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes_String(t *testing.T) {
	t.Parallel()

	for typ, expected := range typeStrings {
		assert.Equal(t, expected, typ.String())
	}
	assert.Equal(t, "_", numberType(-1).String())
}

func TestTypes_MaxType(t *testing.T) {
	t.Parallel()

	assert.Equal(t, numSingle, maxType(numDirect, numSingle))
	assert.Equal(t, numDone, maxType(numDone, numBigOrdinal))
	assert.Equal(t, numTens, maxType(numTens, numTens))
}
