package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Nick"] = "nick@gmail.com"
	emails["Tom"] = "tom@gmail.com"

	arrSlice := []int{33, 42, 5, 709, 43, 121, 5, 786}
	total := 0
	for _, id := range arrSlice {
		total += id
	}
	fmt.Println("Sum =", total)

	for k, v := range emails {
		fmt.Println(k, ":", v)
		fmt.Printf("%s: %s\n", k, v)
	}

}
