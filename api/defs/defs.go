package defs

// response
type SignedUp struct {
	Success   bool
	SessionID string
}

// user information
type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// video
type VideoInfo struct {
	Id  int    `json:"id"`
	Tag string `json:"tag"`
}

// Comment
type Comment struct {
	VideoId int    `json:"video_id"`
	Content string `json:"content"`
}

// Session
type Session struct {
	UserName   string `json:"username"`
	SessionID  string `json:"session_id"`
	ExpireTime int64  `json:"expire_time"`
	CreatedAt  int64  `json:"create_time"`
}

// comment reqs
type CommentReqs struct {
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
}

// create video reqs
type VideoReqs struct {
	AuthodID  string `json:"author_id"`
	VideoName string `json:"name"`
}
