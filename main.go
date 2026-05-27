package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"test-vscode-go/functions"
	"test-vscode-go/greeting"
)

func main() {

	// s1 := []int{3, 4, 6}
	// s2 := s1
		greeting.SayHelloToWorld()


		myNums := []int{2, 7, 9, 15}

	myResult := TwoSum(myNums, 9)
	fmt.Println(myResult)

	

	score := 0

	pointer := &score

	fmt.Println("THIS IS A POINTER ADDRESS: ", pointer)
	fmt.Println("THIS IS POINTER VALUE: ", *pointer)

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

	for i, name := range players {
		fmt.Printf("Player #%d: -> %s ", i+1, name)

		switch name {
		case "Vika":
			fmt.Println("Master of the Game")
		case "Robot":
			fmt.Println("Just a Bot")
		default:
			fmt.Println("Simple player")

		}
	}

	// fmt.Println(equal(s1, s2))
	weekday := "ааа"

	switch weekday {
	case "Friday":
		fmt.Println("Today is not a Monday")
	case "Monday":
		fmt.Println("Today is Monday")
	default:
		fmt.Println("not a weekday")
	}
	fmt.Println("Hello World")
	fmt.Println("Generated random", rand.Intn(10))

	if rand.Intn(4) == 1 {
		fmt.Println("I generated random number")
	} else {
		fmt.Println("I did not hit 1")
	}

	result := aboba(0, 10, 0)
	fmt.Println(result)

	switch {
	case result < 10:
		fmt.Println("result is smaller than 10")
	case result > 10:
		fmt.Println("Result is bigger than 10")
	default:
		fmt.Println("I dont know")
	}

	squaredNum := functions.Squre(2)
	fmt.Println(squaredNum)

	s := 5
	t := "ahahhaa"

	fmt.Println(ahaha(s, t))
	fmt.Println(s)
	fmt.Println(t)


	userBob := User {
		Name: "Bob",
		Age: 40,
		PhoneNumber: "88888888",
		Address: Address{
			City: "London",
			Street: "Shier",
		},
	}

	fmt.Println("Created new user: ->", userBob)
	fmt.Printf("%+v\n", userBob)


	userBob.changeName("Sara")

	fmt.Println(userBob)

	err := userBob.changeNameWithError("")
	
	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("")


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

func hello() {}

func aboba(a, b, c int) int {
	return a + b + c
}

func ahaha(s int, t string) (int, string) {
	s = 100
	t = "bbbb"

	return s, t
}
func qyqyqyq(s *int, t *string) (int, string) {
	*s = 100
	*t = "bbbb"

	return *s, *t
}

func integetToString1(a, b int) string {
	return strconv.Itoa(a + b)
}

func integetToString2(a, b int) string {
	return fmt.Sprint(a+b)
}

func database(){
	defer func(){}()
}