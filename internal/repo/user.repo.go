package repo

import (
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	// h phai ket noi toi sqlc
	sqlc *database.Queries
}

func (up *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	//row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}

	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
