package main

import ("fmt"; "net/http"; "html/template")

func index(w http.ResponseWriter, r *http.Request)  {
	template, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	template.ExecuteTemplate(w, "index", nil)
}

func handleRequest()  {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func main()  {
	handleRequest()
}
