package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"assignment/teacher-api/controllers/classes"
	"assignment/teacher-api/controllers/students"
	"assignment/teacher-api/controllers/teachers"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	router.GET("/api/teachers", teachers.GetTeachers)
	router.POST("/api/teacher", teachers.CreateTeacher)
	router.DELETE("/api/teacher", teachers.DeleteTeacher)

	router.GET("/api/students", students.GetStudents)
	router.POST("/api/student", students.CreateStudent)
	router.DELETE("/api/student", students.DeleteStudent)

	router.POST("/api/suspend", students.SuspendStudent)
	router.POST("/api/unsuspend", students.UnsuspendStudent)

	// TODO: implement class apis (check path for second one)
	router.POST("/api/register", classes.RegisterStudents)
	// router.GET("/api/commonstudents/:teacher", students.GetStudents)

	// TODO: implement notification api
	// router.POST("/api/retrievefornotifications", students.CreateStudent)

	log.Fatal(http.ListenAndServe(":8080", router))
}
