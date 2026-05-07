package main

import (
	"fmt"
)

func main() {

s1 := []int{3, 4, 6}
s2 := s1


var number int 
var text string 
fmt.Println(number)
fmt.Println(text)

score := 0

for i := 0; i < 12; i++ {
	score++


switch {
case score > 10: 
	fmt.Println("Good boy")
case score <= 10:
	fmt.Println("Score is not enough")
default:
	fmt.Println("Play better")
}
}

players := []string{"Vika", "Alex", "Robot", "Max"}

for i, name := range players{
	fmt.Printf("Player #%d: -> %s ", i+1, name)

	switch name {
	case "Vika":
		fmt.Println("Master of the Game")	
	case "Robot":
		fmt.Println("Just a Bot")
	default: 
		fmt.Println("Simple player")

}
fmt.Println(equal(s1, s2))
}
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

