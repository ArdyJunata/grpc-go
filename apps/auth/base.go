package auth

import (
	pb "github.com/ArdyJunata/grpc-go/apps/auth/proto"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func RouterInitGRPC(server *grpc.Server, db *sqlx.DB) {
	repository := newRepository(db)
	service := newService(repository)
	handler := newHandlerGrpc(service)

	pb.RegisterAuthServiceServer(server, handler)
}
