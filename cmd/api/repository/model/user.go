package model

type User struct {
	Id             uint64 `gorm:"primary_key"`
	Username       string `gorm:"index;not null"`
	Email          string `gorm:"unique_index;not null"`
	Password       string
}
