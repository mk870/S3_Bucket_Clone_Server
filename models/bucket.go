package models

type Bucket struct {
	MyModel
	Id      int      `json:"id" gorm:"primary_key"`
	UserId  int      `json:"userId"`
	Name    string   `json:"name"`
	Folders []Folder `gorm:"ForeignKey:BucketId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Objects []Object `gorm:"ForeignKey:BucketId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
