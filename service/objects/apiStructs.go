package objects

// This struct represents a "profile" ad the openAPI specificification describes it
type Profile struct {
	ID             uint64 `json:"user_id"`
	Username       string `json:"username"`
	FollowerCount  uint64 `json:"followers_count"`
	FollowingCount uint64 `json:"following_count"`
	MediaCount     uint64 `json:"media_count"`
}

type CommentPlain struct {
	Text string `json:"comment_text"`
}
type Comment struct {
	ID        uint64 `json:"comment_id"`
	OwnerId   uint64 `json:"owner_id"`
	Username  string `json:"owner_username"`
	Text      string `json:"comment_text"`
	Timestamp string `json:"time_stamp"`
}
type Username struct {
	Text string `json:"username"`
}

type Identifier struct {
	ID uint64 `json:"identifier"`
}
type PhotoMetadata struct {
	ID        uint64    `json:"id"`
	OwnerId   uint64    `json:"owner_id"`
	Comments  []Comment `json:"comments"`
	Likes     []Profile `json:"likes"`
	Timestamp string    `json:"time_stamp"`
}

func (p *Profile) ToDatabase() ProfileDB {
	return ProfileDB{
		ID:             p.ID,
		Username:       p.Username,
		FollowerCount:  p.FollowerCount,
		FollowingCount: p.FollowingCount,
		MediaCount:     p.MediaCount,
	}
}
