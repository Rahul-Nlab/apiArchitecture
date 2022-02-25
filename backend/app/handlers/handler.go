package handlers

import (
	"apiArchitecture/business/services/user"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Api(db *sql.DB, router *echo.Echo) {
	//here, the function accepts the database pointer that we've sent, initializes the userHandler Struct by calling
	// that New() function from service(user) package
	h := userHandlers{
		user: user.New(db),
	}

	// router.HandleFunc("/v1/users/", h.GetUsersRequest).Methods(http.MethodGet)
	router.GET("/v1/users/", h.GetUsersRequest)
	router.GET("/v1/users/:id/", h.GetUsersRequest)

}
