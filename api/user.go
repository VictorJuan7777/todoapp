package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "todoapp/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (server *server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
	}
	users, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (server *server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	users, err := server.store.GetUserID(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err = CheckPassword(req.Password, users.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	calluser := CallbackUser(users)
	accessToken, err, _ := server.tokenMaker.CreateToken(users.Username, CookieTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        *calluser,
	}
	ctx.SetCookie("JWT", accessToken, 0, "/", "", false, true)
	ctx.SetCookie("Username", rsp.User.Username, 0, "/", "", false, false)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserResponse struct {
	User        callUser `json:"user"`
	AccessToken string   `json:"access_token"`
}

func (server *server) Logout(ctx *gin.Context) {
	ctx.SetCookie("JWT", "", -1, "/", "", false, true)
	ctx.SetCookie("Username", "", -1, "/", "", false, false)
	err := fmt.Errorf("successful logout")
	ctx.JSON(http.StatusOK, errorResponse(err))
	return
}

func (server *server) RefreshCookie(ctx *gin.Context) {
	cookie, err := ctx.Cookie("JWT")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	payload, err := server.tokenMaker.VerifyToken(cookie)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	users, err := server.store.GetUserID(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if users.Username != payload.Username {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("Unauthorization Token")))
		return
	}

	accessToken, err, _ := server.tokenMaker.CreateToken(payload.Username, CookieTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.SetCookie("JWT", accessToken, 0, "/", "", false, true)
	ctx.SetCookie("Username", payload.Username, 0, "/", "", false, false)
	ctx.JSON(http.StatusOK, "Refresh cookie sucessful")

}
