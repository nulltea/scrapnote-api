package model

import "time"

type User struct {
	UniqueID           string                  `json:"UniqueID" bson:"unique_id"`
	Username           string                  `json:"Username" bson:"username"`
	Email              string                  `json:"Email" bson:"email"`
	PasswordHash       string                  `json:"PasswordHash" bson:"password_hash"`
	FirstName          string                  `json:"FirstName" bson:"first_name"`
	LastName           string                  `json:"LastName" bson:"last_name"`
	PhoneNumber        string                  `json:"PhoneNumber" bson:"phone_number"`
	Avatar             string                  `json:"Avatar" bson:"avatar"`
	Location           string                  `json:"Location" bson:"location"`
	Confirmed          bool                    `json:"Confirmed" bson:"confirmed"`
	RegisterDate       time.Time               `json:"RegisterDate" bson:"register_date"`
}