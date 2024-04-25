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

func GetProfile(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()
	employee_id := c.Param("employeeid")

	data, err := db.Query("SELECT * FROM data_employee WHERE employeeid = ?",employee_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer data.Close()

	var profile  forms.Employee_profile
	for data.Next() {
		err := data.Scan(&profile.Employee_ID, &profile.Name, &profile.Last_name, &profile.Nickname, &profile.Position)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, profile)
}

func GetallProfile(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM data_employee")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer data.Close()

	profiles := []forms.Employee_profile{}
	for data.Next() {
		var profile forms.Employee_profile
		err := data.Scan(&profile.Employee_ID, &profile.Name, &profile.Last_name, &profile.Nickname, &profile.Position)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		profiles = append(profiles, profile)
	}
	c.JSON(http.StatusOK, profiles)
}

func UpdateOwner(c *gin.Context)  {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	employee_id := c.Param("employeeid")
	_, err = db.Exec("UPDATE data_employee SET  position = ? WHERE position = ?","Staff" , "Owner")
	result, err := db.Exec("UPDATE data_employee SET  position = ? WHERE employeeid = ?","Owner" , employee_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Owner updated successfully"})
}

func UpdateProfile(c *gin.Context)  {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var newprofile forms.Employee_information 
	if err := c.BindJSON(&newprofile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("UPDATE data_employee SET  name = ? , last_name = ? , nickname = ? WHERE employeeid = ?",newprofile.Name, newprofile.Last_name , newprofile.Nickname ,newprofile.Employee_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "profile updated successfully"})
}

func UpdateStatus(c *gin.Context)  {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var newstatus forms.Update_Status 
	if err := c.BindJSON(&newstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("UPDATE data_employee SET  position = ? WHERE employeeid = ?" ,newstatus.Position,newstatus.Employee_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}
