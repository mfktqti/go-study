package model

type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parentID"`
	Url      string `json:"url"`
}
