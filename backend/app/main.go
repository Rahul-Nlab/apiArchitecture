package main

import (
	"apiArchitecture/app/handlers"
	"apiArchitecture/business/database"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db := database.Setup()
	defer db.Close()

	// router := mux.NewRouter()
	router := echo.New()

	//API FUNCTION TO BE CREATED
	handlers.Api(db, router)

	fmt.Println("Up!")

	// headersOk := muxhan.AllowedHeaders([]string{"*"})
	// originsOk := muxhan.AllowedOrigins([]string{"*"})
	// methodsOk := muxhan.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	router.Logger.Fatal(router.Start(":5435"))
}
