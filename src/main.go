package main

import (
	"assignment/teacher-api/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	router.POST("/api/register", controllers.RegisterStudents)
	router.GET("/api/commonstudents", controllers.GetCommonStudents)
	router.POST("/api/suspend", controllers.SuspendStudent)
	router.POST("/api/retrievefornotifications", controllers.RetrieveForNotifications)

	log.Fatal(http.ListenAndServe(":8080", router))
}
