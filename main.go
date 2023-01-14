package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var userName string // here when you need to make variable and after a whole assign a value,you must give the data type
	var userTickets int
	const conferencetTickets = 50 // this is constant can't change during the program
	var remainingTickets = 50     // to calculate the remaining tickets
	// fmt.Println("Wellcome to", conferenceName, " booking app")
	fmt.Printf("Wellcome to %v  booking app\n", conferenceName)
	fmt.Println("We have total", conferencetTickets, "and we have remaining", remainingTickets)
	fmt.Println(" Get your tickets here to attend")

	userName = "bassant"
	userTickets = 2
	fmt.Printf("User %v book %v tickets.\n", userName, userTickets)

	// in go if you didn't use variable will give you an error
	// which help your code to be clean

	// to get the types of the variable use %T
	fmt.Printf("connderanceTickets is %T\n", conferenceName)

}
