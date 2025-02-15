package repo

import (
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func (*userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
