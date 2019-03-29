package users

import (
	"api-mongo/pkg/db"
	"encoding/json"
	"log"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	u := User{}

	conn, err := db.GetDB()
	if err != nil {
		log.Printf("connect db error: %v", err)
		return
	}

	users, err := u.GetUsers(conn)
	if err != nil {
		log.Printf("error : %v", err)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		log.Printf("transform error : %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	u := User{}
	var user User

	conn, err := db.GetDB()
	if err != nil {
		log.Printf("connect db error: %v", err)
		return
	}

	if json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error: %v", err)
		return
	}

	result, err := u.CreateUser(conn, user)
	if err != nil {
		log.Printf("create user error: %v", err)
		return
	}

	userResult, err := json.Marshal(result)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	w.Write(userResult)
}


