package repository

import (
	"reporting-utility/internal/db"
	"reporting-utility/internal/models"
	"strconv"
)

func FetchAllStudents() ([]models.Student, error) {
	rows, err := db.DB.Query("SELECT * FROM students")
	if err != nil { return nil, err }
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		var id int
		rows.Scan(&id, &s.StudentName, &s.Address, &s.State, &s.District, &s.Taluka, &s.Gender, &s.DOB, &s.Photo, &s.Handicapped, &s.Email, &s.MobileNumber, &s.BloodGroup)
		s.ID = strconv.Itoa(id)
		students = append(students, s)
	}
	return students, nil
}