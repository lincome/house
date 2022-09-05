package model

import "time"

// House  。
// 说明:
// 表名:house
// group: House
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.House
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.House
// version:2022-07-29 14:05
type HouseDb struct {
	Id          *int      `gorm:"column:id" json:"Id"`                    //type:*int         comment:id          version:2022-07-29 14:05
	CreatedBy   string    `gorm:"column:created_by" json:"CreatedBy"`     //type:string       comment:创建人      version:2022-07-29 14:05
	CreatedTime time.Time `gorm:"column:created_time" json:"CreatedTime"` //type:*time.Time   comment:创建时间    version:2022-07-29 14:05
	UpdatedBy   string    `gorm:"column:updated_by" json:"UpdatedBy"`     //type:string       comment:更新人      version:2022-07-29 14:05
	UpdatedTime time.Time `gorm:"column:updated_time" json:"UpdatedTime"` //type:*time.Time   comment:更新时间    version:2022-07-29 14:05
	DeletedTime time.Time `gorm:"column:deleted_time" json:"DeletedTime"` //type:*time.Time   comment:删除时间    version:2022-07-29 14:05
	Date        string    `gorm:"column:date" json:"Date"`                //type:*time.Time   comment:            version:2022-07-29 14:05
	Area        string    `gorm:"column:area" json:"area"`                //type:string       comment:            version:2022-07-29 14:05
	RoomArea    string    `gorm:"column:room_area" json:"RoomArea"`       //type:string       comment:            version:2022-07-29 14:05
	RoomNumber  string    `gorm:"column:room_number" json:"RoomNumber"`   //type:*int         comment:            version:2022-07-29 14:05
	Memo        string    `gorm:"column:memo" json:"Memo"`                //type:string       comment:            version:2022-07-29 14:05
}

// TableName Followers 表名
func (HouseDb) TableName() string {
	return "house"
}
