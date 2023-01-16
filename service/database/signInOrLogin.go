package database

import (
	"WASAPhoto/service/objects"
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) SignInOrLogin(username string) (objects.Identifier, error) {
	// TODO make username unique
	var id int64
	err := db.c.QueryRow(fmt.Sprintf(`SELECT id FROM %s WHERE username=?`, USERSTABLE), username).Scan(&id)
	// username not found procede to auth the user and return new id
	// TODO check iff the username is valid
	if errors.Is(err, sql.ErrNoRows) {
		res, err := db.c.Exec(`INSERT INTO users (username) VALUES (?)`, username)
		if err != nil {
			return objects.Identifier{}, err
		}
		id, err = res.LastInsertId()
		if err != nil {
			return objects.Identifier{}, err
		}

		return objects.Identifier{
			ID: uint64(id),
		}, nil
	} else if err != nil {
		return objects.Identifier{}, err
	}

	// found user name return id
	return objects.Identifier{
		ID: uint64(id),
	}, nil

}
