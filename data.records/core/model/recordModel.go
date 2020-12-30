package model

import "time"

type Record struct {
	UniqueID  string    `json:"uniqueID" bson:"unique_id"`
	TopicID  string    `json:"topicID" bson:"topic_id"`
	Content   string    `json:"content" bson:"content"`
	SourceURL string    `json:"sourceURL" bson:"source_url"`
	MarkerURL string    `json:"markerURL" bson:"marker_url"`
	AddedAt   time.Time `json:"addedAt" bson:"added_at"`
}