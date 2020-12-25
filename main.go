package main

import ("fmt"; "net/http"; "html/template"; "database/sql"; "github.com/go-sql-driver/mysql")

func index(w http.ResponseWriter, r *http.Request)  {
	template, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	template.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request)  {
	template, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	template.ExecuteTemplate(w, "create", nil)
}

func saveArticle(w http.ResponseWriter, r *http.Request)  {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")

	db, err := sql.Open("mysql", "login:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, fullText))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func handleRequest()  {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", index)
	http.HandleFunc("/save_article", saveArticle)
	http.ListenAndServe(":5000", nil)
}

func main()  {
	handleRequest()
}
