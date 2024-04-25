package controllers

import (
	// "database/sql"
	// "time"

	// "log"
	// "encoding/json"
	// "fmt"
	// "path/filepath"
	"backend/forms"
	"backend/server"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func GetallProfit(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT date,profit FROM  profit_loss")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	profits := []forms.Profit{}
	for rows.Next() {
		var profit forms.Profit
		err := rows.Scan(&profit.Date, &profit.Profit,)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		profits = append(profits, profit)
	}

	c.JSON(http.StatusOK, profits)
}

func Getahistony(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT D3.name,D2.amount,D2.product_price FROM sales_list AS D1 join sales_record AS D2 ON D1.salesid = D2.salesid join product as D3 ON D2.idproduct = D3.idproduct ORDER BY D1.date DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	histonys := []forms.Histony{}
	for rows.Next() {
		var histony forms.Histony
		var price float32
		err := rows.Scan(&histony.Name, &histony.Amount,&price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		histony.Price_Total = price*float32(histony.Amount)
		histonys = append(histonys, histony)
	}

	c.JSON(http.StatusOK, histonys)
}

func Getincomeprofit(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("select SUM(profit) as totalprofit,SUM(income) as totalincome  from profit_loss where YEAR(date) = YEAR(CURRENT_DATE())")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	ip := forms.Profitincome{}

	for rows.Next() {
		err := rows.Scan(&ip.Profit_Year, &ip.Income_Year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	rows, err = db.Query("select SUM(profit) as totalprofit,SUM(income) as totalincome  from profit_loss where MONTH(date) = MONTH(CURRENT_DATE())")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&ip.Profit_Month, &ip.Income_Month)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	

	c.JSON(http.StatusOK, ip)
}