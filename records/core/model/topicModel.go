package model

type Topic struct {
	Records []Record `json:"Record" bson:"record"`

}
