package api

import (
	model "back_togther/structure"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := GetAllUser()
	res := model.Response{
		Code:    200,
		Message: "sucess",
		Data:    data,
	}

	json.NewEncoder(w).Encode(res)
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&userData)
	login := fmt.Sprintf("%s", userData["login"])

	errorMsg := "user doesn't exist"
	_code := model.StatusNoData

	userExist, _data := GetUser(login)
	if userExist {
		errorMsg = "success"
		_code = model.StatusOK
	}

	resp := model.Response{
		Code:    _code,
		Message: errorMsg,
		Data:    _data,
	}
	json.NewEncoder(w).Encode(resp)
}

func GetForumData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&userData)
	id := fmt.Sprintf("%s", userData["id"])
	errorMsg := "Forum doesn't exist"
	_code := model.StatusNoData

	userExist, _data := GetForm(id)
	if userExist {
		errorMsg = "success"
		_code = model.StatusOK
	}

	resp := model.Response{
		Code:    _code,
		Message: errorMsg,
		Data:    _data,
	}
	json.NewEncoder(w).Encode(resp)
}

func GetMessageForumData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&userData)
	id := fmt.Sprintf("%s", userData["id"])
	errorMsg := "Message Forum doesn't exist"
	_code := model.StatusNoData

	userExist, _data := GetMessageForum(id)
	if userExist {
		errorMsg = "success"
		_code = model.StatusOK
	}

	resp := model.Response{
		Code:    _code,
		Message: errorMsg,
		Data:    _data,
	}
	json.NewEncoder(w).Encode(resp)
}

func GetMessageData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&userData)
	id := fmt.Sprintf("%s", userData["id"])
	errorMsg := "Message doesn't exist"
	_code := model.StatusNoData
	userExist, _data := GetMessage(id)
	if userExist {
		errorMsg = "success"
		_code = model.StatusOK
	}
	resp := model.Response{
		Code:    _code,
		Message: errorMsg,
		Data:    _data,
	}
	json.NewEncoder(w).Encode(resp)
}

func GetAdminData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&userData)
	id := fmt.Sprintf("%s", userData["id"])
	errorMsg := "Message doesn't exist"
	_code := model.StatusNoData
	userExist, _data := GetAdmin(id)
	if userExist {
		errorMsg = "success"
		_code = model.StatusOK
	}
	resp := model.Response{
		Code:    _code,
		Message: errorMsg,
		Data:    _data,
	}
	json.NewEncoder(w).Encode(resp)
}
