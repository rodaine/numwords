package numwords

import (
	"strings"
	"sync"
)

var second = number{2, 1, numSingleOrdinal, true}

var dictionary = struct {
	sync.RWMutex
	m map[string]number
}{
	m: map[string]number{
		// Direct
		"zero":      {0, 1, numDirect, false},
		"a":         {1, 1, numDirect, false},
		"ten":       {10, 1, numDirect, false},
		"eleven":    {11, 1, numDirect, false},
		"twelve":    {12, 1, numDirect, false},
		"thirteen":  {13, 1, numDirect, false},
		"fourteen":  {14, 1, numDirect, false},
		"forteen":   {14, 1, numDirect, false},
		"fifteen":   {15, 1, numDirect, false},
		"sixteen":   {16, 1, numDirect, false},
		"seventeen": {17, 1, numDirect, false},
		"eighteen":  {18, 1, numDirect, false},
		"nineteen":  {19, 1, numDirect, false},
		"ninteen":   {19, 1, numDirect, false},

		// Single
		"one":   {1, 1, numSingle, false},
		"two":   {2, 1, numSingle, false},
		"three": {3, 1, numSingle, false},
		"four":  {4, 1, numSingle, false},
		"five":  {5, 1, numSingle, false},
		"six":   {6, 1, numSingle, false},
		"seven": {7, 1, numSingle, false},
		"eight": {8, 1, numSingle, false},
		"nine":  {9, 1, numSingle, false},

		// Tens
		"twenty":  {20, 1, numTens, false},
		"thirty":  {30, 1, numTens, false},
		"forty":   {40, 1, numTens, false},
		"fourty":  {40, 1, numTens, false},
		"fifty":   {50, 1, numTens, false},
		"sixty":   {60, 1, numTens, false},
		"seventy": {70, 1, numTens, false},
		"eighty":  {80, 1, numTens, false},
		"ninety":  {90, 1, numTens, false},

		// Bigs
		"hundred":  {100, 1, numBig, false},
		"thousand": {1000, 1, numBig, false},
		"million":  {1000000, 1, numBig, false},
		"billion":  {1000000000, 1, numBig, false},
		"trillion": {1000000000000, 1, numBig, false},

		// Fractions
		"half":         {1, 2, numFraction, false},
		"halve":        {1, 2, numFraction, false},
		"halfs":        {1, 2, numFraction, false},
		"halves":       {1, 2, numFraction, false},
		"thirds":       {1, 3, numFraction, false},
		"fourths":      {1, 4, numFraction, false},
		"quarter":      {1, 4, numFraction, false},
		"quarters":     {1, 4, numFraction, false},
		"fifths":       {1, 5, numFraction, false},
		"sixths":       {1, 6, numFraction, false},
		"sevenths":     {1, 7, numFraction, false},
		"eighths":      {1, 8, numFraction, false},
		"nineths":      {1, 9, numFraction, false},
		"tenths":       {1, 10, numFraction, false},
		"elevenths":    {1, 11, numFraction, false},
		"twelfths":     {1, 12, numFraction, false},
		"thirteenths":  {1, 13, numFraction, false},
		"fourteenths":  {1, 14, numFraction, false},
		"fifteenths":   {1, 15, numFraction, false},
		"sixteenths":   {1, 16, numFraction, false},
		"seventeenths": {1, 17, numFraction, false},
		"eighteenths":  {1, 18, numFraction, false},
		"nineteenths":  {1, 19, numFraction, false},
		"twentieths":   {1, 20, numFraction, false},
		"thirtieths":   {1, 30, numFraction, false},
		"fourtieths":   {1, 40, numFraction, false},
		"fiftieths":    {1, 50, numFraction, false},
		"sixtieths":    {1, 60, numFraction, false},
		"seventieths":  {1, 70, numFraction, false},
		"eightieths":   {1, 80, numFraction, false},
		"ninetieths":   {1, 90, numFraction, false},
		"hundredths":   {1, 100, numFraction, false},
		"thousandths":  {1, 1000, numFraction, false},
		"millionths":   {1, 1000000, numFraction, false},
		"billionths":   {1, 1000000000, numFraction, false},
		"trillionths":  {1, 1000000000000, numFraction, false},

		// Direct Ordinals
		"zeroth":      {0, 1, numDirectOrdinal, true},
		"tenth":       {10, 1, numDirectOrdinal, true},
		"eleventh":    {11, 1, numDirectOrdinal, true},
		"twelfth":     {12, 1, numDirectOrdinal, true},
		"thirteenth":  {13, 1, numDirectOrdinal, true},
		"fourteenth":  {14, 1, numDirectOrdinal, true},
		"fifteenth":   {15, 1, numDirectOrdinal, true},
		"sixteenth":   {16, 1, numDirectOrdinal, true},
		"seventeenth": {17, 1, numDirectOrdinal, true},
		"eighteenth":  {18, 1, numDirectOrdinal, true},
		"nineteenth":  {19, 1, numDirectOrdinal, true},

		// Single Ordinals
		"first":   {1, 1, numSingleOrdinal, true},
		"second":  second, // see IncludeSecond
		"third":   {3, 1, numSingleOrdinal, true},
		"fourth":  {4, 1, numSingleOrdinal, true},
		"fifth":   {5, 1, numSingleOrdinal, true},
		"sixth":   {6, 1, numSingleOrdinal, true},
		"seventh": {7, 1, numSingleOrdinal, true},
		"eighth":  {8, 1, numSingleOrdinal, true},
		"ninth":   {9, 1, numSingleOrdinal, true},

		// Tens Ordiinals
		"twentieth":  {20, 1, numTensOrdinal, true},
		"thirtieth":  {30, 1, numTensOrdinal, true},
		"fourtieth":  {40, 1, numTensOrdinal, true},
		"fiftieth":   {50, 1, numTensOrdinal, true},
		"sixtieth":   {60, 1, numTensOrdinal, true},
		"seventieth": {70, 1, numTensOrdinal, true},
		"eightieth":  {80, 1, numTensOrdinal, true},
		"ninetieth":  {90, 1, numTensOrdinal, true},

		// Big Ordinals
		"hundredth":  {100, 1, numBigOrdinal, true},
		"thousandth": {1000, 1, numBigOrdinal, true},
		"millionth":  {1000000, 1, numBigOrdinal, true},
		"billionth":  {1000000000, 1, numBigOrdinal, true},
		"trillionth": {1000000000000, 1, numBigOrdinal, true},

		// Glue
		"and": {0, 0, numAnd, false},
		"&":   {0, 0, numAnd, false},
	},
}

// IncludeSecond toggles whether or not "second" should be included in the
// interpreted words. If true "second" will be read as "2nd", otherwise the
// word will be ignored. The default is set to true.
func IncludeSecond(include bool) {
	dictionary.Lock()
	defer dictionary.Unlock()
	if include {
		dictionary.m["second"] = second
	} else {
		delete(dictionary.m, "second")
	}
}

// LookupNumber safely accesses the dictionary for a number. The input string is
// case insensitive.
func lookupNumber(s string) (n number, ok bool) {
	dictionary.RLock()
	defer dictionary.RUnlock()
	n, ok = dictionary.m[strings.ToLower(s)]
	return
}
