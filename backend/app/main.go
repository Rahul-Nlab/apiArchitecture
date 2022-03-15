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

	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	ErrorLogger.Println("Something went wrong")

	db := database.Setup()
	defer db.Close()

	router := echo.New()

	//API FUNCTION TO BE CREATED
	handlers.Api(db, router)

	fmt.Println("Up and runnin..!")
	// log.Fatalln("Good log", log.LstdFlags, log.Ldate, log.Lmicroseconds, log.Llongfile, log.Lshortfile, log.LUTC)

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	router.Logger.Fatal(router.Start(":5435"))
}
