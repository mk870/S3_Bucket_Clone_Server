package models

type Folder struct {
	MyModel
	Id       int      `json:"id" gorm:"primary_key"`
	BucketId int      `json:"bucketId"`
	Name     string   `json:"name"`
	Objects  []Object `gorm:"ForeignKey:FolderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
