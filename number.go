package numwords

import (
	"math/big"
	"strconv"
	"strings"
)

type number struct {
	numerator   int
	denominator int
	typ         numberType
	ordinal     bool
}

var ordinals = map[int]string{
	1: "st",
	2: "nd",
	3: "rd",
	4: "th",
}

func (n number) Value() float64 {
	return float64(n.numerator) / float64(n.denominator)
}

func (n number) String() string {
	if n.ordinal {
		suffix := "th"
		if r := n.numerator % 10; r > 0 && r < 4 {
			if rr := n.numerator % 100; rr < 11 || rr > 13 {
				suffix = ordinals[r]
			}
		}
		return strconv.Itoa(n.numerator) + suffix
	}

	if n.denominator != 1 {
		s := strconv.FormatFloat(n.Value(), 'f', 6, 64)
		return strings.TrimRight(s, ".0")
	}

	return strconv.Itoa(n.numerator)
}

func maybeNumeric(s string) (n number, ok bool) {
	s = strings.Replace(s, ",", "", -1)

	ord := false
	for _, suffix := range ordinals {
		if strings.HasSuffix(s, suffix) {
			ord = true
			s = strings.TrimSuffix(s, suffix)
			break
		}
	}

	if i, err := strconv.Atoi(s); err == nil {
		ok = true
		n.numerator = i
		n.denominator = 1
		n.ordinal = ord
	} else if f, err := strconv.ParseFloat(s, 64); err == nil {
		rat := big.NewRat(1, 1).SetFloat64(f)
		if rat != nil {
			ok = true
			n.numerator = int(rat.Num().Int64())
			n.denominator = int(rat.Denom().Int64())
			n.typ = numFraction
		}
		return
	} else {
		return
	}

	if !ord {
		switch {
		case n.numerator < 10 && n.numerator > 0:
			n.typ = numSingle
		case n.numerator >= 20 && n.numerator < 100:
			n.typ = numTens
		case n.numerator >= 100:
			n.typ = numBig
		default:
			n.typ = numDirect
		}
	} else {
		switch {
		case n.numerator < 10 && n.numerator > 0:
			n.typ = numSingleOrdinal
		case n.numerator >= 20 && n.numerator < 100:
			n.typ = numTensOrdinal
		case n.numerator >= 100:
			n.typ = numBigOrdinal
		default:
			n.typ = numDirectOrdinal
		}
	}

	return
}
