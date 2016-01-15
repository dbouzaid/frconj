package main

import (
	"fmt"
)

func main() {
	var inputVerb string
	fmt.Print("Please type a french verb: ")
	fmt.Scan(&inputVerb)

	if inputVerb[len(inputVerb)-2:] == "er" {
		conjER(inputVerb)
	} else if inputVerb[len(inputVerb)-2:] == "re" {
		conjRE(inputVerb)
	} else if inputVerb[len(inputVerb)-2:] == "ir" {
		conjIR(inputVerb)
	} else {
		fmt.Println("Not found")
	}
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

func conjRE(input string) {
	je := input[:len(input)-2] + "s"
	tu := input[:len(input)-2] + "s"
	il := input[:len(input)-2]
	nous := input[:len(input)-2] + "ons"
	vous := input[:len(input)-2] + "ez"
	ils := input[:len(input)-2] + "ent"

	fmt.Printf("Je: \t\t%s \nTu: \t\t%s \nIl/Elle/On: \t%s \nNous: \t\t%s \nVous: \t\t%s \nIls/Elles: \t%s \n", je, tu, il, nous, vous, ils)
}

func conjIR(input string) {
	je := input[:len(input)-2] + "is"
	tu := input[:len(input)-2] + "is"
	il := input[:len(input)-2] + "it"
	nous := input[:len(input)-2] + "issons"
	vous := input[:len(input)-2] + "issez"
	ils := input[:len(input)-2] + "issent"

	fmt.Printf("Je: \t\t%s \nTu: \t\t%s \nIl/Elle/On: \t%s \nNous: \t\t%s \nVous: \t\t%s \nIls/Elles: \t%s \n", je, tu, il, nous, vous, ils)
}
