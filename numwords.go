// Package numwords is a utility that converts textual numbers to their
// actual numeric values. The converted numbers can be parsed out as strings,
// integers, or floats as desired.
//
// Source: https://github.com/rodaine/numwords
package numwords

import "strings"

// ParseFloat reads a text string and converts it to its float value. An error
// is returned if the if the string cannot be resolved to a single float value.
func ParseFloat(s string) (float64, error) {
	in := explode(s)
	buf := numbers{}

	ok := false
	for i := range in {
		if buf, ok = readIntoBuffer(i, in, buf); !ok {
			return -1, ErrNonNumber
		}
	}

	return reduce(buf).Float()
}

// ParseInt reads a text string and converts it to its integer value. An error
// is returned if the if the string cannot be resolved to a single integer value.
// Fractional portions of the number will be truncated.
func ParseInt(s string) (int, error) {
	in := explode(s)
	buf := numbers{}

	ok := false
	for i := range in {
		if buf, ok = readIntoBuffer(i, in, buf); !ok {
			return -1, ErrNonNumber
		}
	}

	return reduce(buf).Int()
}

// ParseString reads a text string and converts all numbers contained within to
// their appropriate values. Integers are preserved exactly while floating point
// numbers are limited to six decimal places. The rest of the string is preserved.
func ParseString(s string) string {
	in := explode(s)
	out := ParseStrings(in)
	return strings.Join(out, " ")
}

// ParseStrings performs the same actions as ParseString but operates on a pre-
// sanitized and split string. This method is exposed for convenience if further
// processing of the string is required.
func ParseStrings(in []string) []string {
	out := make([]string, 0, 1)
	buf := numbers{}

	ok := false
	for i, s := range in {
		if buf, ok = readIntoBuffer(i, in, buf); !ok {
			out = buf.flush(out)
			buf = buf[:0]
			out = append(out, s)
		}
	}

	return buf.flush(out)
}

func readIntoBuffer(i int, in []string, buf numbers) (out numbers, ok bool) {
	s := in[i]
	n, ok := lookupNumber(s)

	if ok && n.typ != numAnd {
		buf = append(buf, n)
		return buf, ok
	} else if ok && n.typ == numAnd && shouldIncludeAnd(in, buf, i) {
		buf = append(buf, n)
		return buf, ok
	} else if n, ok = maybeNumeric(s); ok {
		buf = append(buf, n)
		return buf, ok
	}

	return buf, false
}

func shouldIncludeAnd(in []string, buf numbers, idx int) bool {
	if len(buf) == 0 || idx+1 >= len(in) {
		return false
	}

	prev := buf[len(buf)-1]

	if prev.ordinal || prev.typ == numFraction {
		return false
	}

	s := in[idx+1]
	if _, ok := lookupNumber(s); !ok {
		_, ok = maybeNumeric(s)
		return ok
	}

	return true
}
