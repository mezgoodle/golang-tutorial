package main

import ("fmt"; "net/http"; "html/template"; "database/sql"; "github.com/go-sql-driver/mysql")

// Article struct
type Article struct {
	ID uint16
	Title, Anons, FullText string
}

var posts  = []Article

func index(w http.ResponseWriter, r *http.Request)  {
	template, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "login:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * from `articles`")
	if err != nil {
		panic(err)
	}
	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.ID, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	template.ExecuteTemplate(w, "index", posts)
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

	if title == "" || anons == "" || fullText == "" {
		fmt.Fprintf(w, "Enter every data")
	} else {
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
