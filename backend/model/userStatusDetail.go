package model

type UserStatusDetail struct {
	Username     string `gorm:"type:varchar(255);default:null"`
	ContestID    string `gorm:"type:varchar(255);default:null"`
	ProblemID    string `gorm:"type:varchar(255);default:null"`
	TeacherLabel int    `gorm:"default:null"`
	IsLock       int    `gorm:"default:null"`
	IsUnlock     int    `gorm:"default:null"`
}

func (UserStatusDetail) TableName() string {
	return "userstatus_detail"
}
