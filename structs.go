package main

import "errors"


type User struct {
	Name string 
	Age int
	PhoneNumber string
	Address 
	isClose bool
}

type Person struct {
	Name string 
	Rating float64
	Premium bool
}

type Address struct {
	City string  
	Street string
}

func (u *User) changeName (newName string)  {


	if newName != "" {
		u.Name = newName
	}

}

func (u *User) changeNameWithError (newName string) error {

	if newName == "" {
		return errors.New("Can not assign an empty string to name")
	}
	
	u.Name = newName

	return nil
}

func (u *User) ChangeAge (newAge int) error {
	if newAge < 0 {
		return errors.New("Age can not be less than zero")
	} 

	u.Age = newAge

	return nil
}

 
func TwoSum (nums []int, target int) []int {

	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		needed := target - num 
		if index, ok := seen[needed]; ok {
			return []int{index, i}
		}
		seen[num] = i
	}

	return nil
} 

// func NewUser(Name string, age int, PhoneNumber string, Address struct, IsClose bool) 