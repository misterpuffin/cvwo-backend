package api

import (
	"net/http"
	"server/packages/db"
	"server/packages/utils"
)


func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var user db.User
	h.DB.Where("email = ?", r.FormValue("email")).Find(&user)

	if user.CheckPasswordHash(r.FormValue("password")) {
		token, err := user.GenerateJWT()
		if err != nil {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
			return
		}
		utils.NewJSONResponse(w, &token)
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
	utils.NewJSONResponse(w, &user)

}

