package handler

import (
	"net/http"

	"github.com/emersonv25/gowallet/schemas"
	"github.com/gin-gonic/gin"
)

func GetTransactions(ctx *gin.Context) {
	walletId := ctx.Query("walletId")
	if walletId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("walletId", "queryParameter").Error())
		return
	}
	var transactions []schemas.Transaction
	if err := db.Where("wallet_id = ?", walletId).Find(&transactions).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to retrieve transactions")
		return
	}

	var response []schemas.TransactionResponse
	for _, transaction := range transactions {
		response = append(response, schemas.NewTransactionResponse(transaction))
	}
	sendSuccess(ctx, response)
}
