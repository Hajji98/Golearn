package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2

}
func main() {
	fmt.Println(concat("ali, ", "happy birthday"))
	fmt.Println(concat("Hajji, ", "hope that tesla thing works out"))
	fmt.Println(concat("go, ", "is fantastic"))
}
