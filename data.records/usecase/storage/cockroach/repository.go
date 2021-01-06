package cockroach

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/timoth-y/scrapnote-api/lib.common/config"
	"github.com/timoth-y/scrapnote-api/lib.common/util"

	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"github.com/timoth-y/scrapnote-api/data.records/core/repo"
)

type repository struct {
	db *sqlx.DB
	table string
}

func (r repository) RetrieveAll() ([]*model.Record, error) {
	panic("implement me")
}

func NewRepository(config config.DataStoreConfig) (repo.RecordRepository, error) {
	db, err := newClient(config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo := &repository{
		db: db,
		table:  config.Collection,
	}
	return repo, nil
}

func newClient(url string) (*sqlx.DB, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := sqlx.ConnectContext(ctx,"pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "repository.newClient")
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "repository.newPostgresClient")
	}
	return db, nil
}

func (r repository) RetrieveOne(id string) (*model.Record, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sneakerReference := &model.Record{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"unique_id":id}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Record.RetrieveOne")
	}
	if err = r.db.GetContext(ctx, sneakerReference, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.Record.RetrieveOne")
	}
	return sneakerReference, nil
}

func (r repository) RetrieveBy(topic string) ([]*model.Record, error) {
	panic("implement me")
}

func (r repository) Retrieve(ids []string) ([]*model.Record, error) {
	panic("implement me")
}

func (r repository) Store(record *model.Record)  error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(record)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.Record.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.Record.Store")
	}
	return nil
}

func (r repository) Modify(record *model.Record) error {
	panic("implement me")
}

func (r repository) Remove(id string) error {
	panic("implement me")
}