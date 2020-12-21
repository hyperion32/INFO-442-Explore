package gatewaysrc

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //needed
)

//SQLStore holds db info
type SQLStore struct {
	DB *sql.DB
}

//GetUserInfo checks if the user is in the db, and if they are, returns the user info, otherwise creates the user in the databases
func (sqls *SQLStore) GetUserInfo(guser *GoogleUser) (*User, error) {
	user := &User{}

	insq := "select * from user where google_id = ?"

	errQuery := sqls.DB.QueryRow(insq, guser.GoogleID).Scan(&user.UserID, &user.GoogleID, &user.Email, &user.FirstName, &user.LastName, &user.PhotoURL)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			insq = "insert into user(google_id, email, first_name, last_name, photo_url) values(?,?,?,?,?)"

			res, errExec := sqls.DB.Exec(insq, guser.GoogleID, guser.Email, guser.FirstName, guser.LastName, guser.PhotoURL)
			if errExec != nil {
				return nil, errExec
			}

			id, idErr := res.LastInsertId()
			if idErr != nil {
				return nil, idErr
			}
			user.UserID = int(id)
			user.GoogleID = guser.GoogleID
			user.Email = guser.Email
			user.FirstName = guser.FirstName
			user.LastName = guser.LastName
			user.PhotoURL = guser.PhotoURL

			return user, nil
		}
		return nil, errQuery
	}

	return user, nil
}
