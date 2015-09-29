package numwords

type patternHandler func(ns numbers, idx int) numbers

// Patterns describes all the salient number patterns that could be in a
// numbers set. These patterns are evaluated in the order listed here until all
// options are exhausted.
var patterns = []string{
	// tens
	"ts", // twenty three => 23

	// big
	"bdb", // million eighteen thousand => 1018000
	"db",  // eleven hundred            => 1100
	"sb",  // one hundred               => 100
	"btb", // million twenty thousand   => 1020000
	"tb",  // twenty thousand           => 20000
	"bd",  // hundred eleven            => 111
	"bsb", // thousand two hundred      => 1200
	"bs",  // hundred one               => 101
	"bt",  // hundred twenty            => 120
	"bbb", // million hundred thousand  => 1100000
	"bb",  // hundred thousand          => 100000

	// direct
	"dd", // nineteen ten    => 1910
	"dt", // nineteen eighty => 1980
	"td", // twenty fifteen  => 2015

	// fraction
	"df", // fifteen twentieths  => 0.75
	"sf", // three fourths       => 0.75
	"tf", // thirty fourtieths   => 0.75
	"bf", // hundred thousandths => 0.1

	// ordinals that could possibly be singluar fractions
	"dD", // a tenth       => 0.1  || fifteen tenth     => 15 10th
	"dS", // a fourth      => 0.25 || fifteen fourth    => 15 4th
	"dT", // a twentieth   => 0.05 || fifteen twentieth => 15 20th
	"dB", // a hundredth   => 0.01 || fifteen hundredth => 1500th
	"sD", // one tenth     => 0.1  || two tenth         => 2 10th
	"sS", // one fourth    => 0.25 || two fourth        => 2 4th
	"sT", // one twentieth => 0.05 || two twentieth     => 15 20th
	"sB", // one hundredth => 0.01 || fifteen hundredth => 1500th

	// all other ordinals
	"tS", // twenty first       => 21st
	"tB", // twenty thousandth  => 20000th
	"bS", // hundred first      => 101st
	"bB", // hundred thousandth => 100000th

	// glue
	"d&f", // zero and a half      => 0.5
	"s&f", // two and a half       => 2.5
	"t&f", // twenty and a half    => 20.5
	"b&f", // hundred and a half   => 100.5
	"&",   // 100 and 50 => 100 50 => 150
}

var patternHandlers = map[string]patternHandler{
	"tS": add,
	"bS": add,

	"df": multiply,
	"sf": multiply,
	"tf": multiply,
	"bf": multiply,
	"tB": multiply,
	"bT": multiply,

	"ts": combine,
	"db": combine,
	"sb": combine,
	"tb": combine,
	"bd": combine,
	"bs": combine,
	"bt": combine,
	"bb": combine,
	"bB": combine,

	"bbb": combineToLowest,
	"bdb": combineToLowest,
	"bsb": combineToLowest,
	"btb": combineToLowest,

	"dd": yearOrDone,
	"dt": yearOrDone,
	"td": yearOrDone,

	"dD": fractionOrDone,
	"dS": fractionOrDone,
	"dT": fractionOrDone,
	"sD": fractionOrDone,
	"sS": fractionOrDone,
	"sT": fractionOrDone,

	"dB": fractionOrCombine,
	"sB": fractionOrCombine,

	"d&f": addAnd,
	"s&f": addAnd,
	"t&f": addAnd,
	"b&f": addAnd,

	"&": drop,
}

// Done flags the number at the given index as "done" and no longer
// will attempt to resolve it as part of a pattern
func done(ns numbers, idx int) numbers {
	ns[idx].typ = numDone
	return ns
}

// Drop removes the number at the specified index from the set of numbers
func drop(ns numbers, idx int) numbers {
	out := ns[:idx]
	out = append(out, ns[idx+1:]...)
	return out
}

// Add merges two numbers by adding their values together
// This typically occurs when a large number proceeds a smaller one.
func add(ns numbers, idx int) numbers {
	a := ns[idx]
	b := ns[idx+1]

	ns[idx].numerator = a.numerator*b.denominator + a.denominator*b.numerator
	ns[idx].denominator = a.denominator * b.denominator
	ns[idx].typ = maxType(a.typ, b.typ)
	ns[idx].ordinal = b.ordinal

	return drop(ns, idx+1)
}

// AddAnd adds together two numbers split by an (arbitrary) "and" number.
func addAnd(ns numbers, idx int) numbers {
	ns = drop(ns, idx+1)
	return add(ns, idx)
}

// Multiply merges two numbers by multiplying their values together
// This typically occurs when a small number proceeds a larger one.
func multiply(ns numbers, idx int) numbers {
	a := ns[idx]
	b := ns[idx+1]

	ns[idx].numerator = a.numerator * b.numerator
	ns[idx].denominator = a.denominator * b.denominator
	ns[idx].typ = maxType(a.typ, b.typ)
	ns[idx].ordinal = b.ordinal

	return drop(ns, idx+1)
}

// Combine merges two numbers by addition or multiplication depending
// on the relative values of the numbers
func combine(ns numbers, idx int) numbers {
	a := ns[idx]
	b := ns[idx+1]

	if a.numerator > b.numerator {
		return add(ns, idx)
	}

	return multiply(ns, idx)
}

// CombineToLowest combines the two lowest adjacent values in a triple
// of number values. This typically occurs when a number is sandwiched
// between two large values.
func combineToLowest(ns numbers, idx int) numbers {
	l := ns[idx]
	r := ns[idx+2]

	if l.numerator <= r.numerator {
		return combine(ns, idx)
	}

	return combine(ns, idx+1)
}

// YearOrDone potentially combines two double-digit consecutive values
// if they appear to be a colloquial year (eg, ninteen eighty eight => 1988).
// Acceptable years captured by this function range from 1010 to 2100 (exclusive)
// and does not include the first decade of each century (eg, 2000-2009). If the
// heuristic isn't satisfied, the second number in the potential year is marked
// as done to advance evaluation. Likewise, on successful conversion, the date
// is marked done due to its semantic change (from arbitrary number to year).
//
// TODO: capture first decades with "oh": ninteen oh eight => 1908
func yearOrDone(ns numbers, idx int) numbers {
	a := ns[idx]
	b := ns[idx+1]

	if a.numerator > 10 && a.numerator <= 20 && b.numerator >= 10 && b.numerator < 100 {
		ns[idx].numerator *= 100
		ns = add(ns, idx)
		return done(ns, idx)
	}

	return done(ns, idx+1)
}

// FractionOr builds a patternHandler that converts ordinals to 1-numerator
// fractions based on context: one hundredth => 0.001 vs. two hundredth => 200th
// If the heuristic fails, the passed in patternHandler is applied instead.
func fractionOr(ph patternHandler) patternHandler {
	return func(ns numbers, idx int) numbers {
		if ns[idx].numerator == 1 {
			ns[idx+1].ordinal = false
			ns[idx+1].denominator = ns[idx+1].numerator
			ns[idx+1].numerator = 1
			ns[idx+1].typ = numFraction
			return multiply(ns, idx)
		}
		return ph(ns, idx)
	}
}

var (
	// FractionOrDone marks the number done if it is not part of a singular fraction value
	fractionOrDone = fractionOr(done)

	// FractionOrCombine combines two numbers if it is not a singular fraction value
	fractionOrCombine = fractionOr(combine)
)
