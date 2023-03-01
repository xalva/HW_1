package main

import (
	"fmt"
	"strings"
)

type user struct {
	name  string
	age   uint8
	email string
}

var usersMap map[string]user

func (u user) String() string {
	if u == (user{}) {
		return ("Not found!")
	}
	return fmt.Sprintf("%v (%d) %v", u.name, u.age, u.email)
}

func main() {
	users := []user{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}

	FillUsersMap(users)
	findedUser := FindUserByNameInSensitive("JoHn")
	fmt.Println(findedUser)

}

func FillUsersMap(u []user) {
	usersMap = make(map[string]user)
	for _, data := range u {
		usersMap[strings.ToUpper(data.name)] = data
	}
}
func FindUserByNameInSensitive(s string) user {
	return usersMap[strings.ToUpper(s)]
}
