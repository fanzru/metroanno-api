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

	_, err = i.AccountsRepo.CreateUser(ctx, models.User{
		Id:                        0,
		Type:                      1,    // type 1 user not admin
		IsDocumentAnnotator:       true, // default true ?
		IsQuestionAnnotator:       true, // default true ?
		SubjectPreference:         param.SubjectPreference,
		Username:                  param.Username,
		Contact:                   param.Contact,
		Age:                       param.Age,
		NumberOfDocumentAdded:     0,
		NumberOfQuestionAnnotated: 0,
		Status:                    "Aman",
		Password:                  string(cryptPass),
		CreatedAt:                 time.Now(),
	})
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

	_, err = i.AccountsRepo.CreateUser(ctx, models.User{
		Id:                        0,
		Type:                      2,    // type 2 admin
		IsDocumentAnnotator:       true, // default true ?
		IsQuestionAnnotator:       true, // default true ?
		SubjectPreference:         param.SubjectPreference,
		Username:                  param.Username,
		Contact:                   param.Contact,
		Age:                       param.Age,
		NumberOfDocumentAdded:     0,
		NumberOfQuestionAnnotated: 0,
		Status:                    "Aman",
		Password:                  string(cryptPass),
		CreatedAt:                 time.Now(),
	})
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
			SubjectPreference:         user.SubjectPreference,
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
