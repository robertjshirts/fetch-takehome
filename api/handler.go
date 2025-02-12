package api

import (
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/robertjshirts/fetch-takehome/gen"
	"github.com/robertjshirts/fetch-takehome/internal"
)

type ReceiptHandler struct {
	db map[string]int // map[receiptId]points
}

func NewReceiptHandler() *ReceiptHandler {
	return &ReceiptHandler{
		db: make(map[string]int),
	}
}

func (r *ReceiptHandler) GetReceiptsIdPoints(c *gin.Context, id string) {
	points, ok := r.db[id]
	if !ok {
		c.Status(404)
		return
	}
	c.JSON(200, gin.H{
		"points": points,
	})
}

func (r *ReceiptHandler) PostReceiptsProcess(c *gin.Context) {
	var receipt gen.Receipt
	err := c.BindJSON(&receipt)
	if err != nil {
		c.Status(400)
		return
	}
	id := uuid.New().String()
	points, err := internal.GetPoints(&receipt)
	if err != nil {
		c.Status(400)
		return
	}
	r.db[id] = points
	c.JSON(200, gin.H{
		"id": id,
	})
}
