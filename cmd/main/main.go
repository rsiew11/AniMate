package main

import (
	"fmt"

	"github.com/rsiew11/aniMate/pkg/apigetters"
)

func main() {
	fmt.Println("hello world")
	mylist := apigetters.GetAnimeList("scoobertDoobert")
	for i, anime := range mylist {
		fmt.Println(i, anime)
	}
	users := apigetters.GetUsers()
	for i, user := range users {
		fmt.Println(i, user)
		a := apigetters.GetAnimeList(user)
		for i, anime := range a {
			fmt.Println(i, anime)
		}
	}
}
