package api

import (
	"server/packages/db"
	"server/packages/utils"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	task.UserID = r.Header.Get("userID")
	task.Name = r.FormValue("name")
	task.Tag = r.FormValue("tag")

	err := h.DB.Create(&task).Error
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)





}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	reqUserID := r.Header.Get("userID")
	var tasks []db.Task
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	err := h.DB.Where("user_id = ?", reqUserID).Find(&tasks).Error
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tasks)
	
}



func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	reqUserID := r.Header.Get("userID")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	var task db.Task

	if err := h.DB.First(&task, taskID).Error; err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	} else if task.UserID != reqUserID {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Unauthorised access")
		return
	} else {
		task.Name = r.FormValue("name")
		h.DB.Save(&task)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}



	
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	reqUserID := r.Header.Get("userID")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	var task db.Task

	if err := h.DB.First(&task, taskID).Error; err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	} else if task.UserID != reqUserID {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Unauthorised access")
		return
	} else {
		h.DB.Delete(&task)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
	
}