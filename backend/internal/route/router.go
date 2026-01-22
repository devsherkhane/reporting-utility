package route

import (
	"net/http"
	"reporting-utility/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	// Static File Server
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	
	// Student Management
	router.HandleFunc("/students", handler.GetAllStudents).Methods("GET")
	router.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	router.HandleFunc("/students", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students", handler.DeleteStudent).Methods("DELETE")

	// Reporting
	router.HandleFunc("/students/pdf", handler.DownloadStudentsPDF).Methods("GET")

	router.HandleFunc("/", handler.Home).Methods("GET")

	router.HandleFunc("/students/profiles", handler.DownloadStudentProfilesPDF).Methods("GET")

}
