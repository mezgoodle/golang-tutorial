package main

import ("fmt" 
		"net/http")

type User struct {
	name string
	age uint16
	money int16
	avgGrades, happiness float64
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Go is a super lang!")
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
	handleRequest()
}
