package handler

import (
	"net/http"

	"github.com/emersonv25/gowallet/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"message":    message,
		"statusCode": statusCode,
	})
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type CreateWalletResponse struct {
	Data schemas.WalletResponse `json:"data"`
}

type CreateTransactionResponse struct {
	Data schemas.TransactionResponse `json:"data"`
}

type ListTransactionsResponse struct {
	Data []schemas.TransactionResponse `json:"data"`
}
