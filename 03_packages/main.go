package main

import (
	"fmt"
	"math"

	"github.com/NickBansal/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.9876865))
	fmt.Println(math.Ceil(2.9876865))
	fmt.Println(math.Sqrt(169))
	fmt.Println(strutil.Reverse("olleH"))
}
