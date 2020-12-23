package model

import "time"

type Record struct {
	UniqueID  string    `json:"UniqueID" bson:"unique_id"`
	Content   string    `json:"Content" bson:"content"`
	SourceURL string    `json:"SourceURL" bson:"source_url"`
	MarkerURL string    `json:"MarkerURL" bson:"marker_url"`
	AddedAt   time.Time `json:"AddedAt" bson:"added_at"`
}