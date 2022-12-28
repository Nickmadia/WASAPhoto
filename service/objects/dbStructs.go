package objects

//Database representation of a profile
type ProfileDB struct {
	ID             uint64
	Username       string
	FollowerCount  uint64
	FollowingCount uint64
	MediaCount     uint64
}
type PhotoMetadataDB struct {
	ID            uint64
	OwnerId       uint64
	CommentsCount uint
	Comments      []uint64
	LikesCount    uint
	Likes         []Profile
}

func (p *ProfileDB) FromDatabase() Profile {
	return Profile{
		ID:             p.ID,
		Username:       p.Username,
		FollowerCount:  p.FollowerCount,
		FollowingCount: p.FollowingCount,
		MediaCount:     p.MediaCount,
	}
}
