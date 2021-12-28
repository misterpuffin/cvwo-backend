package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"server/packages/middleware"
	

)

type Handler struct {
	DB *gorm.DB
}

func Router(h *Handler) *mux.Router {
	router := mux.NewRouter()

	userRoute := router.PathPrefix("/api/user").Subrouter()
	userRoute.HandleFunc("/login", h.Login).Methods("POST", "OPTIONS")
	userRoute.HandleFunc("/register", h.Register).Methods("POST", "OPTIONS")


	taskRoute := router.PathPrefix("/api/task").Subrouter()
	taskRoute.HandleFunc("/", h.CreateTask).Methods("POST", "OPTIONS")
	taskRoute.HandleFunc("/", h.GetTasks).Methods("GET", "OPTIONS")
	taskRoute.HandleFunc("/{id}", h.UpdateTask).Methods("PUT", "OPTIONS")
	taskRoute.HandleFunc("/{id}", h.DeleteTask).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	taskRoute.Use(middleware.JWTAuthentication)

	router.Use(mux.CORSMethodMiddleware(router))

	return router	
}