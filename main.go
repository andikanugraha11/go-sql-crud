package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID			int
	Name 		string
	Email 		string
	Age 		int
	Division 	string
}

func main()  {
	dsn := "root:hacktiv@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err :=  sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Println("Database connected")

	//createEmployee(db)
	//getEmployees(db)
	//getEmployee(db)
	//updateEmployee(db)
	//deleteEmployee(db)
	countEmployee(db)
}

func createEmployee(db *sql.DB) {
	sql := `INSERT INTO employees (name, age, email, division) VALUES (?,?,?,?)`

	res, err := db.Exec(sql, "Aji", 17, "aji@gmail.com", "IT")
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	log.Println("data inserted with id:", id)
}

func getEmployees(db *sql.DB) {
	var employees []Employee

	sql := `SELECT id, name, email, division, age FROM employees`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tempEmployee Employee
		err := rows.Scan(&tempEmployee.ID, &tempEmployee.Name, &tempEmployee.Email, &tempEmployee.Division, &tempEmployee.Age)
		if err != nil {
			panic(err)
		}

		employees = append(employees, tempEmployee)
	}

	log.Println(employees)
}


func getEmployee(db *sql.DB) {
	var employee Employee

	sql := `SELECT id, name, email, division, age FROM employees WHERE id= ?`

	id:= 2
	err := db.QueryRow(sql, id).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Division, &employee.Age)
	if err != nil {
		panic(err)
	}
	log.Println(employee)
}

func updateEmployee(db *sql.DB) {
	sql := `UPDATE employees SET name=? WHERE id=?`

	res, err := db.Exec(sql, "Andika", 2)
	if err != nil {
		panic(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	log.Println("update success:", affected)
}


func deleteEmployee(db *sql.DB) {
	sql := `DELETE FROM employees WHERE id=?`

	res, err := db.Exec(sql,  1)
	if err != nil {
		panic(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	log.Println("delete success:", affected)
}

func countEmployee(db *sql.DB)  {

	var total int

	sql := `SELECT count(*) from employees`

	err := db.QueryRow(sql).Scan(&total)
	if err != nil {
		panic(err)
	}
	log.Println(total)
}

