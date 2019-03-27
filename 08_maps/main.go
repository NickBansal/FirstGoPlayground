package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Nick"] = "nick@gmail.com"
	emails["Tom"] = "tom@gmail.com"

	fmt.Println(emails)
	delete(emails, "Bob")
	fmt.Println(emails)
}
