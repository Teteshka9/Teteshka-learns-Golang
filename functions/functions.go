package functions


import (

	"fmt"
	"strconv"
)

func Squre(x int) int { return x * x}


func Sum (nums[]int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
} 

func equal(a, b []int) bool {
	if len(a) != len(b) {

		return false
	}

	for i := range a {

		if a[i] != b[i] {

			return false
		}
	}

	return true
}

func SayHelloWorld() string {
	return "My greeting function"
}

func Hello() {}

func Aboba(a, b, c int) int {
	return a + b + c
}

func Ahaha(s int, t string) (int, string) {
	s = 100
	t = "bbbb"

	return s, t
}
func Qyqyqyq(s *int, t *string) (int, string) {
	*s = 100
	*t = "bbbb"

	return *s, *t
}

func IntegetToString1(a, b int) string {
	return strconv.Itoa(a + b)
}

func IntegetToString2(a, b int) string {
	return fmt.Sprint(a + b)
}

func Database() {
	defer func() {}()
}
