package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
)

func Top10(str string) []string {

	//regexp.MustCompile
	//strings.Split
	//strings.Fields
	//sort.Slice

	re, _ := regexp.Compile(`\bcat\b`)
	res := re.FindAllString("cats black cat meowcat cat one cat", -1)
	fmt.Println(res) //
	return nil
}
