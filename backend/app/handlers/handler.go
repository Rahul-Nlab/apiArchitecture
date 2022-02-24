package handlers

import (
	"apiArchitecture/business/services/user"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func Api(db *sql.DB, router *mux.Router) {
	//here, the function accepts the database pointer that we've sent, initializes the userHandler Struct by calling
	// that New() function from service(user) package
	h := userHandlers{
		user: user.New(db),
	}

	router.HandleFunc("/v1/users/", h.GetUsersRequest).Methods(http.MethodGet)

}
