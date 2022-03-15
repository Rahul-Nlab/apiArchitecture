package main

import (
	"apiArchitecture/app/handlers"
	"apiArchitecture/business/database"
	"log"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db := database.Setup()
	defer db.Close()

	router := echo.New()

	// router 

	//API FUNCTION TO BE CREATED
	handlers.Api(db, router)

	fmt.Println("Up and runnin..!")
	log.Print("First log!")

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	router.Logger.Fatal(router.Start(":5435"))
}
