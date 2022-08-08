package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func matchBetweenStrs(str string, head_token string, end_token string) []string {
	r, _ := regexp.Compile(head_token + "\\s*(.*?)\\s*" + end_token)
	matches := r.FindAllStringSubmatch(str, -1)
	var res []string
	for _, group := range matches {
		if len(group[1]) != 0 {
			res = append(res, strings.ReplaceAll(group[1], "&#039;", ""))
		}
	}
	return res
}

func getUsers() []string {
	resp, err := http.Get("https://myanimelist.net/users.php")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	head_token := "<a href=\"/profile/"
	end_token := "\"><img class"
	users := matchBetweenStrs(string(body), head_token, end_token)
	return users
}

/* figure out the status of each of the anime
 * (completed has highest corr, then watching, then plan to watch)
 * (incorp rating?)
 */
func getAnimeList(username string) []string {
	uri := fmt.Sprintf("https://myanimelist.net/animelist/%s", username)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	head_token := "&quot;anime_title_eng&quot;:&quot;"
	end_token := "&quot;,&quot;anime_num_episodes"
	eng_anime_names := matchBetweenStrs(string(body), head_token, end_token)
	return eng_anime_names
}

func main() {
	// mylist := getAnimeList("scoobertDoobert")
	// for i, anime := range mylist {
	// 	fmt.Println(i, anime)
	// }
	users := getUsers()
	for i, user := range users {
		fmt.Println(i, user)
		a := getAnimeList(user)
		for i, anime := range a {
			fmt.Println(i, anime)
		}
	}
}
