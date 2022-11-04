package http

import (
	"metroanno-api/app/accounts/domain/request"
	accountsapp "metroanno-api/app/accounts/usecase"
	"metroanno-api/infrastructure/config"
	"metroanno-api/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	AccountsApp accountsapp.AccountsApp
	Cfg         config.Config
}

func (h AccountHandler) RegisterUser(ctx echo.Context) error {
	userRegisterReq := &request.UserRegisterReq{}

	err := ctx.Bind(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.AccountsApp.UserRegister(ctx, request.UserRegisterReq{
		Username:          userRegisterReq.Username,
		Password:          userRegisterReq.Password,
		SubjectPreference: userRegisterReq.SubjectPreference,
		Contact:           userRegisterReq.Contact,
		Age:               userRegisterReq.Age,
	})
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessCreated(ctx, nil)
}

func (h AccountHandler) RegisterAdmin(ctx echo.Context) error {
	userRegisterReq := &request.UserRegisterReq{}

	err := ctx.Bind(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.AccountsApp.AdminRegister(ctx, request.UserRegisterReq{
		Username:          userRegisterReq.Username,
		Password:          userRegisterReq.Password,
		SubjectPreference: userRegisterReq.SubjectPreference,
		Contact:           userRegisterReq.Contact,
		Age:               userRegisterReq.Age,
	})
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessCreated(ctx, nil)
}

func (h *AccountHandler) Login(ctx echo.Context) error {
	userLoginReq := &request.UserLoginReq{}

	err := ctx.Bind(userLoginReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(userLoginReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	token, err := h.AccountsApp.UserLogin(ctx, request.UserLoginReq{
		Username: userLoginReq.Username,
		Password: userLoginReq.Password,
	})

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, token)
}
