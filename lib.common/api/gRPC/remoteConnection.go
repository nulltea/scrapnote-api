package gRPC

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/timoth-y/scrapnote-api/lib.common/config"
)

func NewRemoteConnection(config config.ConnectionConfig) *grpc.ClientConn {
	var opts []grpc.DialOption
	if config.TLS != nil && config.TLS.EnableTLS {
		tls, err := LoadClientTLSCredentials(config.TLS); if err != nil {
			glog.Fatalln("cannot load TLS credentials: ", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(tls))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(config.URL, opts...); if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}
	return conn
}