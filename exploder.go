package numwords

import (
	"bufio"
	"strings"
)

func explode(s string) (out []string) {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, ",", "", -1)

	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return
}
