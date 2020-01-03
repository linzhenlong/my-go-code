package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// VideoInfo data model.
type VideoInfo struct {
	ID           int    `json:"id"` // id.
	AuthorID     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
	UUID         string `json:"uuid"`
}
type Comment struct {
	ID        int    `json:"id,omitempty"`
	VideoUUID string `json:"video_uuid,omitempty"`
	AuthorID  int    `json:"author_id,omitempty"`
	Content   string `json:"content,omitempty"`
	Time      string `json:"time,omitempty"`
	UserName  string `json:"user_name,omitempty"`
}
