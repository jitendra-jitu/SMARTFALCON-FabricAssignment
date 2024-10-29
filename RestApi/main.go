package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
)

func getGateway(userId string) *gateway.Gateway {
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet in getGateway: %v", err)
	}

	if !wallet.Exists(userId) {
		log.Fatalf("Wallet does not contain '%s' identity in getGateway", userId)
	}

	gateway, err := gateway.Connect(gateway.WithConfig("connection.yaml"), gateway.WithIdentity(wallet, userId))
	if err != nil {
		log.Fatalf("Failed to connect to gateway in getGateway: %v", err)
	}

	return gateway
}

func createAsset(c *gin.Context) {
	userId := c.Query("userId") // Get the user ID from query params

	dealerId := c.PostForm("dealerId")
	msisdn := c.PostForm("msisdn")
	mpin := c.PostForm("mpin")
	balance := c.PostForm("balance")
	status := c.PostForm("status")
	transAmount := c.PostForm("transAmount")
	transType := c.PostForm("transType")
	remarks := c.PostForm("remarks")

	if dealerId == "" || msisdn == "" || mpin == "" || balance == "" {
		c.JSON(400, gin.H{"error": "Required fields are missing"})
		return
	}

	gateway := getGateway(userId)
	defer gateway.Close()

	contract := gateway.GetNetwork("mychannel").GetContract("asset-transfer-basic")

	result, err := contract.SubmitTransaction("CreateAsset", dealerId, msisdn, mpin, balance, status, transAmount, transType, remarks)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": string(result)})
}

func queryAsset(c *gin.Context) {
	userId := c.Query("userId") // Get the user ID from query params

	dealerId := c.Param("dealerId")

	gateway := getGateway(userId)
	defer gateway.Close()

	contract := gateway.GetNetwork("mychannel").GetContract("asset-transfer-basic")

	result, err := contract.EvaluateTransaction("QueryAsset", dealerId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": string(result)})
}

func updateAsset(c *gin.Context) {
	userId := c.Query("userId") // Get the user ID from query params

	dealerId := c.Param("dealerId")
	balance := c.PostForm("balance")
	status := c.PostForm("status")

	if dealerId == "" || balance == "" {
		c.JSON(400, gin.H{"error": "Required fields are missing"})
		return
	}

	gateway := getGateway(userId)
	defer gateway.Close()

	contract := gateway.GetNetwork("mychannel").GetContract("asset-transfer-basic")

	result, err := contract.SubmitTransaction("UpdateAsset", dealerId, balance, status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": string(result)})
}

func main() {
	r := gin.Default()

	r.POST("/assets", createAsset)
	r.GET("/assets/:dealerId", queryAsset)
	r.PUT("/assets/:dealerId", updateAsset)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
