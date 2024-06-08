package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
	Year     int
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("assets/index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola", Year: 1972},
				{Title: "Blade Runner", Director: "Ridley Scott", Year: 1982},
				{Title: "The Thing", Director: "John Carpenter", Year: 1982},
			},
		}

		err := tmpl.Execute(w, films)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		tmpl := template.Must(template.ParseFiles("assets/index.html"))

		err := tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// htmlStr := fmt.Sprintf("<li class='bg-blue text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	// Start http server at port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
