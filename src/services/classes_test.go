package services

import (
	"assignment/teacher-api/database"
	"assignment/teacher-api/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	database.InitialiseTestDB()

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestRegisterStudentsToTeacherService(t *testing.T) {
	clearDB()
	studentSlice := []string{"ann@gmail.com", "ben@gmail.com"}
	RegisterStudentsToTeacherService(studentSlice, "t_tan@gmail.com")

	students, teachers, classes := getStudentsTeachersClasses(t)

	assert.Equal(t, 2, len(students))
	assert.Equal(t, 1, len(teachers))
	assert.Equal(t, 2, len(classes))

	studentSlice = []string{"carl@gmail.com", "ben@gmail.com"}
	RegisterStudentsToTeacherService(studentSlice, "t_teo@gmail.com")

	students, teachers, classes = getStudentsTeachersClasses(t)

	assert.Equal(t, 3, len(students))
	assert.Equal(t, 2, len(teachers))
	assert.Equal(t, 4, len(classes))

	studentSlice = []string{"ben@gmail.com", "daniel@gmail.com"}
	RegisterStudentsToTeacherService(studentSlice, "t_tan@gmail.com")

	students, teachers, classes = getStudentsTeachersClasses(t)

	assert.Equal(t, 4, len(students))
	assert.Equal(t, 2, len(teachers))
	assert.Equal(t, 5, len(classes))

	clearDB()
}

func TestGetCommonStudentsService(t *testing.T) {
	setUp()

	teachersSlice := []string{"t_lee@gmail.com"}
	students, err := GetCommonStudentsService(teachersSlice)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 3, len(students))

	teachersSlice = []string{"t_lee@gmail.com", "t_tan@gmail.com"}
	students, err = GetCommonStudentsService(teachersSlice)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 1, len(students))

	teachersSlice = []string{"t_teo@gmail.com", "t_tan@gmail.com"}
	students, err = GetCommonStudentsService(teachersSlice)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 0, len(students))

	teachersSlice = []string{"t_lee@gmail.com", "t_invalid@gmail.com"}
	students, err = GetCommonStudentsService(teachersSlice)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 0, len(students))

	clearDB()
}

func TestSuspendStudentService(t *testing.T) {
	setUp()

	result := database.DB.QueryRow("SELECT * FROM students WHERE email = ?;", "ben@gmail.com")
	student, err := models.ReadRowAsStudent(result)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, false, student.Is_Suspended)

	SuspendStudentService("ben@gmail.com", true)

	result = database.DB.QueryRow("SELECT * FROM students WHERE email = ?;", "ben@gmail.com")
	student, err = models.ReadRowAsStudent(result)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, true, student.Is_Suspended)

	SuspendStudentService("ben@gmail.com", true)

	result = database.DB.QueryRow("SELECT * FROM students WHERE email = ?;", "ben@gmail.com")
	student, err = models.ReadRowAsStudent(result)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, true, student.Is_Suspended)

	result = database.DB.QueryRow("SELECT * FROM students WHERE email = ?;", "ann@gmail.com")
	student, err = models.ReadRowAsStudent(result)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, true, student.Is_Suspended)

	SuspendStudentService("ann@gmail.com", true)

	result = database.DB.QueryRow("SELECT * FROM students WHERE email = ?;", "ann@gmail.com")
	student, err = models.ReadRowAsStudent(result)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, true, student.Is_Suspended)

	clearDB()
}

func TestRetrieveForNotificationsService(t *testing.T) {
	setUp()

	students, err := RetrieveForNotificationsService("t_lee@gmail.com", "Hello")
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 2, len(students))

	students, err = RetrieveForNotificationsService("t_lee@gmail.com", "Hello @daniel@gmail.com")
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 3, len(students))

	students, err = RetrieveForNotificationsService("t_tan@gmail.com", "Hello")
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 2, len(students))

	students, err = RetrieveForNotificationsService("t_tan@gmail.com", "Hello @ben@gmail.com")
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 3, len(students))

	students, err = RetrieveForNotificationsService("t_tan@gmail.com", "Hello @daniel@gmail.com")
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, 2, len(students))
	clearDB()
}

/* HELPER FUNCTIONS */

func clearDB() {
	database.DB.Exec("DELETE FROM students;")
	database.DB.Exec("DELETE FROM teachers;")
	database.DB.Exec("DELETE FROM classes;")
}

func getStudentsTeachersClasses(t *testing.T) ([]models.Student, []models.Teacher, []models.Class) {
	result, err := database.DB.Query("SELECT * FROM students;")
	if err != nil {
		t.FailNow()
	}
	students, err := models.ReadRowsAsStudents(result)
	if err != nil {
		t.FailNow()
	}
	result, err = database.DB.Query("SELECT * FROM teachers;")
	if err != nil {
		t.FailNow()
	}
	teachers, err := models.ReadRowsAsTeachers(result)
	if err != nil {
		t.FailNow()
	}

	result, err = database.DB.Query("SELECT * FROM classes;")
	if err != nil {
		t.FailNow()
	}
	classes, err := models.ReadRowsAsClass(result)
	if err != nil {
		t.FailNow()
	}

	return students, teachers, classes
}

func seedDB() {
	database.DB.Exec("INSERT INTO students (email, is_suspended) VALUES (?, ?);", "ann@gmail.com", true)
	database.DB.Exec("INSERT INTO students (email, is_suspended) VALUES (?, ?);", "ben@gmail.com", false)
	database.DB.Exec("INSERT INTO students (email, is_suspended) VALUES (?, ?);", "carl@gmail.com", false)
	database.DB.Exec("INSERT INTO students (email, is_suspended) VALUES (?, ?);", "daniel@gmail.com", false)
	database.DB.Exec("INSERT INTO students (email, is_suspended) VALUES (?, ?);", "eve@gmail.com", true)

	database.DB.Exec("INSERT INTO teachers (email) VALUES (?);", "t_lee@gmail.com")
	database.DB.Exec("INSERT INTO teachers (email) VALUES (?);", "t_tan@gmail.com")
	database.DB.Exec("INSERT INTO teachers (email) VALUES (?);", "t_teo@gmail.com")

	result, _ := database.DB.Query("SELECT * FROM students ORDER BY email ASC;")
	students, _ := models.ReadRowsAsStudents(result)

	result, _ = database.DB.Query("SELECT * FROM teachers ORDER BY email ASC;")
	teachers, _ := models.ReadRowsAsTeachers(result)

	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[0].ID, students[0].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[0].ID, students[1].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[0].ID, students[2].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[1].ID, students[2].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[1].ID, students[3].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[1].ID, students[4].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[2].ID, students[0].ID)
	database.DB.Exec("INSERT INTO classes (teacher_id, student_id) VALUES (?, ?);", teachers[2].ID, students[1].ID)
}

func setUp() {
	clearDB()
	seedDB()
}
