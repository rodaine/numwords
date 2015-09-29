package numwords

import (
	"errors"
	"strings"
)

var (
	// ErrNoNumbers is returned if the a value for numbers is requested but there are
	// no number values inside of it
	ErrNoNumbers = errors.New("the input contains no number values")

	// ErrManyNumbers is returned if the value for numbers is requested but there are
	// more than one number values inside of it
	ErrManyNumbers = errors.New("the input contains more than one number")

	// ErrNonNumber is returned if ParseInt or ParseFloat encounters a non-number in
	// the input string.
	ErrNonNumber = errors.New("the string contains a non-number")
)

type numbers []number

// Pattern returns the string that represents the types of numbers in it
func (ns numbers) pattern() (p string) {
	for _, n := range ns {
		p += n.typ.String()
	}
	return
}

// Strings gets the string representations of each contained number
func (ns numbers) strings() []string {
	buf := make([]string, len(ns))
	for i, n := range ns {
		buf[i] = n.String()
	}
	return buf
}

// String returns the space separated string representation of the numbers
func (ns numbers) String() string {
	return strings.Join(ns.strings(), " ")
}

// Float returns a single value for the post-reduced numbers. If the length of
// numbers is not one, an error is returned instead
func (ns numbers) Float() (float64, error) {
	if len(ns) == 0 {
		return -1, ErrNoNumbers
	} else if len(ns) > 1 {
		return -1, ErrManyNumbers
	}

	return ns[0].Value(), nil
}

// Int returns a single integer value for the post-reduced numbers, similar to
// number.Float.
func (ns numbers) Int() (int, error) {
	out, err := ns.Float()
	return int(out), err
}

func (ns numbers) flush(s []string) []string {
	if len(ns) > 0 {
		s = append(s, reduce(ns).strings()...)
	}
	return s
}

// Reduce destructivley converts the numbers set to its minimal form based on
// known numeric patterns. NB: the input slice will be modified significantly
func reduce(ns numbers) numbers {
	for found := true; found; {
		found = false
		pattern := ns.pattern()
		for _, p := range patterns {
			if idx := strings.LastIndex(pattern, p); idx >= 0 {
				found = true
				ns = patternHandlers[p](ns, idx)
				break
			}
		}
	}
	return ns
}
