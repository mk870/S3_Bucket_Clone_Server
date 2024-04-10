package models

type Object struct {
	MyModel
	Id       int    `json:"id" gorm:"primary_key"`
	BucketId int    `json:"bucketId"`
	FolderId int    `json:"folderId"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Size     string `json:"size"`
	Uri      string `json:"uri"`
}
