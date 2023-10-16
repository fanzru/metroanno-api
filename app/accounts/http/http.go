package http

import (
	"errors"
	"fmt"
	errs "metroanno-api/app/accounts/domain/errors"
	"metroanno-api/app/accounts/domain/request"
	accountsapp "metroanno-api/app/accounts/usecase"
	"metroanno-api/infrastructure/config"
	"metroanno-api/pkg/response"
	"strconv"

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
		Username:   userRegisterReq.Username,
		Password:   userRegisterReq.Password,
		Contact:    userRegisterReq.Contact,
		Age:        userRegisterReq.Age,
		SubjectIds: userRegisterReq.SubjectIds,
	})
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessCreated(ctx, nil)
}

func (h AccountHandler) RegisterUserV2(ctx echo.Context) error {
	userRegisterReq := &request.UserRegisterReqV2{}

	err := ctx.Bind(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(userRegisterReq)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.AccountsApp.UserRegister(ctx, request.UserRegisterReq{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
		Contact:  userRegisterReq.Contact,
		Age:      userRegisterReq.Age,
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
		Username:   userRegisterReq.Username,
		Password:   userRegisterReq.Password,
		Contact:    userRegisterReq.Contact,
		Age:        userRegisterReq.Age,
		SubjectIds: userRegisterReq.SubjectIds,
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

func (h *AccountHandler) UserProfile(ctx echo.Context) error {
	r, err := h.AccountsApp.UserProfile(ctx)

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, r)
}

func (h AccountHandler) UpdateStatusUsers(ctx echo.Context) error {
	req := &request.UpdateStatusUserReq{}

	err := ctx.Bind(req)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	err = validator.New().Struct(req)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	if !req.ValidateStatus() {
		return response.ResponseErrorBadRequest(ctx, errors.New("invalid status"))
	}

	err = h.AccountsApp.UpdateStatusUsers(ctx, req.UserID, req.Status)

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, nil)
}
func (h AccountHandler) GetAllUserNonAdmin(ctx echo.Context) error {
	PN := ctx.QueryParam("pageNumber")

	if PN == "" {
		return response.ResponseErrorBadRequest(ctx, errs.ErrPageNumber)
	}

	pageNumber, err := strconv.ParseInt(fmt.Sprintf("%v", PN), 10, 64)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	r, err := h.AccountsApp.GetAllUserNonAdmin(ctx, pageNumber)

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, r)
}
