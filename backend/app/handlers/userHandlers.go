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

	id := requestContext.Param("id")

	userStruct, err := h.user.GetUsers(id)

	if err != "" {
		return requestContext.JSON(http.StatusNotFound, err)
	}

	return requestContext.JSON(http.StatusOK, userStruct)
}

func (h userHandlers) DeleteUserRequest(requestContext echo.Context) error {

	fmt.Println("User DELETE served")

	id := requestContext.Param("id")

	response := h.user.DeleteUser(id)

	return requestContext.JSON(http.StatusOK, response)

}

func (h userHandlers) CreateUserRequest(requestContext echo.Context) error {

	fmt.Println("User POST served")

	id := requestContext.Param("id")

	// userStruct, err := h.user.GetUsers(id)

	reqBody := new(user.Users)
	if err := requestContext.Bind(reqBody); err != nil {
		response := "Please enter a valid json input!"
		return requestContext.JSON(http.StatusBadRequest, response)
	}

	response := h.user.CreateUsers(id, reqBody)

	return requestContext.JSON(http.StatusOK, response)
}

func (h userHandlers) ChangeUserRequest(requestContext echo.Context) error {

	fmt.Println("User PUT served")

	id := requestContext.Param("id")

	reqBody := new(user.Users)
	if err := requestContext.Bind(reqBody); err != nil {
		response := "Please enter a valid json input!"
		return requestContext.JSON(http.StatusBadRequest, response)
	}

	response := h.user.ChangeUser(id, reqBody)

	return requestContext.JSON(http.StatusOK, response)
}
