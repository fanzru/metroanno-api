package repo

import (
	errs "metroanno-api/app/accounts/domain/errors"
	"metroanno-api/app/accounts/domain/models"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Impl interface {
	GetUserByUsername(ctx echo.Context, email string) (models.User, error)
	CreateUser(ctx echo.Context, user models.User) (models.User, error)
}
type AccountsRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a AccountsRepo) AccountsRepo {
	return a
}

func (i *AccountsRepo) GetUserByUsername(ctx echo.Context, username string) (models.User, error) {
	var user models.User
	result := i.MySQL.DB.Table("users").Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errs.ErrInstanceNotFound
		}
		return user, result.Error
	}
	if result.RowsAffected < 1 {
		return user, errs.ErrInstanceNotFound
	}
	return user, nil
}

func (i *AccountsRepo) CreateUser(ctx echo.Context, user models.User) (models.User, error) {
	result := i.MySQL.DB.Table("users").Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
