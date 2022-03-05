package address

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Address struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Address {
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

	rows, err := h.db.Queryx("SELECT * FROM Addresses where a_id in (select a_id from users_addresses where u_id = $1); ", intId)

	if err != nil {
		return nil, err.Error()
	}

	defer rows.Close()

	userAddresses := make([]Addresses, 0)

	for rows.Next() {
		var tempUserAddress Addresses
		err := rows.StructScan(&tempUserAddress)
		if err != nil {
			return nil, "Error while scanning database."
		}
		userAddresses = append(userAddresses, tempUserAddress)
	}

	if err := rows.Err(); err != nil {
		return nil, err.Error()
	}

	if len(userAddresses) == 0 {
		return nil, fmt.Sprintf("User %v has no addresses (or user %v does not exists).", id, id)
	}

	return userAddresses, ""

}
