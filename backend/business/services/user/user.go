package user

import (
	"database/sql"
	"fmt"
	"strconv"
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{
		db: db,
	}
}

func (h User) GetUsers(id string) ([]Users, string) {
	var intId int
	var str string

	if id != "" {
		var err error
		intId, err = strconv.Atoi(id)
		if err != nil {
			return nil, "Index id must be a valid integer!"
		}
		str = "SELECT * FROM Users WHERE u_id = $1 ORDER BY u_id ASC"
	} else {
		str = "SELECT * FROM Users WHERE $1 = $1 ORDER BY u_id ASC"
	}

	rows, err := h.db.Query(str, intId)
	if err != nil {
		return nil, "There was a problem while executing the query"
	}

	var UserStruct []Users
	for rows.Next() {
		var userId int
		var fName, mName, lName string

		err := rows.Scan(&userId, &fName, &mName, &lName)
		if err != nil {
			return nil, "Error while scanning database."
		}
		UserStruct = append(UserStruct, Users{U_id: userId, First_name: fName, Middle_name: mName, Last_name: lName})
	}

	if UserStruct == nil {
		err := fmt.Sprintf("User id %v does not exit!", intId)
		return nil, err
	}

	return UserStruct, ""
}
