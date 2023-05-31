package usecase

import (
	"fmt"
	errs "metroanno-api/app/accounts/domain/errors"
	"metroanno-api/app/accounts/domain/models"
	"metroanno-api/app/accounts/domain/request"
	"metroanno-api/app/accounts/domain/response"
	"metroanno-api/app/accounts/repo"
	"metroanno-api/infrastructure/config"
	"metroanno-api/pkg/jwt"
	"strconv"

	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Impl interface {
	UserRegister(ctx echo.Context, param request.UserRegisterReq) error
	UserLogin(ctx echo.Context, param request.UserLoginReq) (*response.UserLoginRes, error)
	UserProfile(ctx echo.Context) (*response.ProfileRes, error)
	UpdateStatusUsers(ctx echo.Context, userID int64, status string) error
	GetAllUserNonAdmin(ctx echo.Context, pageNumber int64) (response.Pagination, error)
}

type AccountsApp struct {
	AccountsRepo repo.AccountsRepo
	Cfg          config.Config
}

func New(accounts AccountsApp) AccountsApp {
	return accounts
}

func (i AccountsApp) UserRegister(ctx echo.Context, param request.UserRegisterReq) error {
	_, err := i.AccountsRepo.GetUserByUsername(ctx, param.Username)
	if err == nil {
		return errs.ErrEmailUsed
	}
	if !errors.Is(err, errs.ErrInstanceNotFound) {
		return err
	}

	cryptPass, err := bcrypt.GenerateFromPassword([]byte(param.Password), i.Cfg.IntBycrptPassword)
	if err != nil {
		return err
	}

	usersubject := []*models.UsersSubjects{}
	for _, v := range param.SubjectIds {
		usersubject = append(usersubject, &models.UsersSubjects{
			SubjectId: v,
		})
	}
	_, err = i.AccountsRepo.CreateUser(ctx, models.User{
		Id:                        0,
		Type:                      1,    // type 1 user not admin
		IsDocumentAnnotator:       true, // default true ?
		IsQuestionAnnotator:       true, // default true ?
		Username:                  param.Username,
		Contact:                   param.Contact,
		Age:                       param.Age,
		NumberOfDocumentAdded:     0,
		NumberOfQuestionAnnotated: 0,
		Status:                    "REGISTERED",
		Password:                  string(cryptPass),
		CreatedAt:                 time.Now(),
	}, usersubject)
	if err != nil {
		return err
	}

	return nil
}

func (i AccountsApp) AdminRegister(ctx echo.Context, param request.UserRegisterReq) error {
	_, err := i.AccountsRepo.GetUserByUsername(ctx, param.Username)
	if err == nil {
		return errs.ErrEmailUsed
	}
	if !errors.Is(err, errs.ErrInstanceNotFound) {
		return err
	}

	cryptPass, err := bcrypt.GenerateFromPassword([]byte(param.Password), i.Cfg.IntBycrptPassword)
	if err != nil {
		return err
	}

	usersubject := []*models.UsersSubjects{}
	for _, v := range param.SubjectIds {
		usersubject = append(usersubject, &models.UsersSubjects{
			SubjectId: v,
		})
	}

	_, err = i.AccountsRepo.CreateUser(ctx, models.User{
		Id:                        0,
		Type:                      2,    // type 2 admin
		IsDocumentAnnotator:       true, // default true ?
		IsQuestionAnnotator:       true, // default true ?
		Username:                  param.Username,
		Contact:                   param.Contact,
		Age:                       param.Age,
		NumberOfDocumentAdded:     0,
		NumberOfQuestionAnnotated: 0,
		Status:                    "ACTIVED",
		Password:                  string(cryptPass),
		CreatedAt:                 time.Now(),
	}, usersubject)
	if err != nil {
		return err
	}
	return nil
}

func (i AccountsApp) UserLogin(ctx echo.Context, param request.UserLoginReq) (*response.UserLoginRes, error) {
	user, err := i.AccountsRepo.GetUserByUsername(ctx, param.Username)
	if err != nil {
		return nil, err
	}

	if user.Status != "ACTIVED" {
		return nil, errs.ErrActivedAccount
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return nil, err
	}

	token, err := jwt.EncodeToken(user, i.Cfg.JWTTokenSecret)
	if err != nil {
		return nil, err
	}
	return &response.UserLoginRes{
		AccessToken: token,
	}, nil
}

func (i AccountsApp) UserProfile(ctx echo.Context) (*response.ProfileRes, error) {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := i.AccountsRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	document, err := i.AccountsRepo.GetDocumentsById(ctx, user.CurrentDocumentID)
	if err != nil {
		return nil, err
	}

	// construct response
	resp := &response.ProfileRes{
		User: response.UserRes{
			Id:                        user.Id,
			Type:                      user.Type,
			IsDocumentAnnotator:       user.IsDocumentAnnotator,
			IsQuestionAnnotator:       user.IsQuestionAnnotator,
			Username:                  user.Username,
			Contact:                   user.Contact,
			Age:                       user.Age,
			NumberOfDocumentAdded:     user.NumberOfDocumentAdded,
			NumberOfQuestionAnnotated: user.NumberOfQuestionAnnotated,
			Status:                    user.Status,
			CurrentDocumentID:         user.CurrentDocumentID,
		},
		Document: document,
	}

	return resp, nil
}

func (i AccountsApp) UpdateStatusUsers(ctx echo.Context, userID int64, status string) error {
	err := i.AccountsRepo.UpdateStatusUsers(ctx, userID, status)
	if err != nil {
		return err
	}
	return nil
}

func (i AccountsApp) GetAllUserNonAdmin(ctx echo.Context, pageNumber int64) (response.Pagination, error) {
	users, err := i.AccountsRepo.GetAllUserNonAdmin(ctx, pageNumber)
	if err != nil {
		return users, err
	}
	return users, nil
}
