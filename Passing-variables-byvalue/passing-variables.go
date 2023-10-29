package main

import "fmt"

func main() {
	SendsSoFar := 430
	const SendsToAdd = 25

	SendsSoFar = incrementSends(SendsSoFar, SendsToAdd)
	fmt.Println("you've sent", SendsSoFar, "Messages")

}

func incrementSends(SendsSoFar, sendsToAdd int) int {
	SendsSoFar = SendsSoFar + sendsToAdd
	return SendsSoFar
}
