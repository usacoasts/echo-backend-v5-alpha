package handler

import (
	"app/domain/model"
	"app/domain/service"
	"app/usecase"
	"context"
	"fmt"
	"github.com/labstack/echo/v5"
	"net/http"
)

// UserHandler interface
type UserHandler interface {
	CreateUser(c echo.Context) error
	Login(c echo.Context) error
	Refresh(c echo.Context) error
	Logout(c echo.Context) error
}

type ApiUserTokenHandler interface {
	CreateToken(Id int) error
	InsertToken(c echo.Context) error
}

type userHandler struct {
	UserUseCase         usecase.UserUseCase
	ApiUserTokenUseCase usecase.ApiUserTokenUseCase
	AuthUserService     service.AuthUserService
}

type AuthUserService struct {
	AuthUserService service.AuthUserService
}

// NewUserHandler UserHandlerを取得します.
func NewUserHandler(UserUseCase usecase.UserUseCase, ApiUserTokenUseCase usecase.ApiUserTokenUseCase, AuthUserService service.AuthUserService) UserHandler {
	return &userHandler{UserUseCase, ApiUserTokenUseCase, AuthUserService}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandler) Login(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.Login(ctx, user)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Certification Failed")
	}

	token, err := h.ApiUserTokenUseCase.CreateToken(user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User Token Cannot Create")
	}

	apiUserToken := &model.ApiUserToken{
		User_ID: user.ID, Token: token,
	}

	fmt.Println(apiUserToken)

	_, err = h.ApiUserTokenUseCase.InsertToken(apiUserToken)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h *userHandler) Logout(c echo.Context) error {
	user, err := h.AuthUserService.AuthUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist.")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) Refresh(c echo.Context) error {

	user, err := h.AuthUserService.AuthUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist.")
	}

	return c.JSON(http.StatusOK, user)
}
