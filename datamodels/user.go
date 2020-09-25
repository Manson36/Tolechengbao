package datamodels

import "time"

type User struct {
	ID 			int64 `json:"id, string"`
	Username 	string `json:"username" gorm:"type:varchar(50);not null; unique; pq_comment:用户名"`
	Nickname 	string `json:"nickname" gorm:"type: varchar(50); not null; pq_comment:用户昵称"`
	Password	string `json:"password" gorm:"type: varchar(200); not null; pq_comment:用户密码"`

	CreateAt     *time.Time   `json:"createAt" gorm:"type:timestamptz;not null;default:now();pq_comment:该用户的创建时间"`
	UpdateAt     *time.Time   `json:"updateAt" gorm:"type:timestamptz;default:now();pq_comment:该用户的更新时间"`
	RemovedAt    *time.Time   `json:"removedAt" gorm:"type:timestamptz;pq_comment:用户的移除时间"`
	Removed      bool         `json:"removed" gorm:"pq_comment:该用户是否被移除"`
	Disabled     bool         `json:"disabled" gorm:"pq_comment:该用户是否被禁用"`
	DisabledAt   *time.Time   `json:"disabledAt" gorm:"type:timestamptz;pq_comment:该用户被禁用的时间"`
}
