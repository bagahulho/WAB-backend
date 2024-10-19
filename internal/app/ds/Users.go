package ds

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Login       string `gorm:"type:varchar(25);unique;not null" json:"login"`
	Password    string `gorm:"type:varchar(100);not null" json:"password"`
	IsModerator bool   `gorm:"type:boolean;default:false" json:"is_moderator"`
}
