package handlers

import (
	"apiArchitecture/business/services/address"
	"apiArchitecture/business/services/user"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func nextIdMiddleware(handler echo.HandlerFunc) echo.HandlerFunc {
// 	fmt.Println("Next Id Middleware")
// 	return func(c echo.Context) error {

// 		a, _ := strconv.Atoi(c.Param("id"))
// 		c.SetParamValues(fmt.Sprint(a+1))

// 		fmt.Println("Inside m3, modified id = ", c.Param("id") )

// 		c.Response().Header().Set("Context-Type", "application/notjson")
// 		// c.Response().Write()

// 		return handler(c)
// 	}
// }

func middleware1(handler echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Middleware 1")
	return func(c echo.Context) error {
		fmt.Println("Inside m1, accessed id = " + c.Param("id"))
		// log.Print("Inside m1, accessed id = (" + c.Param("id") + ")")
		// fmt.Printf("c.Response: %v\n", c.Response)
		c.Response().After(func() {
			fmt.Println("Response written")
			// var tempStruct user.JsonResponse
			// tempstruct := c.Bind(&tempStruct)
			// fmt.Println(c.JSON())
			// c.Reset()
		})
		return handler(c)

		// ("Response middleware1")

	}
}

func middleware2(handler echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Middleware 2")
	return func(c echo.Context) error {
		// fmt.Println()
		log.Println("Inside m2, accessed id = " + c.Param("id"))
		return handler(c)
	}
}

func Api(db *sqlx.DB, router *echo.Echo) {
	//here, the function accepts the database pointer that we've sent, initializes the userHandler Struct by calling
	// that New() function from service(user) package
	hUser := userHandlers{
		user: user.New(db),
	}

	// router.Group("/v1/")
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${method}  ${status}  ${latency_human}` + "\n",
	}))

	log.Println("Second log in file...")

	// routes defining here for users

	router.GET("/v1/users/", hUser.GetUsersRequest, middleware2)
	router.GET("/v1/users/:id/", hUser.GetUsersRequest)

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
