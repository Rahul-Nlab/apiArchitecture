package handlers

import (
	"apiArchitecture/business/services/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandlers struct {
	user user.User
}

// GetUsersRequest handles the getUser request and returns the JSON response
// It won't be processing the request btw, GetUsers will do it from business package
func (h userHandlers) GetUsersRequest(requestContext echo.Context) error {

	fmt.Println("User GET served")

	id := requestContext.Param("id")

	userStruct, msg := h.user.GetUsers(id)
	var jsonResponse user.JsonResponse
	jsonResponse.Message = msg
	
	if msg != "" {
		jsonResponse.Data = nil
		return requestContext.JSON(http.StatusNotFound, jsonResponse)
	}

	jsonResponse.Data = userStruct

	return requestContext.JSON(http.StatusOK, jsonResponse)
}

func (h userHandlers) DeleteUserRequest(requestContext echo.Context) error {

	fmt.Println("User DELETE served")

	id := requestContext.Param("id")

	msg := h.user.DeleteUser(id)

	var jsonResponse user.JsonResponse

	if msg != "" {
		jsonResponse.Data = nil
		jsonResponse.Message = msg
		return requestContext.JSON(http.StatusBadRequest, jsonResponse)
	}

	jsonResponse.Message = "Deleted Successfully"

	return requestContext.JSON(http.StatusOK, jsonResponse)
}

func (h userHandlers) CreateUserRequest(requestContext echo.Context) error {

	fmt.Println("User POST served")

	id := requestContext.Param("id")

	reqBody := user.Users{}
	var jsonResponse user.JsonResponse

	e := requestContext.Bind(&reqBody)
	if e != nil {
		fmt.Println("problem while binding your input")
	}

	msg := h.user.CreateUsers(id, reqBody)
	jsonResponse.Message = msg
	jsonResponse.Data = nil

	if msg != "" {
		return requestContext.JSON(http.StatusBadRequest, jsonResponse)
	}

	jsonResponse.Message = "Successfully added!"

	return requestContext.JSON(http.StatusOK, jsonResponse)
}

func (h userHandlers) ChangeUserRequest(requestContext echo.Context) error {

	fmt.Println("User PUT served")

	id := requestContext.Param("id")

	reqBody := user.Users{}
	var jsonResponse user.JsonResponse

	err := requestContext.Bind(&reqBody)
	if err != nil {
		fmt.Println("problem while binding your input ")
	}

	msg := h.user.ChangeUser(id, reqBody)
	jsonResponse.Message = msg
	jsonResponse.Data = nil

	if msg != "" {
		return requestContext.JSON(http.StatusBadRequest, jsonResponse)
	}

	return requestContext.JSON(http.StatusOK, jsonResponse)
}
