package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"assignment/teacher-api/controllers"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/teachers", controllers.GetTeachers)
	router.POST("/teacher/:email", controllers.CreateTeacher)
	router.DELETE("/teacher/:email", controllers.DeleteTeacher)

	router.GET("/students", Index)
	router.POST("/student/:email", Index)
	router.DELETE("/student/:email", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
