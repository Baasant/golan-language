package main

import (
	"booking-app/helper"
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
/*
func main() {
	conferenceName := "Go Conference"
	greatusers(conferenceName)

	//var conferenceName = "Go Conference"

	const conferencetTickets = 50 // this is constant can't change during the program
	// uint prevent the number from being -ve
	var remainingTickets uint = 50 // to calculate the remaining tickets

	var bookings []string //the main different that you havn't to add the length of the array
	//////////////////////////////
	for {

		//conferenceName := "Go Conference" // is the same as var conferenceName = "Go Conference"
		var firstName string // here when you need to make variable and after a whole assign a value,you must give the data type
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

		//check if valid name
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isvalidticketnumber := userTickets > 0 && userTickets <= remainingTickets
		//isvalidcity:=city !="singpore"  && city !="london"

		if isValidName && isvalidticketnumber {

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
			if !isValidName {
				fmt.Println("YOUR INPUT name is invalid,please try again")
				continue
			}
			if !isvalidticketnumber {
				fmt.Println("YOUR INPUT tickets nnumber is invalid,please try again")
				continue
			}
			//fmt.Printf("we only have %v tickets\n", remainingTickets)

		}

		///////////////////////////////////////arrays & slices
	}

	//break instead of breaking the loop we will skip this iteration by adding continus

}

// you need to write the type of the variable
func greatusers(confname string) {
	fmt.Printf("welcome to %v booking application\n", confname)
}
*/
// organize everything in functions

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

/*
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
*/

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
