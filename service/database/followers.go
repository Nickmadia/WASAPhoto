package database

import (
	"fmt"
)

func (db *appdbimpl) FollowUser(idReq uint64, followedId uint64) error {
	var exist uint64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d",
		USERSTABLE, idReq)
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserIsNotAuthenticated
	} else if err != nil {
		return err
	}

	// check if the follow target exist
	query = fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d",
		USERSTABLE, followedId)
	err = db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserDoesNotExist
	} else if err != nil {
		return err
	}
	isBanned, err := db.IsBanned(followedId, idReq)
	if err != nil {
		return err
	}
	if isBanned {
		return ErrUserIsBanned
	}
	// check if he already follows
	query = fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d AND follow_id=%d",
		FOLLOWERSTABLE, idReq, followedId)
	err = db.c.QueryRow(query).Scan(&exist)
	if err != nil {
		return err
	}
	if exist == 1 {
		return nil
	}
	// TODO ENFORCE PRIMARY KEYS
	query = fmt.Sprintf(`INSERT INTO %s (id, follow_id) VALUES (%d,%d)`,
		FOLLOWERSTABLE, idReq, followedId)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}

	// update followers number
	query = fmt.Sprintf(`UPDATE %s SET followers_count = followers_count + 1 WHERE id=%d `,
		USERSTABLE, followedId)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf(`UPDATE %s SET following_count = following_count + 1 WHERE id=%d `,
		USERSTABLE, idReq)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) UnfollowUser(idReq uint64, followedId uint64) error {
	var exist uint64

	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d",
		USERSTABLE, idReq)
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserIsNotAuthenticated
	} else if err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id=%d AND follow_id=%d`,
		FOLLOWERSTABLE, idReq, followedId)
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
	query = fmt.Sprintf(`UPDATE %s SET followers_count = followers_count - 1 WHERE id=%d `,
		USERSTABLE, followedId)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf(`UPDATE %s SET following_count = following_count - 1 WHERE id=%d `,
		USERSTABLE, idReq)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
