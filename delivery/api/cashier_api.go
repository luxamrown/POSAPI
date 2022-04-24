package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mohamadelabror.com/posapi/model"
	"mohamadelabror.com/posapi/usecase"
)

type CashierApi struct {
	showCashier   usecase.ShowAllCashierUseCase
	getCashier    usecase.GetCashierDetailUseCase
	createCashier usecase.CreateCashierUseCase
	updateCashier usecase.UpdateCashierUseCase
	deleteCashier usecase.DeleteCashierUseCase
}

func (a *CashierApi) ShowAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := c.Query("limit")
		skip := c.Query("skip")
		cashiers, err := a.showCashier.ShowAll(limit, skip)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cashiers)
	}
}

func (a *CashierApi) GetDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		cashierId := c.Param("cashierId")
		cashier, err := a.getCashier.GetDetail(cashierId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"succes":  true,
			"message": "success",
			"data":    cashier,
		})
	}
}

func (a *CashierApi) CreateCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cashier model.Cashier
		err := c.ShouldBindJSON(&cashier)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ErrorJSON"})
			return
		}
		selectedCashier, err := a.createCashier.Insert(cashier.Name, cashier.Passcode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    selectedCashier,
		})
	}
}

func (a *CashierApi) UpdateCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cashier model.Cashier
		cashierId := c.Param("cashierId")
		err := c.ShouldBindJSON(&cashier)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ErrorJSON"})
			return
		}
		err = a.updateCashier.UpdateCashier(cashierId, cashier.Name, cashier.Passcode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "succes",
		})
	}
}

func (a *CashierApi) DeleteCashier() gin.HandlerFunc {
	return func(c *gin.Context) {
		cashierId := c.Param("cashierId")
		err := a.deleteCashier.DeleteCashier(cashierId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
		})
	}

}

func NewCashierApi(cashierRoute *gin.RouterGroup, showCashier usecase.ShowAllCashierUseCase, getCashier usecase.GetCashierDetailUseCase, createCashier usecase.CreateCashierUseCase, updateCashier usecase.UpdateCashierUseCase, deleteCashier usecase.DeleteCashierUseCase) {
	api := CashierApi{
		showCashier:   showCashier,
		getCashier:    getCashier,
		createCashier: createCashier,
		updateCashier: updateCashier,
		deleteCashier: deleteCashier,
	}
	cashierRoute.GET("", api.ShowAll())
	cashierRoute.GET("/:cashierId", api.GetDetail())
	cashierRoute.POST("", api.CreateCashier())
	cashierRoute.PUT("/:cashierId", api.UpdateCashier())
	cashierRoute.DELETE("/:cashierId", api.DeleteCashier())
}
