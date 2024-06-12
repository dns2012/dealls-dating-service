package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Invalid(message string) error  {
	return status.New(codes.InvalidArgument, message).Err()
}
