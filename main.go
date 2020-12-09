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

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and" +
	 "he has %d dollars", u.name, u.age, u.money)
}

func (u *User) setNewName(name string) {
	u.name = name
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	bob := User{"Bob", 24, -50, 4.2, 0.8}
	bob.setNewName("Max")
	fmt.Fprintf(w, bob.getAllInfo())
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
