package api

import (
	"net/http"
	"encoding/json"
	"server/packages/db"
	"server/packages/utils"
)


func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var user db.User
	h.DB.Where("email = ?", r.FormValue("email")).Find(&user)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if user.CheckPasswordHash(r.FormValue("password")) {
		token, err := user.GenerateJWT()
		if err != nil {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
			return
		}
		json.NewEncoder(w).Encode(&token)
	} else {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Password incorrect")
		return
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user db.User
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	user.Password = user.HashPassword(r.FormValue("password"))

	err := h.DB.Create(&user).Error
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)

}

