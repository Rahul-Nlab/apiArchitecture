package address

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Address struct {
	db *sql.DB
}

func New(db *sql.DB) Address {
	return Address{
		db: db,
	}
}

func (h Address) AddressOfUser(id string) ([]Addresses, string) {

	if id == "" {
		return nil, "Please enter a User Id for fetching addresses"
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Sprintf("'%v' is not a valid integer!", id)
	}

	rows, err := h.db.Query("SELECT * FROM Addresses where a_id in (select a_id from users_addresses where u_id = $1); ", intId)

	if err != nil {
		return nil, err.Error()
	}

	defer rows.Close()

	userAddresses := make([]Addresses, 0)

	for rows.Next() {
		var addressPincode, addressId int
		var addressStreet, addressArea, addressCity string

		err := rows.Scan(&addressId, &addressStreet, &addressArea, &addressPincode, &addressCity)
		if err != nil {
			return nil, err.Error()
		}

		userAddresses = append(userAddresses, Addresses{A_id: addressId, Street: addressStreet, Area: addressArea, Pincode: addressPincode, City: addressCity})
	}

	if len(userAddresses) == 0 {
		return nil, fmt.Sprintf("User %v has no addresses (or user %v does not exists).", id, id)
	}

	return userAddresses, ""

}
