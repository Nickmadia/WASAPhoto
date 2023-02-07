package database

import (
	"fmt"
)

func (db *appdbimpl) IsBanned(id uint64, banId uint64) (bool, error) {
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d AND ban_id=%d",
		BANSTABLE, id, banId)
	var exist int
	err := db.c.QueryRow(query).Scan(&exist)
	if err != nil {
		return false, err
	} else if exist != 0 {
		return true, nil
	}
	return false, nil
}

// TODO implement removing the user from the followers
func (db *appdbimpl) BanUser(idReq uint64, banId uint64) error {
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d",
		USERSTABLE, idReq)
	var exist int
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserIsNotAuthenticated
	} else if err != nil {
		return err
	}
	query = fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d",
		USERSTABLE, banId)

	err = db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserDoesNotExist
	} else if err != nil {
		return err
	}
	query = fmt.Sprintf(`INSERT INTO %s (id, ban_id) VALUES (%d, %d)`,
		BANSTABLE, idReq, banId)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) UnbanUser(idReq uint64, banId uint64) error {
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d",
		USERSTABLE, idReq)
	var exist uint64
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserIsNotAuthenticated
	} else if err != nil {
		return err
	}
	query = fmt.Sprintf(`DELETE FROM %s WHERE id=%d AND ban_id=%d`,
		BANSTABLE, idReq, banId)
	res, err := db.c.Exec(query)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrResourceDoesNotExist
	}
	return nil
}
