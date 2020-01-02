package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`

}

// VideoInfo data modle.
type VideoInfo struct {
	ID int  `json:"id"`// id.
	AuthorID int `json:"author_id"`
	Name string  `json:"name"`
	DisplayCtime string `json:"display_ctime"`
	UUID string `json:"uuid"`
}
