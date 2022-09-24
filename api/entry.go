package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/xdot2012/simple-bank/db/sqlc"
)

type listEntriesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listEntries(ctx *gin.Context) {
	var req listEntriesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entries)
}

type createEntryRequest struct {
	AccountID int64 `json:"account_id" binding:"required,min=1"`
	Amount    int64 `json:"amount" binding:"required"`
}

func (server *Server) createEntry(ctx *gin.Context) {
	var req createEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type getEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEntry(ctx *gin.Context) {
	var req getEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type updateEntryURIRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateEntryBodyRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}

func (server *Server) updateEntry(ctx *gin.Context) {
	var req updateEntryURIRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateEntryBodyRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateEntryParams{
		ID:     req.ID,
		Amount: reqBody.Amount,
	}

	entry, err := server.store.UpdateEntry(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type deleteEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteEntry(ctx *gin.Context) {
	var req deleteEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
