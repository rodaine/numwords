package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatterns_Done(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 20, denominator: 1, typ: numTens},
	}

	out := done(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(20), out[0].Value())
	assert.Equal(t, numDone, out[0].typ)
}

func TestPatterns_Drop(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1},
		number{numerator: 2},
	}

	out := drop(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, 2, out[0].numerator)
}

func TestPatterns_Add(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 20, denominator: 1, typ: numTens},
		number{numerator: 3, denominator: 1, typ: numSingle},
	}

	out := add(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(23), out[0].Value())
	assert.Equal(t, numTens, out[0].typ)

	ns = numbers{
		number{numerator: 100, denominator: 1, typ: numBig},
		number{numerator: 1, denominator: 2, typ: numFraction},
	}

	out = add(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(100.5), out[0].Value())
	assert.Equal(t, numFraction, out[0].typ)
}

func TestPatterns_Multiply(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 3, denominator: 1, typ: numSingle},
		number{numerator: 100, denominator: 1, typ: numBig},
	}

	out := multiply(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(300), out[0].Value())
	assert.Equal(t, numBig, out[0].typ)

	ns = numbers{
		number{numerator: 100, denominator: 1, typ: numBig},
		number{numerator: 1, denominator: 4, typ: numFraction},
	}

	out = multiply(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(25), out[0].Value())
	assert.Equal(t, numFraction, out[0].typ)
}

func TestPatterns_Combine(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 3, denominator: 1, typ: numSingle},
		number{numerator: 100, denominator: 1, typ: numBig},
	}

	out := combine(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(300), out[0].Value())
	assert.Equal(t, numBig, out[0].typ)

	ns = numbers{
		number{numerator: 100, denominator: 1, typ: numBig},
		number{numerator: 3, denominator: 1, typ: numSingle},
	}

	out = combine(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(103), out[0].Value())
	assert.Equal(t, numBig, out[0].typ)
}

func TestPatterns_CombineToLowest(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1000, denominator: 1, typ: numBig},
		number{numerator: 3, denominator: 1, typ: numSingle},
		number{numerator: 100, denominator: 1, typ: numBig},
	}

	out := combineToLowest(ns, 0)
	assert.Len(t, out, 2)
	assert.Equal(t, float64(300), out[1].Value())
	assert.Equal(t, numBig, out[1].typ)
	assert.Equal(t, float64(1000), out[0].Value())

	ns = numbers{
		number{numerator: 100, denominator: 1, typ: numBig},
		number{numerator: 3, denominator: 1, typ: numSingle},
		number{numerator: 1000, denominator: 1, typ: numBig},
	}

	out = combineToLowest(ns, 0)
	assert.Len(t, out, 2)
	assert.Equal(t, float64(103), out[0].Value())
	assert.Equal(t, numBig, out[0].typ)
	assert.Equal(t, float64(1000), out[1].Value())
}

func TestPatterns_YearOrDone(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 19, denominator: 1, typ: numDirect},
		number{numerator: 88, denominator: 1, typ: numTens},
	}

	out := yearOrDone(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(1988), out[0].Value())
	assert.Equal(t, numDone, out[0].typ)

	ns = numbers{
		number{numerator: 20, denominator: 1, typ: numTens},
		number{numerator: 15, denominator: 1, typ: numDirect},
	}

	out = yearOrDone(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(2015), out[0].Value())
	assert.Equal(t, numDone, out[0].typ)

	ns = numbers{
		number{numerator: 30, denominator: 1, typ: numTens},
		number{numerator: 0, denominator: 1, typ: numDirect},
	}

	out = yearOrDone(ns, 0)
	assert.Len(t, out, 2)
	assert.Equal(t, numDone, out[1].typ)
}

func TestPatterns_FractionOrDone(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1, denominator: 1, typ: numSingle},
		number{numerator: 4, denominator: 1, typ: numSingleOrdinal, ordinal: true},
	}

	out := fractionOrDone(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(0.25), out[0].Value())
	assert.Equal(t, numFraction, out[0].typ)
	assert.False(t, out[0].ordinal)

	ns = numbers{
		number{numerator: 2, denominator: 1, typ: numSingle},
		number{numerator: 4, denominator: 1, typ: numSingleOrdinal, ordinal: true},
	}

	out = fractionOrDone(ns, 0)
	assert.Len(t, out, 2)
	assert.Equal(t, float64(2), out[0].Value())
	assert.Equal(t, numDone, out[0].typ)
	assert.Equal(t, float64(4), out[1].Value())
	assert.Equal(t, numSingleOrdinal, out[1].typ)
	assert.True(t, out[1].ordinal)
}

func TestPatterns_FractionOrCombine(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1, denominator: 1, typ: numDirect},
		number{numerator: 4, denominator: 1, typ: numSingleOrdinal, ordinal: true},
	}

	out := fractionOrCombine(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(0.25), out[0].Value())
	assert.Equal(t, numFraction, out[0].typ)
	assert.False(t, out[0].ordinal)

	ns = numbers{
		number{numerator: 20, denominator: 1, typ: numTens},
		number{numerator: 4, denominator: 1, typ: numSingleOrdinal, ordinal: true},
	}

	out = fractionOrCombine(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(24), out[0].Value())
	assert.Equal(t, numSingleOrdinal, out[0].typ)
	assert.True(t, out[0].ordinal)
}

func TestPatterns_AddAnd(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 2, denominator: 1, typ: numSingle},
		number{numerator: 0, denominator: 0, typ: numAnd},
		number{numerator: 1, denominator: 2, typ: numFraction},
	}

	out := addAnd(ns, 0)
	assert.Len(t, out, 1)
	assert.Equal(t, float64(2.5), out[0].Value())
	assert.Equal(t, numFraction, out[0].typ)
}

func TestPatterns_AllHaveHandlers(t *testing.T) {
	t.Parallel()

	for _, p := range patterns {
		_, ok := patternHandlers[p]
		assert.True(t, ok, "no handler for pattern `%s`", p)
	}
}
