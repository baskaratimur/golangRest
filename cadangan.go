// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// )

// // secret key dari JWT disimpan disini
// var secretKey = []byte("your-secret-key")

// type User struct {
// 	Username string
// 	Password string
// }

// type LoginResponse struct {
// 	Code   int    `json:"Code"`
// 	Status string `json:"Status"`
// 	Data   struct {
// 		Token string `json:"Token"`
// 	} `json:"Data"`
// }

// type idResponse struct {
// 	Code   int         `json:"Code"`
// 	Status string      `json:"Status"`
// 	Data   interface{} `json:"data,omitempty"`
// }

// // sqldb tuh kaya koneksi ke database buat melakukan eksekusi
// func GetConec() *sql.DB {
// 	db, dbErr := sql.Open("mysql", "root:@tcp(localhost:3306)/golang-database")
// 	if dbErr != nil {
// 		log.Fatal(dbErr)
// 	}
// 	return db
// }

// func () {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/register", RegisterHandler).Methods("POST")
// 	r.HandleFunc("/login", LoginHandler).Methods("POST")
// 	r.HandleFunc("/homepage", HomePageHandler).Methods("GET")
// 	r.HandleFunc("/findId", FindByIdHandler).Methods("GET")

// 	fmt.Println("Server started on :8080")
// 	http.Handle("/", r)
// 	http.ListenAndServe(":8080", nil)
// }

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	db := GetConec()
// 	defer db.Close()
// 	var user User
// 	// request body dalam json di decode kedalam data go
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// prepare digunakan kalau u mau ngehit berkali2, lebih irit daya, kalau exec cukup 1x hit
// 	// Simpan data pengguna ke database MySQL
// 	// stmt, err := db.Prepare("INSERT INTO author (username, password) VALUES (?, ?)")
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	// defer stmt.Close()

// 	script := "insert into author (username, password) values (?,?)"
// 	_, err := db.Exec(script, user.Username, user.Password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	db := GetConec()
// 	defer db.Close()
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	// periksa password user
// 	if user.Username == "" || user.Password == "" {
// 		http.Error(w, "Username and password are required", http.StatusBadRequest)
// 		return
// 	}

// 	// Periksa kredensial pengguna dalam database MySQL
// 	var existingUser User
// 	err := db.QueryRow("SELECT username, password FROM author WHERE username = ?", user.Username).Scan(&existingUser.Username, &existingUser.Password)
// 	if err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["username"] = existingUser.Username
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	response := LoginResponse{
// 		Code:   http.StatusOK,
// 		Status: "ok",
// 		Data: struct {
// 			Token string `json:"Token"`
// 		}{
// 			Token: tokenString,
// 		},
// 	}

// 	// marshal digunakan buat decode ke format jsongo
// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonResponse)

// }

// func HomePageHandler(w http.ResponseWriter, r *http.Request) {
// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})
// 	if err != nil || !token.Valid {
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	w.Write([]byte("Welcome to the homepage!"))
// }

// func FindByIdHandler(w http.ResponseWriter, r *http.Request) {
// 	db := GetConec()
// 	defer db.Close()
// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})
// 	if err != nil || !token.Valid {
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Mendapatkan username dari token
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
// 		return
// 	}
// 	username, exists := claims["username"].(string)
// 	if !exists || username == "" {
// 		http.Error(w, "Invalid username in token", http.StatusUnauthorized)
// 		return
// 	}
// 	// fmt.Println(username)
// 	// Lakukan pencarian berdasarkan username di sini
// 	var datanya User
// 	// Ganti ika pencarian berdasarkan username sesuai dengan kebutuhan Anda
// 	// Misalnya, Anda dapat menjalankan query SQL dengan username
// 	errr := db.QueryRow("SELECT username, password FROM author WHERE username = ?", username).Scan(&datanya.Username, &datanya.Password)
// 	fmt.Println(username)

// 	if errr != nil {
// 		// Data tidak ditemukan
// 		response := idResponse{
// 			Code:   http.StatusOK,
// 			Status: "ok",
// 			Data:   "Data Tidak ditemukan",
// 		}
// 		jsonResponse, err := json.Marshal(response)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(jsonResponse)
// 		return
// 	}

// 	// Data ditemukan, maka Anda dapat mengakses "datanya" di sini
// 	response := idResponse{
// 		Code:   http.StatusOK,
// 		Status: "ok",
// 		Data:   datanya,
// 	}
// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonResponse)
// }
