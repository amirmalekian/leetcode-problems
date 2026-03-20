package main

import (
	// "firstapp/student"
	// "firstapp/differenceOfSquares"
	// "firstapp/luhn"
	"fmt"
	"math"
	"strings"
	"unicode"
)

func add(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	if x := sqrt(float64(4)); x == "2" {
		fmt.Println("hhhhhh")
	}

	return fmt.Sprint(math.Sqrt(x))
}

// func SendMessageWithName(name string) (str string) {
// 	if name != "" {
// 		str = "One for " + name + ", one for me."
// 	} else {
// 		str = "One for you, one for me."
// 	}
// 	return
// }

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me. My Age is %d", name, 25)
	// return "One for " + name + ", one for me."
}

func ScoreOf(word string) (score int) {
	for _, rune := range word {
		switch unicode.ToUpper(rune) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}
	return
}

func scoreOf(c rune) int {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}

func ScrabbleScore(word string) (sum int) {

	letters := make(map[rune]int)
	var points = map[rune]int{
		'A': 1,
		'B': 3,
		'C': 3,
		'D': 2,
	}
	for _, l := range word {
		sum += points[unicode.ToUpper(l)]
	}
	for _, l := range []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'} {
		letters[l] = 1
	}

	for _, l := range []rune{'D', 'G'} {
		letters[l] = 2
	}

	for _, l := range []rune{'B', 'C', 'M', 'P'} {
		letters[l] = 3
	}

	for _, l := range []rune{'F', 'H', 'V', 'W', 'Y'} {
		letters[l] = 4
	}

	letters['K'] = 5

	for _, l := range []rune{'J', 'X'} {
		letters[l] = 8
	}

	for _, l := range []rune{'Q', 'Z'} {
		letters[l] = 10
	}

	for _, l := range strings.ToUpper(word) {
		fmt.Println("char is", l)
		if v, ok := letters[l]; ok {
			sum += v
		}
	}

	return
}

func main() {
	// type student struct {
	// 	Name string
	// 	ID   string
	// }
	// nill map
	// var mayMap map[string]string

	// initial map
	// var myMap =map[string]string{}
	// ideomatic
	// myMap := make(map[string]string)
	// var myMap = map[string]string{
	// 	"a": "hello",
	// 	"b": "chalghoz",
	// }

	// a := make(map[string]int)
	// a["h"] = 3
	// if v, ok := a["hh"]; ok {
	// 	fmt.Println(v, ok)
	// } else {
	// 	fmt.Println("is not exist")
	// }
	// fmt.Println(ShareWith(""))
	// fmt.Println(SendMessageWithName(""))
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// fmt.Println(add(5, 4))
	// fmt.Println(split(17))
	// sum := 0
	// for i := 0; i < 10; i += 2 {
	// 	sum += i
	// }
	// fmt.Println(sqrt(2), sqrt(-4))

	// j := 0
	// for ; j < 100 ; {
	// 	j+=1
	// }
	// for j < 100 {

	// }
	// for {

	// }
	// fmt.Println(sum)

	// var u student.Student
	// u.Age = 12
	// u.FirstName = "Amirhossein"
	// // u.isActive = true
	// fmt.Println(u.Age)
	// fmt.Println(u.FirstName)

	// var j = student.Student{
	// 	ID:        1,
	// 	FirstName: "Hossein",
	// 	Age:      25,
	// 	LastName:  "Hesami",
	// }

	// fmt.Println(j.Age)
	// fmt.Println(j.FirstName)

	// var z = student.Student {12, "00110", 24, "first", "last"}

	// fmt.Println(z)
	// "My name is amirhossein , and I'm 25 years old. My National code is 124124"

	// fmt.Printf("My nane is %s, and I'm %d years old. My National code is %s\n", z.FirstName, z.Age, z.NationNumber)
	// fmt.Println("HHHH")

	// Array
	// var a [6]int

	// a = [6]int{1, 2, 3, 4, 5, 6}
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	// Slice
	// var b []int
	// builtin function

	// b = append(b, 5)
	// fmt.Println(b)
	// fmt.Println(cap(b), len(b))
	// c := make([]int, 5, 10)
	// fmt.Println(cap(c))
	// fmt.Println(len(c))

	// Exercism
	// fmt.Println(diffsquares.SquareOfSum(10))
	// fmt.Println(diffsquares.SumOfSquares(10))
	// fmt.Println(diffsquares.Difference(10))

	// fmt.Println(luhn.Valid("4539 3195 0343 6467"))

	// score := ScrabbleScore("cabbage")
	// fmt.Println("score is", score)
	// word := "cabbage"
	// score := 0
	// score2 := 0
	// for _, c := range strings.ToLower(word) {
	// 	score += scoreOf(rune(c))
	// 	score2 += scoreOf(c)
	// }
	// fmt.Println("score 1 is :", score)
	// fmt.Println("score 2 is: ", score2)
}
