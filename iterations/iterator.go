package iterations

import "strings"

func Repeat(stringToRepeat string, numberOfRepeat int) string {
	// or strings.Repeat(s, count)
	var repeated strings.Builder
	for i := 0; i < numberOfRepeat; i++ {
		repeated.WriteString(stringToRepeat)
	}
	return repeated.String()
}
