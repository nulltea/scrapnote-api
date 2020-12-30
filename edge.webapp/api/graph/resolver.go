package graph

import "github.com/timoth-y/scrapnote-api/edge.webapp/core/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	records service.RecordService
}

func NewResolver(records service.RecordService) *Resolver {
	return &Resolver{
		records: records,
	}
}