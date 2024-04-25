package controllers

import (
	// "database/sql"
	"time"

	// "log"
	// "encoding/json"
	// "fmt"
	// "path/filepath"
	"backend/server"
	"backend/forms"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func SellProduct(c *gin.Context)  {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	employee_id := c.Param("employee_id")
	var products []forms.SellProduct 
	if err := c.BindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentTime := time.Now()
	
	result, err := db.Exec("INSERT INTO sales_list (employeeid,total_price,date) VALUES(?,?,?)",employee_id,0,currentTime)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	var total float32
	var total_cost float32
	for _, product := range products {
		total+= float32(product.Amount) * product.Price
		total_cost += float32(product.Amount) * product.Price_Cost
        _, err = db.Exec("INSERT INTO sales_record (salesid,idproduct,amount,product_price) VALUES(?,?,?,?)",lastInsertID,product.IDProduct,product.Amount,product.Price)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
		_, err = db.Exec("UPDATE productpa SET amount = amount - ? WHERE idproduct = ?",product.Amount,product.IDProduct)
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		}
		
    }
	_, err = db.Exec("UPDATE  sales_list SET total_price = ? WHERE salesid = ? ",total,lastInsertID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

	date := currentTime.Format("2006-01-02")
	var totalprofit = total - total_cost
	_, err = db.Exec("INSERT INTO profit_loss (date, profit, income) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE date = ?, profit = profit + ?,income = income + ?",date,totalprofit,total,date,totalprofit,total)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
	c.JSON(http.StatusOK, gin.H{"message": "sell successfully"})
}


