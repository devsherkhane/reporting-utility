package handler

import (
	"net/http"
	"reporting-utility/internal/report"
	"reporting-utility/internal/repository"
)

func DownloadStudentsPDF(w http.ResponseWriter, r *http.Request) {
	students, err := repository.FetchAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pdf, err := report.GenerateStudentsPDF(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=students.pdf")
	pdf.Write(w)
}
// backend/internal/handler/report_handler.go

func DownloadStudentProfilesPDF(w http.ResponseWriter, r *http.Request) {
	students, err := repository.FetchAllStudents()
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	pdf, err := report.GenerateStudentProfilesPDF(students)
	if err != nil {
		http.Error(w, "PDF generation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=student_profiles.pdf")
	
	// Use WriteTo to stream the PDF buffer to the response writer
	_, err = pdf.WriteTo(w)
	if err != nil {
		http.Error(w, "Error streaming PDF: "+err.Error(), http.StatusInternalServerError)
	}
}
