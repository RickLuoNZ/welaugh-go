package entity

import (
	"time"
)

type User struct {
	Id             uint64 
	Username       string 
	Email          string 
	Password       string 
	PhoneNumber    string 
	CountryCode    string 
	Avatar         string
	CoverImage     string
	Signature      string
	Rank           uint8    
	IsMember       uint8  
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time   
}


