package main

import (
	"fmt"
)

type user struct {
	name string
	id   int
}

// Print prints user info
func (u *user) Print() {
	fmt.Printf("Name: %v, ID: %v\n", u.name, u.id)
}

var people = []user{
	{name: "John", id: 1},
	{name: "Jacob", id: 2},
	{name: "Josh", id: 3},
	{name: "Eve", id: 4},
	{name: "Daniel", id: 5},
	{name: "John", id: 6},
	{name: "Wes", id: 7},
	{name: "John", id: 8},
	{name: "Eve", id: 9},
}

func main() {
	m := groupUsersByName(people)
	fmt.Println(m)

	dupes := findUsersWithMoreThanOneID(m)
	fmt.Println("dupes:", dupes)
}

func groupUsersByName(users []user) map[string][]int {
	m := make(map[string][]int)

	for _, v := range users {
		ids, found := m[v.name]
		// if not found instantiate map
		if !found {
			m[v.name] = append(ids, v.id)
		}
		m[v.name] = append(ids, v.id)
	}

	return m
}

func findUsersWithMoreThanOneID(m map[string][]int) []string {
	var usersWithMultipleIds []string

	for k, v := range m {
		if len(v) > 1 {
			r := fmt.Sprintf("%v:%v", k, v)
			usersWithMultipleIds = append(usersWithMultipleIds, r)
		}
	}

	return usersWithMultipleIds
}

//You have a text file with name and id comma separated. Print out all the names which are duplicates and their corresponding ids.
// first step is to build the data object that will be passed to the function
//
//eg.
//Name, Id
//John, 1
//Jacob, 2
//Josh, 3
//Eve, 4
//Daniel, 5
//John, 6
//Wes, 7
//John, 10
//
//Output:
//John 1, 6, 10
