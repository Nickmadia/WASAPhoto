package database

import (
	obj "WASAPhoto/service/objects"
	"fmt"
	"time"
)

// methods for managing likes
func (db *appdbimpl) LikeMedia(idReq uint64, postId uint64) error {
	//TODO check if the user is authenticated + ENCAPS
	var exist uint64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", USERSTABLE, idReq)
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrUserIsNotAuthenticated
	} else if err != nil {
		return err
	}
	//check if the post exist
	query = fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", MEDIATABLE, postId)
	err = db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return ErrResourceDoesNotExist
	} else if err != nil {
		return err
	}
	query = fmt.Sprintf("SELECT count(*) FROM %s WHERE media_id=%d and user_id=%d ", LIKESTABLE, postId, idReq)
	err = db.c.QueryRow(query).Scan(&exist)
	if exist == 1 {
		return nil
	} else if err != nil {
		return err
	}
	query = fmt.Sprintf(`INSERT INTO %s (media_id, user_id) VALUES  (%d,%d)`,
		LIKESTABLE, postId, idReq)
	_, err = db.c.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) UnlikeMedia(idReq uint64, postId uint64) error {
	var exist uint64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", USERSTABLE, idReq)
	err := db.c.QueryRow(query).Scan(&exist)
	if err != nil {
		return err
	} else if exist != 1 {
		return ErrUserIsNotAuthenticated
	}
	query = fmt.Sprintf(`DELETE FROM %s WHERE media_id=%d AND user_id=%d `,
		LIKESTABLE, postId, idReq)
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

func (db *appdbimpl) CommentMedia(idReq uint64, postId uint64, text string) (int64, error) {
	var exist uint64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", USERSTABLE, idReq)
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return 0, ErrUserIsNotAuthenticated
	} else if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", MEDIATABLE, postId)
	//check if the post exist
	err = db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return 0, ErrResourceDoesNotExist
	} else if err != nil {
		return 0, err
	}

	//Getting the timestamp + adding comment
	ts := time.Now().Unix()
	query = fmt.Sprintf(`INSERT INTO %s ( owner_id, media_id, text , time_stamp) VALUES (%d, %d, "%s", %d)`,
		COMMENTSTABLE, idReq, postId, text, ts)
	res, err := db.c.Exec(query)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (db *appdbimpl) UncommentMedia(idReq uint64, postId uint64, commentId uint64) error {
	var exist uint64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", USERSTABLE, idReq)
	err := db.c.QueryRow(query).Scan(&exist)
	if err != nil {
		return err
	} else if exist != 1 {
		return ErrUserIsNotAuthenticated
	}
	query = fmt.Sprintf(`DELETE FROM %s WHERE id=%d AND media_id=%d AND owner_id=%d `,
		COMMENTSTABLE, commentId, postId, idReq)
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

// gets a list of comments without the one made by the banend users
func (db *appdbimpl) GetComments(idReq uint64, postId uint64) (*[]obj.Comment, error) {
	var exist int
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", MEDIATABLE, postId)
	err := db.c.QueryRow(query).Scan(&exist)
	if exist != 1 {
		return nil, ErrResourceDoesNotExist
	} else if err != nil {
		return nil, err
	}

	//query the result while removing comments from banned users
	//TODO put everything in one query
	query = fmt.Sprintf(`SELECT id, owner_id, text, time_stamp 
						FROM %s 
						WHERE media_id=%d AND owner_id NOT IN 
						(SELECT ban_id FROM %s WHERE id=%d)`, COMMENTSTABLE, postId, BANSTABLE, idReq)
	rows, err := db.c.Query(query)
	if err != nil {
		return nil, err
	}
	var commentsList []obj.Comment
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		var comment obj.Comment
		var ts int64
		err = rows.Scan(&comment.ID, &comment.OwnerId, &comment.Text, &ts)
		if err != nil {
			return nil, err
		}
		//TODO encapsulate date time converter in a utils package
		comment.Timestamp = time.Unix(ts, 0).Format("2006-01-02T15:04:05.999Z")
		commentsList = append(commentsList, comment)
	}
	for x, y := range commentsList {
		query = fmt.Sprintf(`SELECT username from %s WHERE id=%d`, USERSTABLE, y.OwnerId)
		var tmp string
		err = db.c.QueryRow(query).Scan(&tmp)
		if err != nil {
			return nil, err
		}

		commentsList[x].Username = tmp
	}
	return &commentsList, nil

}

func (db *appdbimpl) GetLikes(idReq uint64, postId uint64) (*[]obj.Profile, error) {
	var exist int
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d", MEDIATABLE, postId)
	err := db.c.QueryRow(query).Scan(&exist)
	if err != nil {
		return nil, err
	} else if exist != 1 {
		return nil, ErrResourceDoesNotExist
	}

	//query the result while removing profiles from banned users
	query = fmt.Sprintf(`SELECT u.id, u.username, u.followers_count, u.following_count, u.media_count 
						FROM %s AS u 
						JOIN %s AS l ON u.id = l.user_id
						WHERE l.media_id=%d AND l.user_id NOT IN (SELECT ban_id FROM %s WHERE id=%d)`,
		USERSTABLE, LIKESTABLE, postId, BANSTABLE, idReq)
	rows, err := db.c.Query(query)
	if err != nil {
		return nil, err
	}
	//TODO keep an eye on pointed list that could be erased after function calling
	var likesList = new([]obj.Profile)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		var like obj.Profile
		err = rows.Scan(&like.ID, &like.Username, &like.FollowerCount, &like.FollowingCount, &like.MediaCount)
		if err != nil {
			return nil, nil
		}
		*likesList = append(*likesList, like)
	}

	return likesList, nil

}
