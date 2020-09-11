package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iotexproject/iotex-blockchain-iot/blockchain"
)

func addDeviceData(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	deviceTopic := ctx.Param("topic")
	if err := blockchain.ExecuteContract(deviceTopic, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, nil)
}
