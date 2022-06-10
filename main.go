package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mysql connected...")
	dbm = db
}

func createTable() {
	query := `Create table student (
              sid int auto_increment,
			  name text not null,
			  course text not null,
			  gender text not null,
			  dob    text not null,
			  created_at datetime,
			  primary key(sid)
	        );`
	_, err := dbm.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" student table created...\n")
}

func insertTable() {
	sid := 5
	name := "Bhashkar kumar"
	course := "B.tech"
	gender := "Male"
	dob := "11-02-1969"
	created_at := time.Now()
	result, err := dbm.Exec(`insert into student (sid,name,course,gender,dob,created_at) values (?,?,?,?,?,?);`,
		sid, name, course, gender, dob, created_at)
	if err != nil {
		log.Fatal(err)
	} else {
		value, _ := result.LastInsertId()
		fmt.Println("Record inserted for student Id=", value)
	}
}

func updateTable() {
	sid := 2
	name := "Iffco Raja"
	_, err := dbm.Exec(`update student set name=? where sid=?`,
		name, sid)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Record updated...")
	}
}

func getStudentData() {
	type student struct {
		sid        int
		name       string
		course     string
		gender     string
		dob        string
		created_at time.Time
	}
	var s student
	rows, err := dbm.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			rows.Scan(&s.sid, &s.name, &s.course, &s.gender, &s.dob, &s.created_at)
			fmt.Println("Student Id=", s.sid)
			fmt.Println("Student Name=", s.name)
			fmt.Println("Student Course=", s.course)
			fmt.Println("Student Gender=", s.gender)
			fmt.Println("Student Dob=", s.dob)
			fmt.Println("Student Id create at=", s.created_at)
		}
	}
}

func deleteTable() {
	_, err := dbm.Exec(`delete from student where sid=?`, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record delete for sid", 5)
}

func main() {
	ConnectDB()
	createTable()
	insertTable()
	updateTable()
	deleteTable()
	getStudentData()

}
