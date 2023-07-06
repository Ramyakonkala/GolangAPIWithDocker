package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ProcessReceipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

var receipts = []ProcessReceipt{
	{ID: "7fb1377b-b223-49d9-a31a-5a02701dd310", Retailer: "Walgreens", PurchaseDate: "2022-01-02", PurchaseTime: "08:13",
		Items: []Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "Dasani", Price: "1.40"}}, Total: "2.65"},
}

func getReceiptsById(id string) (*ProcessReceipt, error) {
	for i, b := range receipts {
		if b.ID == id {
			return &receipts[i], nil
		}
	}

	return nil, errors.New("Receipt not found")
}

func calculateReceiptPoints(recepit *ProcessReceipt) int {

	points := 0
	//retailer Name char len
	charCount := 0
	var alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	name := string(recepit.Retailer)
	for i := 0; i < len(name); i++ {
		if alphanumericRegex.MatchString(string(name[i])) {
			charCount++
		}
	}

	points = points + charCount

	//receipt total
	val, err := strconv.ParseFloat(recepit.Total, 64)

	if err == nil {
		if math.Mod(val, 0.25) == 0 {
			points = points + 25
		}
		if val == math.Trunc(val) {
			points = points + 50
		}
	}

	//receipt items
	points = points + (int(len(recepit.Items)/2) * 5)

	//receipt items trimmed description
	for _, item := range recepit.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			priceVal, _ := strconv.ParseFloat(item.Price, 64)

			points = points + int(math.Ceil(priceVal*0.2))
		}
	}

	//receipt date yyyy-mm-dd
	datetime, _ := time.Parse("2006-01-02", recepit.PurchaseDate)
	if datetime.Day()%2 != 0 {
		points = points + 6
	}

	//receipt time
	timeHH, _ := strconv.Atoi(strings.Split(recepit.PurchaseTime, ":")[0])

	if timeHH >= 14 && timeHH < 16 {
		points = points + 10
	}

	return points
}

func getReceiptPoints(c *gin.Context) {
	id := c.Param("id")
	receipt, err := getReceiptsById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Receipt not found for given ID."})
		return
	}

	// calculate receipt points
	points := calculateReceiptPoints(receipt)
	c.IndentedJSON(http.StatusOK, gin.H{"points": points})
}

func getProcessReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func addProcessReceipts(c *gin.Context) {
	var newReceipt ProcessReceipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt.ID = uuid.New().String()
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

func main() {
	r := gin.Default()
	r.GET("/receipts/process", getProcessReceipts)
	r.GET("/receipts/:id/points", getReceiptPoints)
	r.POST("/receipts/process", addProcessReceipts)
	r.Run()
}
