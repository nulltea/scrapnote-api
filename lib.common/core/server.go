package core

import (
	"crypto/rsa"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"

	"github.com/timoth-y/scrapnote-api/lib.common/core/meta"
)

type Server interface {
	SetupEncryption(cert *meta.TLSCertificate)
	SetupAuth(pb *rsa.PublicKey, accessRoles meta.AccessConfig) // Must be configured before rest & gRPC sub servers!
	SetupREST(router chi.Router) // Setup REST sub server configuration
	SetupGRPC(fn func(srv *grpc.Server)) // Setup gRPC sub server configuration
	SetupAMQP(handler Handler) // Setup AMQP events handler sub server configuration
	SetupGrQL(handler *handler.Server) // Setup GraphQL sub server configuration
	SetupLogger()
	Start()
	Shutdown()
}
