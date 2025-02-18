package impl

import (
	"context"
	"example.com/go-ecommerce-backend-api/internal/database"
	"example.com/go-ecommerce-backend-api/internal/services"
)

type sUserLogin struct {
	// implement the IUserLogin interface here
	r *database.Queries
}

// Implement the IUserLogin interface here
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}
func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}
func (s *sUserLogin) UpdatePassword(ctx context.Context) error {
	return nil
}

func NewUserLoginImpl(r *database.Queries) services.IUserLogin {
	return &sUserLogin{r: r}
}
