package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"apiArchitecture/business/services/user"

	"github.com/gorilla/mux"
)

//userHandler struct to be made here
type userHandlers struct {
	user user.User
}

func (h userHandlers) GetUsersRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("User GET served")
	// w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") 

	vars := mux.Vars(r)
	id := vars["id"]

	userStruct, err := h.user.GetUsers(id)

	if err != "" {
		w.Write([]byte(err))
		return
	}

	json.NewEncoder(w).Encode(userStruct)
}
