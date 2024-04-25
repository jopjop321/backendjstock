package main

import (
	// "database/sql"
	// "time"

	// "log"
	// "encoding/json"
	// "fmt"
	// "path/filepath"
	// "backend/forms"
	"backend/controllers"
	// "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Define a struct to represent your data model



func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.GET("/products", controllers.GetProducts)
	router.GET("/product/:idproduct", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.DELETE("/products/:idproduct", controllers.DeleteProduct)
	router.PUT("/product/edit", controllers.UpdateProduct)
	router.PUT("/product/addstock",controllers.AddStockProduct)
	router.POST("/product/sell/:employee_id",controllers.SellProduct)
	router.POST("/register",controllers.PostRegister)
	router.POST("/login",controllers.Postlogin)
	router.GET("/profile/:employeeid",controllers.GetProfile)
	router.GET("/profile/all",controllers.GetallProfile)
	router.PUT("/profile/ownerupdate/:employeeid" , controllers.UpdateOwner)
	router.PUT("/profile/updateprofile" , controllers.UpdateProfile)
	router.PUT("/profile/updatestatus", controllers.UpdateStatus)
	router.GET("/profit" , controllers.GetallProfit)
	router.GET("/histony" , controllers.Getahistony)
	router.GET("/profit/ip" , controllers.Getincomeprofit)
	router.GET("/product/addstock/recode" , controllers.GetAddRecordProducts)
	router.PUT("/product/addstock/edit" , controllers.EditRecordProducts)
	router.Run(":8080")
}
