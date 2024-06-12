package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NotFound(message string) error  {
	return status.New(codes.NotFound, message).Err()
}
