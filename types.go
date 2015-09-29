package numwords

type numberType int8

const (
	numAnd numberType = iota
	numDirect
	numSingle
	numTens
	numBig
	numFraction
	numDirectOrdinal
	numSingleOrdinal
	numTensOrdinal
	numBigOrdinal
	numDone
)

var typeStrings = map[numberType]string{
	numAnd:           "&",
	numDirect:        "d",
	numSingle:        "s",
	numTens:          "t",
	numBig:           "b",
	numFraction:      "f",
	numDirectOrdinal: "D",
	numSingleOrdinal: "S",
	numTensOrdinal:   "T",
	numBigOrdinal:    "B",
}

func (t numberType) String() string {
	if s, ok := typeStrings[t]; ok {
		return s
	}
	return "_"
}

func maxType(a, b numberType) numberType {
	if a > b {
		return a
	}
	return b
}
