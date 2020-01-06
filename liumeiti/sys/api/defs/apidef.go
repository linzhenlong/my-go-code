package defs


// UserCredential 用户信息.
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// response
// SignedUp 信息.
type SignedUp struct {
	Suceess   bool   `json:"suceess,omitempty"`
	SessionID string `json:"session_id,omitempty"`
}

// VideoInfo data model.
type VideoInfo struct {
	ID           int    `json:"id"` // id.
	AuthorID     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
	UUID         string `json:"uuid"`
}
//Comment  评论结构体.
type Comment struct {
	ID        int    `json:"id,omitempty"`
	VideoUUID string `json:"video_uuid,omitempty"`
	AuthorID  int    `json:"author_id,omitempty"`
	Content   string `json:"content,omitempty"`
	Time      string `json:"time,omitempty"`
	UserName  string `json:"user_name,omitempty"`
}
// SimpleSession session 相关结构体.
type SimpleSession struct {
	UserName string `json:"user_name,omitempty"`
	TTL      int64  `json:"ttl,omitempty"`

}