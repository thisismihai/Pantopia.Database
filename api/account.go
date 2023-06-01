package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/mbaxamb3/pantopia/db/sqlc"
	"github.com/mbaxamb3/pantopia/token"
)

type createAccountRequest struct {
	CompanyName    string `json:"company_name" binding:"required"`
	Industry       string `json:"industry" binding:"required,oneof=Education Construction Healthcare Finance"`
	TargetIndustry string `json:"target_industry" binding:"required,oneof=Education Construction Healthcare Finance"`
	TargetPosition string `json:"target_position" binding:"required,oneof=Management CEO CTO Finance"`
	TargetSize     int64  `json:"target_size" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:          authPayload.Username,
		CompanyName:    req.CompanyName,
		Industry:       req.Industry,
		TargetIndustry: req.TargetIndustry,
		TargetPosition: req.TargetPosition,
		TargetSize:     req.TargetSize,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Print(pqErr.Code.Name()) //
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, account)

}
