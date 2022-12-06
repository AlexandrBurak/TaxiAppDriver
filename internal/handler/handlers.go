package handlers

import (
	"TaxiAppDriver/internal/model"
	"TaxiAppDriver/internal/service"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{service: service}
}

// @Summary Register
// @Tags auth
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body model.Login true "account info"
// @Router /register [post]
// @Success 200 {string}
func (h *Handler) Register(c *gin.Context) {
	var user model.Driver
	ctx := c.Request.Context()
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleError(err, c)
		return
	}
	err = h.service.Save(ctx, user)
	if err != nil {
		h.handleError(err, c)
		return
	}
	err = service.SignIn(model.Login{Phone: user.Phone, Password: user.Password}, c)
	if err != nil {
		h.handleError(err, c)
		return
	}
}

// @Summary Login
// @Tags auth
// @Description login user
// @ID login-user
// @Accept json
// @Produce json
// @Param input body model.Login true "account info"
// @Router /login [post]
// @Success 200 {string}
func (h *Handler) Login(c *gin.Context) {
	var user model.Login
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleError(err, c)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "phone",
		Value:   user.Phone,
		Expires: time.Now().Add(time.Hour * 2),
	})
	c.Set("phone", user.Phone)
	ctx := c.Request.Context()
	err = h.service.Authorize(ctx, user)
	if err != nil {
		h.handleError(err, c)
		return
	}
	err = service.SignIn(user, c)
	if err != nil {
		h.handleError(err, c)
		return
	}
	c.JSON(http.StatusOK, user.Phone)
}

// @Summary Logout
// @Tags auth
// @Description logout user
// @ID logout-user
// @Accept json
// @Produce json
// @Router /logout [get]
// @Success 200  {string}
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
}

func (h *Handler) handleError(err error, c *gin.Context) {
	switch {
	case errors.Is(err, service.ErrWrongPassword):
		c.JSON(http.StatusBadRequest, "Wrong password")
	case errors.Is(err, service.ErrUserAlreadyExists):
		c.JSON(http.StatusBadRequest, "User already exists")
	case errors.Is(err, service.ErrUserNotFound):
		c.JSON(http.StatusBadRequest, "User not found")
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (h *Handler) SetDriverStatus(c *gin.Context) {
	phone, err := c.Cookie("phone")
	if err != nil {
		h.handleError(err, c)
		return
	}
	err = h.service.SetDriverStatus(c.Request.Context(), phone)
	if err != nil {
		h.handleError(err, c)
		return
	}
}
