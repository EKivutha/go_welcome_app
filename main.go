package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//interface
type Welcome struct {
	Name string
	Time string
}

func main() {
	//get current time from the os and the error message print
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//get front-end page to send the database data to
	templates := template.Must(template.ParseFiles("template/welcome-template.html"))

	//link our css styling to the html page
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."+"/static/"))))

	//initialize out url handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//get name from url and print out on the front-end
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		//handle url errors
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//start the server at specified port
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
