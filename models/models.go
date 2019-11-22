package models

type Note struct {
	ID          string `json:"id,omitempty"`
	Timestamp   int `json:"time,omitempty"`
	Content    	string `json:"content,omitempty"`
	User `json:"User,omitempty"`
}
type User struct {	
	ID 			string `json:"id,omitempty"`
	Name 		string `json:"username,omitempty"`
}
