package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name        string          `json:"username"`
	Password    string          `json:"-"`
	Permissions map[string]bool `json:"perms,omitempty"`
}

type json_user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"abc", "1234", nil},
		{"god", "42", map[string]bool{"admin": true}},
		{"devil", "666", map[string]bool{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))

	var json_users []json_user

	errr := json.Unmarshal(out, &json_users)
	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println(json_users)

}
