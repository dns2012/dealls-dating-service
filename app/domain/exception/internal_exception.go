package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Internal(message string) error  {
	return status.New(codes.Internal, message).Err()
}
