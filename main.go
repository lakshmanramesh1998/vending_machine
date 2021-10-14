package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Orders struct {
	Menu           []string `json:"menu"`
	ChangeInReturn int64    `json:"changeInReturn"`
	Message        string   `json:"message"`
	Error          string    `json:"error"`
}

type Vendormacine struct {
	Currency int64 `json:"currency"`
	Items    []Items
}

type Items struct {
	Drinkname string `json:"drinkName"`
	Quantity  int64  `json:"quantity"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/vendor", func(c *gin.Context) {
		vendor := Vendormacine{}
		err := c.ShouldBindJSON(&vendor)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		validCurrency := false
		message := ""
		val, err := vendor.VendorCalculation(validCurrency, message)
		if err != nil {
			log.Println("Internal error", err)
		}
		c.JSON(http.StatusOK, val)
	})
	r.Run()
}

func (c *Vendormacine) VendorCalculation(validCurrency bool, message string) (orderPlaced Orders, err error) {

	// Checking whether the given currency is lower than 2000
	currency := fmt.Sprintf("%v", c.Currency)
	switch {
	case currency == "20":
		validCurrency = true
		order, err := ReceiveOrders(*c, message)
		if err != nil {
			fmt.Println("Something went wrong---", err)
		}
		return *order, nil

	case currency == "50":
		order, err := ReceiveOrders(*c, message)
		if err != nil {
			fmt.Println("Something went wrong---", err)
		}
		return *order, nil

	case currency == "100":
		validCurrency = true
		order, err := ReceiveOrders(*c, message)
		if err != nil {
			fmt.Println("Something went wrong---", err)
		}
		return *order, nil

	case currency == "200":
		validCurrency = true
		order, err := ReceiveOrders(*c, message)
		if err != nil {
			fmt.Println("Something went wrong---", err)
		}
		return *order, nil

	case currency == "500":
		validCurrency = true
		order, err := ReceiveOrders(*c, message)
		if err != nil {
			fmt.Println("Something went wrong---", err)
		}
		return *order, nil

	case currency == "2000":
		validCurrency = false
		order, err := InvalidOrder()
		if err != nil {
			fmt.Println("Something went wrong---", err)
		}
		return *order, nil

	default:
		validCurrency = true
	}
	return
}

func ReceiveOrders(c Vendormacine, message string) (orders *Orders, err error) {
	order := Orders{}
	cost := map[string]int{
		"Pepsi":       25,
		"Coke":        30,
		"Mirinda":     30,
		"Sevenup":     35,
		"Thumbsup":    30,
		"Mountaindew": 40,
		"Redbull":     80,
		"Sprite":      30,   
	}
	total := 0

	// finding total cost of the drink
	for _, pricelist := range c.Items {
		total += cost[pricelist.Drinkname] * int(pricelist.Quantity)
	}
	enteredCurrency := int(c.Currency)

    //  checking whether the selected item exceeds the currency's worth
	if total > enteredCurrency {
		order.Message = "Selected item exceeds the amount you entered"
		order.Error = "Value invalid"
		return &order, nil
	}
	remainingChange := enteredCurrency - total

	order.ChangeInReturn = int64(remainingChange)
	order.processOrder(c)
	return &order, nil
}

func (n *Orders) processOrder(c Vendormacine) {
	var deliverables []string
	for _, val := range c.Items {
		for i := 0; i < int(val.Quantity); i++ {
			deliverables = append(deliverables, val.Drinkname)
		}
	}
	n.Menu = deliverables
	n.Message = "Thank you for using our service"
	n.Error = ""
}

func InvalidOrder() (order *Orders, err error) {
	orderStruct := Orders{}
	orderStruct.Message = "Sorry we dont accept currency above 2000"
	orderStruct.ChangeInReturn = 2000
	orderStruct.Error = "We don't have change for 2000 notes"
	return &orderStruct, nil
}
