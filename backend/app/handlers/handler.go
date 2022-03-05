package handlers

import (
	"apiArchitecture/business/services/address"
	"apiArchitecture/business/services/user"
	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
)

func Api(db *sqlx.DB, router *echo.Echo) {
	//here, the function accepts the database pointer that we've sent, initializes the userHandler Struct by calling
	// that New() function from service(user) package
	hUser := userHandlers{
		user: user.New(db),
	}

	// routes defining here for users

	router.GET("/v1/users/", hUser.GetUsersRequest)
	router.GET("/v1/users/:id/", hUser.GetUsersRequest)

	router.POST("/v1/users/", hUser.CreateUserRequest)
	router.POST("/v1/users/:id/", hUser.CreateUserRequest)

	router.PUT("/v1/users/:id/", hUser.ChangeUserRequest)

	router.DELETE("/v1/users/:id/", hUser.DeleteUserRequest)

	// new addressHandlres is created here for better organization

	hAddress := addressHandlers{
		address: address.New(db),
	}

	router.GET("/v1/users/:id/addresses/", hAddress.GetAddressesRequest)
}
