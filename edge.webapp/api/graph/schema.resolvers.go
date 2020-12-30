package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	native "github.com/timoth-y/scrapnote-api/data.records/core/model"

	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph/generated"
	"github.com/timoth-y/scrapnote-api/edge.webapp/core/model"
)

func (r *mutationResolver) ModifyRecord(ctx context.Context, input model.RecordInput) (*native.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Records(ctx context.Context, topic string) ([]*native.Record, error) {
	return r.records.GetFrom(topic)
}

func (r *recordResolver) TopicID(ctx context.Context, obj *native.Record) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Record returns generated.RecordResolver implementation.
func (r *Resolver) Record() generated.RecordResolver { return &recordResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type recordResolver struct{ *Resolver }
