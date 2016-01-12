package main

import (
	"fmt"
)

func main() {
	var inputVerb string
	fmt.Print("Please type a french ER verb: ")
	fmt.Scan(&inputVerb)
	conjER(inputVerb)
}

func conjER(input string) {
	je := input[:len(input)-2] + "e"
	tu := input[:len(input)-2] + "es"
	il := input[:len(input)-2] + "e"
	nous := input[:len(input)-2] + "ons"
	vous := input[:len(input)-2] + "ez"
	ils := input[:len(input)-2] + "ent"

	fmt.Printf("Je: \t\t%s \nTu: \t\t%s \nIl/Elle/On: \t%s \nNous: \t\t%s \nVous: \t\t%s \nIls/Elles: \t%s \n", je, tu, il, nous, vous, ils)
}
