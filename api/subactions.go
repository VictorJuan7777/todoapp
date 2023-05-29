package api

import (
	"fmt"
	"net/http"
	"strconv"
	db "todoapp/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateSubActionReq struct {
	Action_id string `json:"action_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
}

func (server *server) createSubAction(ctx *gin.Context) {
	var req CreateSubActionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, _ := strconv.Atoi(req.Action_id)
	arg := db.CreateSubActionParams{
		ActionsID: int64(id),
		Title:     req.Title,
	}
	action, err := server.store.CreateSubAction(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			//log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, action)
}

type CompletedSubActionReq struct {
	Completed *bool  `json:"completed" binding:"required"`
	ID        string `json:"id" binding:"required"`
	Title     string `json:"title" binding:"required"`
}

func (server *server) completedSubAction(ctx *gin.Context) {
	var req CompletedSubActionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, _ := strconv.Atoi(req.ID)
	arg := db.UpdateSubActionParams{
		ID:        int64(id),
		Completed: *req.Completed,
		Title:     req.Title,
	}
	action, err := server.store.UpdateSubAction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, action)

}

type DeleteSubActionReq struct {
	ID string `json:"id" binding:"required"`
}

func (server *server) deletSubedAction(ctx *gin.Context) {
	var req DeleteSubActionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, _ := strconv.Atoi(req.ID)
	err := server.store.DeleteSubAction(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, errorResponse(fmt.Errorf("Deleted Susseccfully")))
}
