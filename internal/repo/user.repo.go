package repo

import (
	"database/sql"
	"errors"
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/database"
	"fmt"
)

// import "context"
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	// h phai ket noi toi sqlc
	sqlc *database.Queries
}

func (up *userRepository) GetUserByEmail(email string) bool {
	if up.sqlc == nil {
		fmt.Println("sqlc is nil")
		return false
	}

	//ctx := context.Background() // Đảm bảo có ctx
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Không tìm thấy user:", email)
			return false // User không tồn tại
		}
		fmt.Println("Lỗi truy vấn database:", err)
		return false
	}

	fmt.Println("User tìm thấy:", user)
	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
