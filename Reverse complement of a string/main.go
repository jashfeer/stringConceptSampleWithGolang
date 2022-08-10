

package main

import "fmt"

func ReverseComplement(s string) string {
	return Reverse(Complement(s))
}
func Reverse(s string) string {
	s2:=""
	for i:=0;i<len(s);i++{
		s2 +=string(s[len(s)-1-i])
	}
	return s2
}
func Complement(s string) string {
	s2:=""
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			s2 += "X"
		case 'B':
			s2 += "Y"
		case 'C':
			s2 += "Z"
		}
		 
	}

	return s2
}

func main() {
fmt.Println(ReverseComplement("ABCBACD"))
}
