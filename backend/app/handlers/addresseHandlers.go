package handlers

import (
	"apiArchitecture/business/services/address"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type addressHandlers struct {
	address address.Address
}

func (h addressHandlers) GetAddressesRequest(requestContext echo.Context) error {
	fmt.Println("User Addresses GET served")

	id := requestContext.Param("id")

	addressStruct, msg := h.address.AddressOfUser(id)
	var jsonResponse address.JsonResponse
	jsonResponse.Message = msg
	
	if msg != "" {
		jsonResponse.Data = nil
		return requestContext.JSON(http.StatusNotFound, jsonResponse)
	}

	jsonResponse.Data = addressStruct
	return requestContext.JSON(http.StatusAccepted, jsonResponse)
}
