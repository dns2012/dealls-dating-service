package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Unauthenticated(message string) error  {
	return status.New(codes.Unauthenticated, message).Err()
}
