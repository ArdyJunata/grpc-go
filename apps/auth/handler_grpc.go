package auth

import (
	"context"

	auth "github.com/ArdyJunata/grpc-go/apps/auth/proto"
	internalerror "github.com/ArdyJunata/grpc-go/internal/error"
	"google.golang.org/grpc/status"
)

func newHandlerGrpc(svc service) handlerGrpc {
	return handlerGrpc{
		svc: svc,
	}
}

type handlerGrpc struct {
	svc service
	auth.UnimplementedAuthServiceServer
}

// Register implements auth.AuthServiceServer.
func (h handlerGrpc) Register(ctx context.Context, req *auth.RegisterRequest) (resp *auth.RegisterResponse, err error) {
	reqData := registerRequest{
		Username: req.Username,
		Password: req.Password,
	}

	err = h.svc.register(ctx, reqData)
	if err != nil {
		errMsg, ok := internalerror.ErrorBase[err.Error()]
		if !ok {
			errMsg = internalerror.ErrorInternal
		}

		err = status.Error(errMsg.Code, errMsg.Error())

		return nil, err
	}

	resp = &auth.RegisterResponse{
		Success: true,
	}

	return
}

// Login implements auth.AuthServiceServer.
func (h handlerGrpc) Login(ctx context.Context, req *auth.LoginRequest) (resp *auth.LoginResponse, err error) {
	reqData := loginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := h.svc.login(ctx, reqData)
	if err != nil {
		errMsg, ok := internalerror.ErrorBase[err.Error()]
		if !ok {
			errMsg = internalerror.ErrorUnauthenticated
		}

		err = status.Error(errMsg.Code, errMsg.Error())

		return nil, err
	}

	resp = &auth.LoginResponse{
		Token: token,
	}

	return
}
