package defs

// user information
type UserCredential struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
}

// video
type VideoInfo struct {
	Id   int     `json:"id"`
    Tag  string  `json:"tag"`
}

type Comment struct {
	VideoId   int     `json:"video_id"`
	Content   string  `json:"content"`
}