package repo

import (
	errs "metroanno-api/app/accounts/domain/errors"
	"metroanno-api/app/accounts/domain/models"
	"metroanno-api/app/accounts/domain/response"
	modelsannotation "metroanno-api/app/annotation/domain/models"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Impl interface {
	GetUserByUsername(ctx echo.Context, email string) (models.User, error)
	CreateUser(ctx echo.Context, user models.User) (models.User, error)
	GetUserByID(ctx echo.Context, ID int64) (models.User, error)
	GetDocumentsById(ctx echo.Context, documentId int64) (*modelsannotation.Document, error)
	UpdateStatusUsers(ctx echo.Context, userID int64, status string) error
	GetAllUserNonAdmin(ctx echo.Context, pageNumber int64) (response.Pagination, error)
	CreateUsersSubjects(ctx echo.Context, usersSubjects []*models.UsersSubjects) ([]*models.UsersSubjects, error)
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

func (i *AccountsRepo) GetUserByID(ctx echo.Context, ID int64) (models.User, error) {
	var user models.User
	result := i.MySQL.DB.Table("users").Where("id = ?", ID).First(&user)
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

func (a *AccountsRepo) GetDocumentsById(ctx echo.Context, documentId int64) (*modelsannotation.Document, error) {
	document := &modelsannotation.Document{}
	result := a.MySQL.DB.Table("documents").Where("id = ?", documentId).First(&document)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return document, nil
}

func (i *AccountsRepo) CreateUser(ctx echo.Context, user models.User, us []*models.UsersSubjects) (models.User, error) {
	tx := i.MySQL.DB.Begin()
	result := tx.Table("users").Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return user, result.Error
	}
	for _, v := range us {
		v.UserId = user.Id
	}

	result = tx.Table("users_subjects").CreateInBatches(&us, 10)
	if result.Error != nil {
		tx.Rollback()
		return user, result.Error
	}

	tx.Commit()
	return user, nil
}

func (i *AccountsRepo) CreateUsersSubjects(ctx echo.Context, usersSubjects []*models.UsersSubjects) ([]*models.UsersSubjects, error) {
	result := i.MySQL.DB.Table("users_subjects").CreateInBatches(&usersSubjects, 10)
	if result.Error != nil {
		return nil, result.Error
	}
	return usersSubjects, nil
}

func (i *AccountsRepo) UpdateStatusUsers(ctx echo.Context, userID int64, status string) error {
	result := i.MySQL.DB.Table("users").Where("id = ?", userID).UpdateColumn("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (i *AccountsRepo) GetAllUserNonAdmin(ctx echo.Context, pageNumber int64) (response.Pagination, error) {
	resp := response.Pagination{}
	var totalRows int64

	err := i.MySQL.DB.Model(&models.User{}).Where("type = ?", 1).Count(&totalRows).Error
	if err != nil {
		return resp, err
	}

	// calculate page dan offset
	// contoh halaman yang diminta
	pageSize := 10 // contoh ukuran halaman
	offset := (int(pageNumber) - 1) * pageSize

	users := []*models.UserWithoutPassword{}
	result := i.MySQL.DB.Table("users").Where("type = ?", 1).Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return resp, result.Error
	}

	for _, user := range users {
		var totalAllChecked, totalAll int64

		err := i.MySQL.DB.Table("question_annotations").Where("user_id = ?", user.Id).Count(&totalAll).Error
		if err != nil {
			return resp, err
		}

		err = i.MySQL.DB.Table("question_annotations").Where("user_id = ? AND is_checked_admin = ?", user.Id, 1).Count(&totalAllChecked).Error
		if err != nil {
			return resp, err
		}

		user.TotalQuestionCheckedAdmin = totalAllChecked
		user.TotalQuestion = totalAll

	}
	var prevPage, nextPage int64
	if pageNumber > 1 {
		prevPage = pageNumber - 1
	}
	if int(offset)+len(users) < int(totalRows) {
		nextPage = pageNumber + 1
	}

	var totalPages int64
	if totalRows%int64(pageSize) == 0 {
		totalPages = totalRows / int64(pageSize)
	} else {
		totalPages = (totalRows / int64(pageSize)) + 1
	}

	var start int64 = 1
	if totalRows == 0 {
		start = 0
	}
	resp = response.Pagination{
		Page:  pageNumber,
		Limit: int64(pageSize),
		Prev:  prevPage,
		Next:  nextPage,
		Start: start,
		End:   totalPages,
		Data:  users,
	}

	return resp, nil
}
