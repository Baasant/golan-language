package main

import (
	"fmt"
	"strings"
)

/*
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

		//////////////arrays/////////////////
		//var bookings [50]string

		///////////////////////////////////

		/////////slices//////////////
		var bookings []string //the main different that you havn't to add the length of the array
		//////////////////////////////

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

		///////////////array accessing element///////////
		//bookings[0] = firstName + " " + lastName
		/////////////////////////////////////////

		///////////////slices accessing element///////////
		bookings = append(bookings, firstName+" "+lastName)
		/////////////////////////////////////////

		////////////////////////array//////////////////////////////////////////////////
		//fmt.Printf("the whole array: %v \n", bookings)
		//fmt.Printf("the first value: %v \n", bookings[0])
		//fmt.Printf("array type :%T \n", bookings)
		//fmt.Printf("array length :%v \n", len(bookings))
		///////////////////////////////////////////////////////////////////////

		////////////////////////slices//////////////////////////////////////////////////
		fmt.Printf("the whole slice: %v \n", bookings)
		fmt.Printf("the first value: %v \n", bookings[0])
		fmt.Printf("array type :%T \n", bookings)
		fmt.Printf("array length :%v \n", len(bookings))
		///////////////////////////////////////////////////////////////////////

		fmt.Printf(" user name is %v %v booked %v tickets\n", firstName, lastName, userTickets)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		///////////////////////////////////////arrays & slices

}
*/
func main() {

	//var conferenceName = "Go Conference"

	const conferencetTickets = 50 // this is constant can't change during the program
	// uint prevent the number from being -ve
	var remainingTickets uint = 50 // to calculate the remaining tickets

	var bookings []string //the main different that you havn't to add the length of the array
	//////////////////////////////
	for {

		conferenceName := "Go Conference" // is the same as var conferenceName = "Go Conference"
		var firstName string              // here when you need to make variable and after a whole assign a value,you must give the data type
		var lastName string
		var userTickets uint
		fmt.Printf("Wellcome to %v  booking app\n", conferenceName)
		fmt.Println("We have total", conferencetTickets, "and we have remaining", remainingTickets)
		fmt.Println(" Get your tickets here to attend")

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //here we get the pointer (location where the variable will be in memory)
		//fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)

		fmt.Println("Enter your sec name: ")
		fmt.Scan(&lastName)

		fmt.Println("How msny tickets you want")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets

			///////////////slices accessing element///////////
			bookings = append(bookings, firstName+" "+lastName)
			/////////////////////////////////////////

			////////////////////////slices//////////////////////////////////////////////////
			fmt.Printf("the whole slice: %v \n", bookings)
			fmt.Printf("the first value: %v \n", bookings[0])
			fmt.Printf("array type :%T \n", bookings)
			fmt.Printf("array length :%v \n", len(bookings))
			///////////////////////////////////////////////////////////////////////

			//fmt.Printf(" user name is %v %v booked %v tickets\n", firstName, lastName, userTickets)
			fmt.Printf(" user name is %v booked %v tickets\n", firstName, userTickets)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) // split booking to first and sec name
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("the first name of the user is  %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("comeback next year")
				break
			}
		} else {
			fmt.Printf("we only have %v tickets\n", remainingTickets)
			continue
		}

		///////////////////////////////////////arrays & slices
	}

	//break instead of breaking the loop we will skip this iteration by adding continus

}
