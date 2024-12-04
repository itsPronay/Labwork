package main

import (
	"http"
	"log"
	"net/http"
	"routes"
)

func main() {
	http.HandleFunc("/", routes.HomePage())
	http.HandleFunc("/students/getAllStudents", routes.getAllStudents())
	http.HandleFunc("/students/getStudent", routes.getStudent())
	http.HandleFunc("/students/deleteStudent{id}", func() {
		// decrypt id and update
	})
	http.HandleFunc("/students/updateStudent{id}{studentData}", func() {
		// decrypt and pass
	})

	http.HandleFunc("/students/insertStudent{student}", func() {
		// do somtething
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}

}
