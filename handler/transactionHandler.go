package handler

import (
	"net/http"

	"github.com/emersonv25/gowallet/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Get Transactions
// @Description Get all transactions for a wallet
// @Tags Transaction
// @Accept json
// @Produce json
// @Param walletId query string true "Wallet ID"
// @Success 200 {object} ListTransactionsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [get]
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

// @Summary Create Transaction
// @Description Create a new transaction (credit or debit)
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body CreateTransactionRequest true "Transaction data"
// @Success 200 {object} CreateTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [post]
func CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("requestBody", "json").Error())
		return
	}
	if err := req.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	wallet := schemas.Wallet{}
	if err := db.First(&wallet, req.WalletID).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Wallet not found")
		return
	}

	switch req.Type {
	case schemas.TransactionTypeDebit:
		if wallet.Balance < req.Amount {
			sendError(ctx, http.StatusBadRequest, "Insufficient balance for debit transaction")
			return
		}
		wallet.Balance -= req.Amount
	case schemas.TransactionTypeCredit:
		wallet.Balance += req.Amount

	}
	if err := db.Save(&wallet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to update wallet balance")
		return
	}

	transaction := schemas.NewTransaction(uint(req.WalletID), req.Amount, req.Type, req.Description)
	if err := db.Create(&transaction).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to create transaction")
		return
	}

	sendSuccess(ctx, schemas.NewTransactionResponse(transaction))
}
