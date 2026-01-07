package handler

import (
	"net/http"

	"github.com/emersonv25/gowallet/schemas"
	"github.com/gin-gonic/gin"
)

func GetWallet(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	wallet := schemas.Wallet{}
	if err := db.First(&wallet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Wallet not founds")
		return
	}

	response := schemas.NewWalletResponse(wallet)

	sendSuccess(ctx, response)
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
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	wallet := schemas.Wallet{}
	if err := db.First(&wallet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Wallet not found")
		return
	}

	if err := db.Delete(&wallet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to delete wallet")
		return
	}
	sendSuccess(ctx, nil)
}
