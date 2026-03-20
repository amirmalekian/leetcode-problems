package raindrops

import (
	"bytes"
	"strconv"
)

var translation = []struct {
	Num  int
	Word string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(amount int) string {
	var b bytes.Buffer

	for _, trans := range translation {
		if amount%trans.Num == 0 {
			b.WriteString(trans.Word)
		}
	}

	if b.String() == "" {
		b.WriteString(strconv.Itoa(amount))
	}

	return b.String()
}
