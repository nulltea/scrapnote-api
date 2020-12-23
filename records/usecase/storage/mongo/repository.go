package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	"go.kicksware.com/api/service-common/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/core/meta"

	"github.com/timoth-y/scrapnote-api/records/core/model"
	"github.com/timoth-y/scrapnote-api/records/core/repo"
)

type repository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}


func NewRepository(config config.DataStoreConfig) (repo.RecordRepository, error) {
	repo := &repository{
		timeout:  time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config); if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	repo.collection = database.Collection(config.Collection)
	return repo, nil
}

func newMongoClient(config config.DataStoreConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(config.URL).
		SetTLSConfig(newTLSConfig(config.TLS)).
		SetAuth(options.Credential{
			Username: config.Login, Password: config.Password,
		}),
	)
	err = client.Ping(ctx, readpref.Primary()); if err != nil {
		return nil, err
	}
	return client, nil
}

func newTLSConfig(tlsConfig *meta.TLSCertificate) *tls.Config {
	if !tlsConfig.EnableTLS {
		return nil
	}
	certs := x509.NewCertPool()
	pem, err := ioutil.ReadFile(tlsConfig.CertFile); if err != nil {
		glog.Fatalln(err)
	}
	certs.AppendCertsFromPEM(pem)
	return &tls.Config{
		RootCAs: certs,
	}
}

func (r *repository) RetrieveOne(id string) (*model.Record, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := r.buildQueryPipeline(bson.M{"unique_id": id})
	cursor, err := r.collection.Aggregate(ctx, query); if err != nil {
		return nil, errors.Wrap(err, "repository.Record.FetchOne")
	}
	defer cursor.Close(ctx)

	var orders []*model.Record
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, errors.Wrap(err, "repository.Record.FetchOne")
	}
	if orders == nil || len(orders) == 0 {
		if err == mongo.ErrNoDocuments{
			return nil, errors.Wrap(err, "repository.Record.FetchOne")
		}
		return nil, errors.Wrap(err, "repository.Record.FetchOne")
	}
	return orders[0], nil
}

func (r *repository) Retrieve(topic string) ([]*model.Record, error) {
	panic("implement me")
}

func (r *repository) Store(record *model.Record) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		return errors.Wrap(err, "repository.Record.Store")
	}
	return nil
}

func (r *repository) Modify(record *model.Record) error {
	panic("implement me")
}

func (r *repository) Remove(id string) error {
	panic("implement me")
}

func (r *repository) buildQueryPipeline(matchQuery bson.M) mongo.Pipeline {
	pipe := mongo.Pipeline{}
	pipe = append(pipe, bson.D{{"$match", matchQuery}})

	return pipe
}

