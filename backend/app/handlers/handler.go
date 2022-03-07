package handlers

import (
	"apiArchitecture/business/services/address"
	"apiArchitecture/business/services/user"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
)

func middleware1(handler echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Middleware 1")
	return func(c echo.Context) error {
		fmt.Println("Inside m1, accessed id = " + c.Param("id"))
		return handler(c)
	}
}

func middleware2(handler echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Middleware 2")
	return func(c echo.Context) error {
		fmt.Println("Inside m2, accessed id = " + c.Param("id"))
		return handler(c)
	}
}

func nextIdMiddleware(handler echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Next Id Middleware")
	return func(c echo.Context) error {

		a, _ := strconv.Atoi(c.Param("id"))
		c.SetParamValues(fmt.Sprint(a+1))

		fmt.Println("Inside m3, modified id = ", c.Param("id") )
		return handler(c)
	}
}

// func accessNameMiddleware (handler echo.HandlerFunc) echo.HandlerFunc {
// 	fmt.Println("Access name Middleware")
// 	return func(c echo.Context) error {

// 		reqBody := user.Users{}
// 		// var userStruct user.Users

// 		err := c.Bind(&reqBody)
// 		if err != nil {
// 			// fmt.Println("problem while binding your input ")
// 			fmt.Println(err.Error())
// 		}

// 		fmt.Println(reqBody.First_name, reqBody.Middle_name, reqBody.Last_name)

// 		// fmt.Println("Inside m3, modified id = ", c.Param("id") )
// 		return handler(c)
// 	}
// }

func Api(db *sqlx.DB, router *echo.Echo) {
	//here, the function accepts the database pointer that we've sent, initializes the userHandler Struct by calling
	// that New() function from service(user) package
	hUser := userHandlers{
		user: user.New(db),
	}

	// routes defining here for users

	router.GET("/v1/users/", hUser.GetUsersRequest)
	router.GET("/v1/users/:id/", hUser.GetUsersRequest, middleware1, middleware2, nextIdMiddleware)

	router.POST("/v1/users/", hUser.CreateUserRequest)
	router.POST("/v1/users/:id/", hUser.CreateUserRequest) //, accessNameMiddleware

	router.PUT("/v1/users/:id/", hUser.ChangeUserRequest)

	router.DELETE("/v1/users/:id/", hUser.DeleteUserRequest)

	// new addressHandlres is created here for better organization

	hAddress := addressHandlers{
		address: address.New(db),
	}

	router.GET("/v1/users/:id/addresses/", hAddress.GetAddressesRequest)
}
