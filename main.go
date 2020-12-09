package main

import ("fmt" 
		"net/http")

// User model
type User struct {
	name string
	age uint16
	money int16
	avgGrades, happiness float64
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	bob := User{"Bob", 24, -50, 4.2, 0.8}
	bob.name = "Max"
	fmt.Fprintf(w, "User name is: " + bob.name)
}

func contactsPage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":5000", nil)
}

func main()  {
	// var bob User = ...
	// bob := User{name: "Bob", age: 24, money: -50, avgGrades: 4.2, happiness: 0.8}

	handleRequest()
}
