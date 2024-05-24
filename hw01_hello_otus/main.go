package main

import (
	"fmt"
	"regexp"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello, OTUS!"))

	matched, _ := regexp.MatchString(`I am here`, "I am there")
	fmt.Println(matched) // false

	matched1, _ := regexp.MatchString(`Banana`, "Banana")
	fmt.Println(matched1) // true

	matched2, _ := regexp.MatchString(`\bcat\b`, "black cat meow")
	fmt.Println(matched2) // true

	re, _ := regexp.Compile(`\bcat\b`)
	res := re.FindAllString("black cat meowcat cat one cat", -1)
	fmt.Println(res) // [cat cat]

	matched3, _ := regexp.MatchString(`.....`, "any trash with 5 chars")
	fmt.Println(matched3) // true

	matched4, _ := regexp.MatchString(`^I\nam\nhere$`, "I\nam\nhere")
	fmt.Println(matched4) // true

	matched5, _ := regexp.MatchString(`a\|b=c`, "a|b=c")
	fmt.Println(matched5) // true

	re1, _ := regexp.Compile(`\d+`)
	res1 := re1.FindAllString("A123AA455AAA2A89", -1)
	fmt.Println(res1) // [123  455  2  89]

	matched6, _ := regexp.Compile(`A`)
	res2 := matched6.FindAllString("AGGAA!!", -1)
	fmt.Println(res2)

	re3, _ := regexp.Compile(`<.*>`)
	res3 := re3.FindAllString("<p><b>Golang</b> <i>VS</i> <b>Python</b></p>", -1)
	fmt.Println(res3) // [<p><b>Golang</b> <i>VS</i> <b>Python</b></p>] (len=1) :(

	re4, _ := regexp.Compile(`<.*?>`)
	res4 := re4.FindAllString("<p><b>Golang</b> <i>VS</i> <b>Python</b></p>", -1)
	fmt.Println(res4) // [<p>  <b>  </b>  <i>  </i>  <b>  </b>  </p>] :)

	re5, _ := regexp.Compile(`.(\d+)`)
	res5 := re5.FindAllStringSubmatch("Funny1 2020 ye12ar", -1)
	fmt.Println(res5) // [[y1  1] [ 2020  2020] [e12  12]]

}
