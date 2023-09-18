package gapi

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

func (*LoginRpcServer) ServiceHealth(context.Context, *empty.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Hello World"),
		Extensions:  nil,
	}, nil
}
