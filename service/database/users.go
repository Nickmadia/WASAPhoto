package database

import (
	obj "WASAPhoto/service/objects"
	"database/sql"
	"errors"
	"fmt"
)

//TODO Add checks for authentication

// Retrives a user profile while checking bans permissions
func (db *appdbimpl) GetUserProfile(id uint64, idReq uint64) (*obj.ProfileDB, error) {
	//TODO encapsulate ban check
	isBanned, err := db.IsBanned(id, idReq)
	if isBanned {
		return nil, ErrUserIsBanned
	} else if err != nil {
		return nil, err
	}
	var user obj.ProfileDB
	err = db.c.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(&user.ID, &user.Username, &user.FollowerCount, &user.FollowingCount, &user.MediaCount)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrProfileDoesNotExist
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// Updates the username for the given id
func (db *appdbimpl) UpdateUsername(id uint64, username string) error {
	var uptQuery = fmt.Sprintf(`UPDATE %s SET username=? WHERE id=?`, USERSTABLE)

	res, err := db.c.Exec(uptQuery, username, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil
	} else if affected == 0 {
		return ErrProfileDoesNotExist
	}
	return nil
}

// Returns a maximum number of profiles that match the given query
func (db *appdbimpl) FetchUsername(username string) ([]obj.ProfileDB, error) {

	var res []obj.ProfileDB
	query := fmt.Sprintf(`SELECT * FROM %s WHERE username LIKE "%s%%"`, USERSTABLE, username)
	rows, err := db.c.Query(query)
	if errors.Is(err, sql.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		var user obj.ProfileDB
		err = rows.Scan(&user.ID, &user.Username, &user.FollowerCount, &user.FollowingCount, &user.MediaCount)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	if err = rows.Err(); err != nil {

		return nil, err
	}

	return res, err
}

// Returns a list of followers or/and following users of the given profile
func (db *appdbimpl) GetUserInfo(id uint64, idReq uint64) ([]obj.ProfileDB, []obj.ProfileDB, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=%d AND ban_id=%d", BANSTABLE, id, idReq)
	err := db.c.QueryRow(query).Scan()
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrUserIsBanned
	}
	var followersList, followingList []obj.ProfileDB

	//TODO change limit and change var
	//following
	query = fmt.Sprintf("SELECT * FROM %s WHERE id=%d LIMIT %d", FOLLOWERSTABLE, id, FETCHLIMIT)
	rows, err := db.c.Query(query)
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		var user obj.ProfileDB
		err = rows.Scan(&user.ID, &user.Username, &user.FollowerCount, &user.FollowingCount, &user.MediaCount)
		if err != nil {
			return nil, nil, err
		}
		followingList = append(followingList, user)
	}

	rows.Close()
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}
	//followers
	rows, err = db.c.Query("SELECT * FROM ? WHERE follower_id=? LIMIT ?", FOLLOWERSTABLE, id, FETCHLIMIT)
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		var user obj.ProfileDB
		err = rows.Scan(&user.ID, &user.Username, &user.FollowerCount, &user.FollowingCount, &user.MediaCount)
		if err != nil {
			return nil, nil, err
		}
		followersList = append(followersList, user)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}
	rows.Close()

	return followersList, followingList, nil
}
