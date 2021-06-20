package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proeftuin/todd-goweb/10-mongo-db/06-hands-on-2/models"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

// UserController holds a session
type UserController struct {
	session map[string]models.User
}

// NewUserController is a factory
func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

// GetUser is a GET handler
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser is a POST handler
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	id, err := uuid.NewV4()
	u.ID = id.String()

	// store the user
	uc.session[u.ID] = u
	models.StoreUsers(uc.session)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser is a DELETE handler
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	models.StoreUsers(uc.session)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user ", id, "\n")
}
