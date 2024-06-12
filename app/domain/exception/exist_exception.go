package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Exist(message string) error  {
	return status.New(codes.AlreadyExists, message).Err()
}
