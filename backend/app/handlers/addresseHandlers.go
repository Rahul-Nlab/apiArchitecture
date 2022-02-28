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

	addressStruct, err := h.address.AddressOfUser(id)

	if err != "" {
		return requestContext.JSON(http.StatusNotFound, err)
	}

	return requestContext.JSON(http.StatusAccepted, addressStruct)
}
