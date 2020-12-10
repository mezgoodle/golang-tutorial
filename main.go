package main

import ("fmt"; "net/http"; "html/template")

// User model
type User struct {
	Name string
	Age uint16
	Money int16
	AvgGrades, Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and" +
	 "he has %d dollars", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(name string) {
	u.Name = name
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	bob := User{"Bob", 24, -50, 4.2, 0.8, []string{"Football", "Dancing"}}
	bob.setNewName("Max")
	// fmt.Fprintf(w, bob.getAllInfo())
	template, _ := template.ParseFiles("templates/home_page.html")
	template.Execute(w, bob)
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
