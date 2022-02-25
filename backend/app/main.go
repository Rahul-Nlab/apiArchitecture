package main

import (
	"apiArchitecture/app/handlers"
	"apiArchitecture/business/database"
	
	"fmt"
	"log"
	"net/http"

	muxhan "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_"github.com/lib/pq"
)

func main() {
	db := database.Setup()
	defer db.Close()

	router := mux.NewRouter()
	
	//API FUNCTION TO BE CREATED
	handlers.Api(db, router)

	// router := echo.New()

	fmt.Println("Up!")

	headersOk := muxhan.AllowedHeaders([]string{"*"})
	originsOk := muxhan.AllowedOrigins([]string{"*"})
	methodsOk := muxhan.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":5434", muxhan.CORS(headersOk, originsOk, methodsOk)(router)))
}
