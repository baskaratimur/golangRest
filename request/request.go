package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "project_baskara/config"
	response "project_baskara/data"
	"project_baskara/helper"
	"project_baskara/model"
)

type ResponseData struct {
	Result string `json:"Result"`
}

func RegisFunc(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	var user model.AuthorStruct
	err := json.NewDecoder(r.Body).Decode(&user)
	helper.Panicerr(err)

	cekUsername := db.QueryRow("SELECT username, password FROM author WHERE username = ?", user.Username).Scan(&user.Username, &user.Password)
	// artinya berhasil, keluarannya kosong
	if cekUsername == nil {
		// Username already exists, return a response indicating the conflict
		response := response.ResponseResult{
			Code:   http.StatusConflict, // 409 Conflict status code
			Status: "Username already exists",
			Data:   "",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write(jsonResponse)
		return
	}

	script := "insert into author (username, password) values (?,?)"
	test, query := db.Exec(script, user.Username, user.Password)
	fmt.Println(test)
	if query != nil {
		http.Error(w, query.Error(), http.StatusInternalServerError)
	}

	response := response.ResponseResult{
		Code:   http.StatusOK,
		Status: "ok",
		Data: ResponseData{
			Result: "berhasil",
		},
	}
	fmt.Println(response)
	jsonResponse, err := json.Marshal(response)
	fmt.Println(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func Findall(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	script := "SELECT username, password from author"
	result, err := db.Query(script)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var users []model.AuthorStruct
	for result.Next() {
		var user model.AuthorStruct
		err := result.Scan(&user.Username, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)

	}

	response := response.ResponseResult{
		Code:   200,
		Status: "ok",
		Data:   users,
	}

	jsonResponse, err := json.Marshal(response)
	helper.Panicerr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func Findid(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	var username = r.URL.Query().Get("username")

	var author model.AuthorStruct
	script := "SELECT username, password from author where username = ?"
	query := db.QueryRow(script, username).Scan(&author.Username, &author.Password)
	if query != nil {
		response := response.ResponseResult{
			Code:   404,
			Status: "Data tidak ditemukan",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	response := response.ResponseResult{
		Code:   200,
		Status: "ok",
		Data:   author,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func Update(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	var user model.AuthorStruct
	err := json.NewDecoder(r.Body).Decode(&user)
	helper.Panicerr(err)

	queryParams := r.URL.Query()
	username := queryParams.Get("username")
	script := "UPDATE author SET password =? WHERE username=?"
	_, dberr := db.Exec(script, user.Password, username)
	if dberr != nil {
		http.Error(w, dberr.Error(), http.StatusInternalServerError)
	}

	response := response.ResponseResult{
		Code:   200,
		Status: "ok",
		Data:   "data sudah diupdate",
	}

	jsonResponse, err := json.Marshal(response)
	helper.Panicerr(err)
	w.Write(jsonResponse)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	defer db.Close()

	queryParams := r.URL.Query()
	username := queryParams.Get("username")

	script := "DELETE FROM author WHERE username =?"
	_, dberr := db.Exec(script, username)
	if dberr != nil {
		http.Error(w, dberr.Error(), http.StatusInternalServerError)
	}

	response := response.ResponseResult{
		Code:   200,
		Status: "ok",
		Data:   "data dihapus",
	}
	jsonResponse, err := json.Marshal(response)
	helper.Panicerr(err)
	w.Write(jsonResponse)

}
