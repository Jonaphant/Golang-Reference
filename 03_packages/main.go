package main

import (
	"fmt"
	"math"

	// This imports the package that we created containing a function called Reverse. This path can look different based on the way you've structured your Go workspace.
	"github.com/jonaphant/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(4))

	//This is how we used the imported package
	fmt.Println(strutil.Reverse("olleh"))

}
