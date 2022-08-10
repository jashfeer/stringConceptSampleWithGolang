package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hi jashfeer")
	fmt.Println(strings.Count("jashfeer", "e"))
	fmt.Println(strings.Count("abadababa", "aba"))
	fmt.Println("jashfeer " + "abdullah")
	s := "jone honayi"
	fmt.Println("length of s =", len(s))
	fmt.Println("2nd chare of s is ", s[1])
	fmt.Println("2nd chare of s is ", string(s[1]))
	fmt.Println(s)
	fmt.Println(s[2:7])
	subStrOfS:=s[3:9]
	fmt.Println(subStrOfS)
	fmt.Println(strings.Contains("hi yolo","hi"))


}
