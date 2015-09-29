package numwords

import "fmt"

func Example() {
	s := "I've got three apples and two and a half bananas"
	fmt.Println(ParseString(s))

	s = "My chili won second place at the county fair"
	fmt.Println(ParseString(s))

	i, _ := ParseInt("fourteen ninety two")
	fmt.Println(i)

	f, _ := ParseFloat("eight and three quarters")
	fmt.Println(f)

	// Output:
	// I've got 3 apples and 2.5 bananas
	// My chili won 2nd place at the county fair
	// 1492
	// 8.75
}

func ExampleParseString() {
	s := "I've got three apples and two and a half bananas"
	fmt.Println(ParseString(s))

	// Output:
	// I've got 3 apples and 2.5 bananas
}

func ExampleParseFloat() {
	f, _ := ParseFloat("eight and three quarters")
	fmt.Println(f)

	// Output:
	// 8.75
}

func ExampleParseInt() {
	i, _ := ParseInt("fourteen ninety two")
	fmt.Println(i)

	// Output:
	// 1492
}

func ExampleIncludeSecond() {
	s := "My chili won second place at the county fair"
	fmt.Println(ParseString(s))

	s = "One second ago"
	IncludeSecond(false)
	fmt.Println(ParseString(s))

	// Output:
	// My chili won 2nd place at the county fair
	// 1 second ago
}
