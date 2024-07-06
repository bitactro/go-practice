package main

import "fmt"

func main() {
	username := ""
	password := ""
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)

	if username == "ankit" && password == "loginme" {
		fmt.Println("Authenticaion successful.")
	} else {
		fmt.Println("Authentication failed.")
	}
}
