package handlers

import (
	"apiArchitecture/business/services/address"
	"apiArchitecture/business/services/user"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Api(db *sql.DB, router *echo.Echo) {
	//here, the function accepts the database pointer that we've sent, initializes the userHandler Struct by calling
	// that New() function from service(user) package
	hUser := userHandlers{
		user: user.New(db),
	}

	router.GET("/v1/users/", hUser.GetUsersRequest)
	router.GET("/v1/users/:id/", hUser.GetUsersRequest)

	router.POST("/v1/users/", hUser.CreateUserRequest)
	router.POST("/v1/users/:id/", hUser.CreateUserRequest)

	router.PUT("/v1/users/:id/", hUser.ChangeUserRequest)

	router.DELETE("/v1/users/:id/", hUser.DeleteUserRequest)

	hAddress := addressHandlers{
		address: address.New(db),
	}

	router.GET("/v1/users/:id/addresses/", hAddress.GetAddressesRequest)

}
