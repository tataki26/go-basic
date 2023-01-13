package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"mongodb/06_hands-on/models"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	users map[string]map[string]string
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	u := models.User{}

	// get user from map
	u.Id = uc.users[id]["id"]
	u.Name = uc.users[id]["name"]
	u.Gender = uc.users[id]["gender"]
	u.Age, _ = strconv.Atoi(uc.users[id]["age"])

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = strconv.Itoa(rand.Intn(100))

	um := map[string]string{
		"id":     u.Id,
		"name":   u.Name,
		"gender": u.Gender,
		"age":    strconv.Itoa(u.Age),
	}

	uc.users = make(map[string]map[string]string, 1)

	uc.users[um["id"]] = um

	uj, err := json.Marshal(uc.users)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.users, id)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", id, "\n")
}
