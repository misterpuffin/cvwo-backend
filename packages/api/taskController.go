package api

import (
	"server/packages/db"
	"server/packages/utils"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	r.ParseForm()
	reqUserID := r.Header.Get("userID")
	task.UserID = reqUserID
	task.Name = r.FormValue("name")
	if tags := r.Form["tags"]; tags == nil {
		task.Tags = make([]string, 0)
	} else {
		task.Tags = tags
	}

	err := h.DB.Create(&task).Error
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	for _, tagName := range task.Tags {
		var tag db.Tag 
		tag.Name = tagName
		tag.UserID = reqUserID
		tag.TaskID = strconv.FormatUint(uint64(task.ID), 10)

		err := h.DB.Create(&tag).Error
		if err != nil {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
			return
		}
	}

	utils.NewJSONResponse(w, &task)
}

type GetResponse struct {
	Tasks 	[]db.Task `json:"tasks"`
	Tags	[]string `json:"tags"`
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	reqUserID := r.Header.Get("userID")
	var tasks []db.Task

	err := h.DB.Where("user_id = ?", reqUserID).Find(&tasks).Error
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	for index, task := range tasks {
		var tags []db.Tag
		err := h.DB.Select("name").Where("task_id = ?", task.ID).Find(&tags).Error
		if err != nil {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
			return
		}

		tagNames := make([]string, len(tags))
		for i, tag := range tags {
			tagNames[i] = tag.Name
		}

		tasks[index].Tags = tagNames
	}

	var tags []db.Tag

	err = h.DB.Distinct("name").Select("name").Where("user_id = ?", reqUserID).Find(&tags).Error
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	tagNames := make([]string, len(tags))
	for i, tag := range tags {
		tagNames[i] = tag.Name
	}

	utils.NewJSONResponse(w, GetResponse{tasks, tagNames})
	
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
		r.ParseForm()
		task.Name = r.FormValue("name")
		if done := r.FormValue("done"); done != "" {
			task.Done, err = strconv.ParseBool(done)
			if err != nil {
				utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
				return
			}
		}

		if tags := r.Form["tags"]; tags == nil {
			task.Tags = make([]string, 0)
		} else {
			task.Tags = tags
		}
		h.DB.Save(&task)

		h.DB.Where("task_id = ?", taskID).Delete(db.Tag{})
		for _, tagName := range task.Tags {
			var tag db.Tag 
			tag.Name = tagName
			tag.UserID = reqUserID
			tag.TaskID = strconv.FormatUint(uint64(task.ID), 10)
	
			err := h.DB.Create(&tag).Error
			if err != nil {
				utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
				return
			}
		}

		utils.NewJSONResponse(w, &task)
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
		h.DB.Where("task_id = ?", taskID).Delete(db.Tag{})
		utils.NewJSONResponse(w, &task)
	}
	
}