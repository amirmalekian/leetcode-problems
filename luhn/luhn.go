package luhn

import (
	"strings"
)

func Valid(id string) bool {

	// newId := strings.Replace(id, " ", "", -1)
	myId := strings.ReplaceAll(id, " ", "")
	if len(myId) < 2 {
		return false
	}
	sum := 0
	for i := 0; i < len(myId); i++ {
		digit := myId[i] - '0'
		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += int(digit)
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}

	// newId := strings.Split(id, " ")
	// sum := 0
	// for _, item := range newId {
	// 	for j, char := range item {
	// 		digit := char - '0'
	// 		if j%2 == 0 {
	// 			digit *= 2
	// 			if digit > 9 {
	// 				digit -= 9
	// 			}
	// 		}
	// 		sum += int(digit)
	// 	}
	// }

	// if sum%10 == 0 {
	// 	return true
	// } else {
	// 	return false
	// }
}
