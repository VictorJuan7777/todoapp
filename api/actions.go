package api

import (
	"fmt"
	"net/http"
	"strconv"
	db "todoapp/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreateActionReq struct {
	Username string `json:"username" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

func (server *server) createAction(ctx *gin.Context) {
	var req CreateActionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateActionParams{
		Username: req.Username,
		Title:    req.Title,
	}
	action, err := server.store.CreateAction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, action)
}

type CompletedActionReq struct {
	Completed *bool  `json:"completed" binding:"required"`
	ID        string `json:"id" binding:"required"`
	Title     string `json:"title" binding:"required"`
}

func (server *server) completedAction(ctx *gin.Context) {
	var req CompletedActionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, _ := strconv.Atoi(req.ID)
	arg := db.UpdateActionParams{
		ID:        int64(id),
		Completed: *req.Completed,
		Title:     req.Title,
	}
	action, err := server.store.UpdateAction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, action)

}

type DeleteActionReq struct {
	ID string `json:"id" binding:"required"`
}

func (server *server) deletedAction(ctx *gin.Context) {
	var req DeleteActionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, _ := strconv.Atoi(req.ID)
	err := server.store.DeleteAllSubAction(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteAction(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, errorResponse(fmt.Errorf("Deleted Susseccfully")))
}

type listActionReq struct {
	Username string `json:"username" binding:"required"`
}

type allList struct {
	Action    db.Actions      `json:"action"`
	Subaction []db.Subactions `json:"subaction"`
}

func (server *server) listAction(ctx *gin.Context) {
	var req listActionReq
	all := []allList{}
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

	req.Username = payload.Username
	list, err := server.store.ListAction(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for i := range list {
		list2, err := server.store.ListSubAction(ctx, int64(list[i].ID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		var res allList
		res.Action = list[i]
		res.Subaction = list2
		all = append(all, res)
	}
	ctx.JSON(http.StatusOK, all)

}
