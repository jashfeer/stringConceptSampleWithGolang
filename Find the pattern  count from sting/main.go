package main

import "fmt"

func countPattern(text,pattern string)int{
	count:=0
	n:=len(text)
	k:=len(pattern)
	for i:=0;i<=n-k;i++{
		if pattern==text[i:i+k]{
			count++
		}
	}
	return count
}


func main(){
	var s,p string
fmt.Println("Enter the string")
fmt.Scanln(&s)
fmt.Println("Enter the pattern")
fmt.Scanln(&p)
fmt.Println("Count is = ",countPattern(s,p))
}