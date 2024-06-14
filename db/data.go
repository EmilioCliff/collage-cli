package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDatabas() error {
	var err error
	DB, err = sql.Open("sqlite3", "./sqlite-collage.db")
	if err != nil {
		return err
	}

	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return err
	}

	return DB.Ping()
}

func CreateTables() {
	query := `
		CREATE TABLE IF NOT EXISTS students (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"f_name" TEXT NOT NULL,
			"l_name" TEXT NOT NULL,
			"dob" TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS courses (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"course_name" TEXT NOT NULL,
			"total_credit" INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS enrollment (
			"course_id" INTEGER,
			"student_id" INTEGER,
			"semester" TEXT,
			"score" FLOAT,
			FOREIGN KEY (course_id) REFERENCES courses(id),
			FOREIGN KEY (student_id) REFERENCES students(id)
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tables" + err.Error())
	}
}
