package calculservice

import (
	"context"
	"log"

	calcul "github.com/idirall22/grpc_calcul/api/pb"
)

// CalculService struct
type CalculService struct {
	calcul.UnsafeCalculServiceServer
}

// NewService create CalculService.
func NewService() *CalculService {
	return &CalculService{}
}

// Add add two number and return result.
func (s *CalculService) Add(ctx context.Context, in *calcul.AddReq) (*calcul.AddRes, error) {
	log.Printf("New Add request with params, A+B= %d+%d", in.A, in.B)
	var result int64
	result = in.A + in.B
	return &calcul.AddRes{Result: result}, nil
}
