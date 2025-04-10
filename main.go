package main

import "fmt"

func main() {
	m := map[string]string{
		"PurpleSchool": "purpleschool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"])
	m["PurpleSchool"] = "https://purpleschool.ru"
	fmt.Println(m["PurpleSchool"])
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	delete(m, "Y") // delete not exist
	fmt.Println(m["Y"]) // empty
	fmt.Println(m)
}