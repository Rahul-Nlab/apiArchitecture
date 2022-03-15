package user

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) User {
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

	rows, err := h.db.Queryx(str, intId)
	if err != nil {
		return nil, "There was a problem while executing the query"
	}

	defer rows.Close()

	var UserStruct []Users
	for rows.Next() {

		var tempUserStruct Users

		err := rows.StructScan(&tempUserStruct)
		if err != nil {
			return nil, "Error while scanning database."
		}
		UserStruct = append(UserStruct, tempUserStruct)

	}

	if err := rows.Err(); err != nil {
		return nil, err.Error()
	}

	if UserStruct == nil {
		err := fmt.Sprintf("User id %v does not exist!", intId)
		return nil, err
	}

	return UserStruct, ""
}

// CreateUsers still has some error, I don't know why pointer o the user table is passed here,
// a user struct which is de-pointered has to be send here
func (h User) CreateUsers(id string, reqBody Users) string {
	var str string

	if id != "" {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return "Index id must be a valid integer!"
		}

		str = "INSERT INTO Users(first_name, middle_name, last_name, u_id) VALUES ($1, $2, $3, $4)"
		_, e := h.db.Exec(str, reqBody.First_name, reqBody.Middle_name, reqBody.Last_name, intId)

		if e != nil {
			return fmt.Sprintf("User id %v does already exists!", intId)
		}

	} else {

		if reqBody.First_name == "" || reqBody.Middle_name == "" || reqBody.Last_name == "" {
			return "Make sure you enter full name!"
		}
		str = "INSERT INTO Users(first_name, middle_name, last_name) VALUES ($1, $2, $3)"
		_, e := h.db.Exec(str, reqBody.First_name, reqBody.Middle_name, reqBody.Last_name)

		if e != nil {
			return "Problem while executing the database query"
		}

	}
	return ""
}

// DeleteUser need more optimized flow
func (h User) DeleteUser(id string) string {

	if id == "" {
		return "Please enter a User Id for deleting an entry."
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return "Index id must be a valid integer!"
	}

	rows, err := h.db.Query("SELECT * FROM Users_Addresses WHERE u_id = $1", intId)
	if err != nil {
		return err.Error()
	}

	defer rows.Close()

	if !rows.Next() {

		rows1, err1 := h.db.Query("SELECT * FROM Users WHERE u_id = $1", intId)
		if err1 != nil {
			return err.Error()
		}

		if !rows1.Next() {
			return "No such user exists."
		}

		h.db.Exec("DELETE FROM Users WHERE u_id = $1;", intId)
		return ""

	} else {
		if err := rows.Err(); err != nil {
			return err.Error()
		}
		return "User has addresses, so cannot be deleted."
	}
}

//ChangeUser needs to be executed with exec function
func (h User) ChangeUser(id string, reqBody Users) string {

	var intId int
	var str string

	if id == "" {
		return "Please enter a User Id for updating an entry."
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return "Index id must be a valid integer!"
	}

	tempStr := "SELECT * from Users WHERE u_id = $1"
	rows, err := h.db.Query(tempStr, intId)

	if err != nil {
		return "Problem while searching the data."
	}

	defer rows.Close()

	if !rows.Next() {
		return fmt.Sprintf("User id %v does already exists!", intId)
	}

	if err := rows.Err(); err != nil {
		return err.Error()
	}

	str = "UPDATE Users SET first_name = $1, middle_name = $2, last_name = $3 WHERE u_id = $4"
	_, e := h.db.Exec(str, reqBody.First_name, reqBody.Middle_name, reqBody.Last_name, intId)

	if e != nil {
		return e.Error()
	}

	return ""
}
