package internal

import (
	"strings"
)

type User struct {
	Name  string
	Age   uint8
	Email string
}

var usersMap map[string]User

func FillUsersMap(u []User) {
	usersMap = make(map[string]User)
	for _, data := range u {
		usersMap[strings.ToUpper(data.Name)] = data
	}
}
func FindUserByNameInSensitive(s string) User {
	return usersMap[strings.ToUpper(s)]
}
