package handler

import (
	"net/http"

	"github.com/emersonv25/gowallet/schemas"
	"github.com/gin-gonic/gin"
)

func GetWallet(ctx *gin.Context) {
	// Implementation for creating a client
	ctx.JSON(http.StatusOK, gin.H{
		"status": "GET Wallet",
	})
}

func CreateWalletHandler(ctx *gin.Context) {
	request := CreateWalletRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	wallet := schemas.NewWallet(
		request.Name,
		request.Document,
		request.Email,
	)

	if err := db.Create(&wallet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to create wallet")
		return
	}

	response := schemas.NewWalletResponse(wallet)

	sendSuccess(ctx, response)
}
func UpdateWalletHandler(ctx *gin.Context) {
	// Implementation for updating a client
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Wallet Updated",
	})
}

func DeleteWalletHandler(ctx *gin.Context) {
	// Implementation for deleting a client
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Wallet Deleted",
	})
}
