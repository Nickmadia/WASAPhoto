package database

import (
	"WASAPhoto/service/objects"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func (db *appdbimpl) UploadImage(id uint64, img *string) (uint64, error) {
	var exist uint64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", USERSTABLE, id)
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return 0, ErrUserIsNotAuthenticated
	} else if err != nil {
		return 0, err
	}
	// Getting the timestamp
	ts := time.Now().Unix()
	query = fmt.Sprintf(`INSERT INTO %s (owner_id, png, time_stamp) VALUES (%d,"%s", %d)`,
		MEDIATABLE, id, *img, ts)
	res, err := db.c.Exec(query)
	if err != nil {
		return 0, err
	}

	upId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(upId), nil
}
func (db *appdbimpl) GetMedia(idReq uint64, postId uint64) (*string, error) {
	// TODO user auth
	// TODO check bans in all the file
	query := fmt.Sprintf(`SELECT png FROM %s WHERE id=%d`, MEDIATABLE, postId)
	var b64Img string
	err := db.c.QueryRow(query).Scan(&b64Img)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrResourceDoesNotExist
	} else if err != nil {
		return nil, err
	}
	return &b64Img, nil

}
func (db *appdbimpl) GetMediaMetadata(idReq uint64, postId uint64) (*objects.PhotoMetadata, error) {
	// TODO Check if user is auth
	var res = new(objects.PhotoMetadata)
	var ts int64
	query := fmt.Sprintf(`SELECT id, owner_id, time_stamp  FROM %s WHERE id=%d`, MEDIATABLE, postId)
	err := db.c.QueryRow(query).
		Scan(&res.ID, &res.OwnerId, &ts)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrResourceDoesNotExist
	} else if err != nil {
		return nil, err
	}
	res.Timestamp = time.Unix(ts, 0).Format("2006-01-02T15:04:05.999Z")
	commmestList, err := db.GetComments(idReq, postId)
	if err != nil {
		return nil, err
	}
	likesList, err := db.GetLikes(idReq, postId)
	if err != nil {
		return nil, err
	}
	res.Comments = *commmestList
	res.Likes = *likesList
	return res, nil

}
func (db *appdbimpl) DeleteMedia(idReq uint64, postId uint64) error {
	// check allso if the user is the actual owner
	query := fmt.Sprintf(`SELECT id FROM %s WHERE id=%d AND owner_id=%d`, MEDIATABLE, postId, idReq)
	var exist int
	err := db.c.QueryRow(query).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrResourceDoesNotExist
	} else if err != nil {
		return err
	}
	query = fmt.Sprintf(`DELETE FROM %s WHERE id=%d AND owner_id=%d `,
		MEDIATABLE, postId, idReq)
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
