package handler

import (
	"net/http"

	"github.com/emersonv25/gowallet/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/wallet
// @Summary Get Wallet
// @Description Get wallet by ID
// @Tags Wallet
// @Accept json
// @Produce json
// @Param id query string true "Wallet ID"
// @Success 200 {object} CreateWalletResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /wallet [get]
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

// @Summary Create Wallet
// @Description Create a new wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body WalletRequest true "Wallet data"
// @Success 200 {object} CreateWalletResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /wallet [post]
func CreateWalletHandler(ctx *gin.Context) {
	request := WalletRequest{}
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
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	request := WalletRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	wallet := schemas.Wallet{}
	if err := db.First(&wallet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Wallet not found")
		return
	}

	wallet.Name = request.Name
	wallet.Document = request.Document
	wallet.Email = request.Email
	if err := db.Save(&wallet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to update wallet")
		return
	}
	response := schemas.NewWalletResponse(wallet)
	sendSuccess(ctx, response)
}

// @Summary Delete Wallet
// @Description Delete wallet by ID
// @Tags Wallet
// @Accept json
// @Produce json
// @Param id query string true "Wallet ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /wallet [delete]
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
