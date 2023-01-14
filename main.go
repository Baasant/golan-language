package main

import "fmt"

func main() {
	//var conferenceName = "Go Conference"
	conferenceName := "Go Conference" // is the same as var conferenceName = "Go Conference"
	var firstName string              // here when you need to make variable and after a whole assign a value,you must give the data type
	var lastName string
	var userTickets uint
	const conferencetTickets = 50 // this is constant can't change during the program
	// uint prevent the number from being -ve
	var remainingTickets uint = 50 // to calculate the remaining tickets
	// fmt.Println("Wellcome to", conferenceName, " booking app")
	fmt.Printf("Wellcome to %v  booking app\n", conferenceName)
	fmt.Println("We have total", conferencetTickets, "and we have remaining", remainingTickets)
	fmt.Println(" Get your tickets here to attend")

	//userName = "bassant"
	//userTickets = 2
	//fmt.Printf("User %v book %v tickets.\n", userName, userTickets)

	//ask user for their name
	//fmt.Println("Enter your ffirst name: ")

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //here we get the pointer (location where the variable will be in memory)
	//fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)

	// in go if you didn't use variable will give you an error
	// which help your code to be clean

	// to get the types of the variable use %T
	//fmt.Printf("connderanceTickets is %T\n", conferenceName)

	fmt.Println("Enter your sec name: ")
	fmt.Scan(&lastName)

	fmt.Println("How msny tickets you want")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf(" user name is %v %v booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	///////////////////////////////////////arrays & slices

}
