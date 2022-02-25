package handlers

import (
	"apiArchitecture/business/services/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//userHandler struct to be made here
type userHandlers struct {
	user user.User
}

// GetUsersRequest handles the getUser request and returns the JSON response
// It won't be processing the request btw, GetUsers will do it from business package
func (h userHandlers) GetUsersRequest(requestContext echo.Context) error {

	fmt.Println("User GET served")

	// vars := mux.Vars(r)
	// id := vars["id"]

	id := requestContext.Param("id")

	userStruct, err := h.user.GetUsers(id)

	if err != "" {
		// w.Write([]byte(err))
		response := err

		return requestContext.JSON(http.StatusNotFound, response)
	}

	// json.NewEncoder(w).Encode(userStruct)
	return requestContext.JSON(http.StatusOK, userStruct)
}
