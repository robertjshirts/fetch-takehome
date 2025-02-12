package api

import (

	"github.com/gin-gonic/gin"
	_ "github.com/robertjshirts/fetch-takehome/gen"
)

type ReceiptHandler struct {
}

func NewReceiptHandler() *ReceiptHandler {
	return &ReceiptHandler{}
}

func (ReceiptHandler) GetReceiptsIdPoints(c *gin.Context, id string) {
	c.JSON(200, gin.H{
		"message": "Get Receipts",
	})
}

func (ReceiptHandler) PostReceiptsProcess(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Receipts",
	})
}
