/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"WASAPhoto/service/objects"
	"database/sql"
	"errors"
	"fmt"
)

// Const for db's table names
const USERSTABLE string = "users"
const FOLLOWERSTABLE string = "followers"
const MEDIATABLE string = "media"
const LIKESTABLE string = "likes"
const COMMENTSTABLE string = "comments"
const BANSTABLE string = "bans"

// Common errors that could happen while interacting with the db
var ErrProfileDoesNotExist = errors.New("profile does not exist")
var ErrResourceDoesNotExist = errors.New("resource does not exist")
var ErrUserIsBanned = errors.New("user is Banned")
var ErrUserDoesNotOwnTheResource = errors.New("the user doesn't own the resource")
var ErrUserDoesNotExist = errors.New("user does now exists")
var ErrUserIsNotAuthenticated = errors.New("user must be authenticated")

// API LIMITS
const FETCHLIMIT = 20

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	//Methods for managing the users
	GetUserProfile(id uint64, idReq uint64) (*objects.ProfileDB, error)
	UpdateUsername(id uint64, username string) error
	FetchUsername(username string) ([]objects.ProfileDB, error)
	GetUserInfo(id uint64, idReq uint64) ([]objects.ProfileDB, []objects.ProfileDB, error)

	//Methods for managing media uploads and downloads
	UploadImage(id uint64, img *string) (uint64, error)
	GetMedia(idReq uint64, postId uint64) (*string, error)
	GetMediaMetadata(idReq uint64, postId uint64) (*objects.PhotoMetadata, error)
	DeleteMedia(idReq uint64, postId uint64) error

	//methods for managing likes
	LikeMedia(idReq uint64, postId uint64) error
	UnlikeMedia(idReq uint64, postId uint64) error

	//methods for managing comments
	CommentMedia(idReq uint64, postId uint64, text string) (int64, error)
	UncommentMedia(idReq uint64, postId uint64, commentId uint64) error

	//follow
	FollowUser(idReq uint64, followedId uint64) error
	UnfollowUser(idReq uint64, followedId uint64) error

	//bans
	IsBanned(id uint64, idBanner uint64) (bool, error)
	BanUser(idReq uint64, banId uint64) error
	UnbanUser(idReq uint64, banId uint64) error

	//Login
	SignInOrLogin(username string) (objects.Identifier, error)

	//Stream
	GetStream(idReq uint64) ([]objects.PhotoMetadata, error)

	//ping
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	//TODO drop prebad written table
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		userTblStmt := `CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL,
			 followers_count INTEGER NOT NULL DEFAULT 0, following_count INTEGER NOT NULL DEFAULT 0, 
			 media_count INTEGER NOT NULL DEFAULT 0);`
		followersTbl := `CREATE TABLE followers (id INTEGER NOT NULL, follow_id INTEGER NOT NULL, PRIMARY KEY (id,follow_id))`

		commentsTbl := `CREATE TABLE comments(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, owner_id INTEGER NOT NULL,
			media_id INTEGER NOT NULL, text TEXT NOT NULL, time_stamp INTEGER NOT NULL)`

		likesTbl := `CREATE TABLE likes(media_id INTEGER NOT NULL, user_id INTEGER NOT NULL, PRIMARY KEY (media_id, user_id))`

		bansTbl := `CREATE TABLE bans (id INTEGER NOT NULL, ban_id INTEGER NOT NULL, PRIMARY KEY(id, ban_id))`

		mediaTbl := `CREATE TABLE media (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, owner_id INTEGER NOT NULL, 
				png TEXT NOT NULL, time_stamp INTEGER NOT NULL)`

		err := execMultipleQuerys(db, userTblStmt, followersTbl, commentsTbl, likesTbl, bansTbl, mediaTbl)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func execMultipleQuerys(db *sql.DB, s ...string) error {
	for _, el := range s {
		_, err := db.Exec(el)

		if err != nil {

			return fmt.Errorf("error creating db structurev: %w", err)
		}
	}
	return nil
}
