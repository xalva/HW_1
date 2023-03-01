package main

import (
	"HW_1/internal"
	"fmt"
)

type user internal.User

func (u user) String() string {
	if u == (user{}) {
		return ("Not found!")
	}
	return fmt.Sprintf("%v (%d) %v", u.Name, u.Age, u.Email)
}

func main() {
	users := []internal.User{
		{
			Name:  "Mike",
			Age:   32,
			Email: "mike@gmail.com",
		},
		{
			Name:  "John",
			Age:   54,
			Email: "john@gmail.com",
		},
		{
			Name:  "Abobus",
			Age:   2,
			Email: "abobus@gmail.com",
		},
	}

	internal.FillUsersMap(users)
	findedUser := internal.FindUserByNameInSensitive("JoHn")
	fmt.Println(findedUser)

}
