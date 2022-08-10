package main

import (
    "strconv"
    "fmt"
)

func main() {
	i:=123
    s := strconv.Itoa(i)
	fmt.Printf("value: %v,type: %T",i ,i)
	fmt.Println()
    fmt.Printf("value: %v,type: %T",s ,s)
	fmt.Println()
}