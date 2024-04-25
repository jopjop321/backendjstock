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
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func PostRegister(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var register forms.Register
	if err := c.BindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	check, err := db.Query("SELECT name,last_name,nickname FROM data_employee WHERE name = ? AND last_name = ? AND nickname = ?", register.Name, register.Last_name, register.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer check.Close()

	if check.Next() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ซ้ำ"})
		return
	}
	password := []byte(register.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	result, err := db.Exec("INSERT INTO data_employee (name, last_name, nickname, position) VALUES (?, ?, ?, ?)", register.Name, register.Last_name, register.Nickname, "NoStatus")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("INSERT INTO data_login (employeeid, id, password) VALUES (?, ?, ?)", lastInsertID, register.ID, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register created successfully", "id": lastInsertID})
}

func Postlogin(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var login forms.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	check, err := db.Query("SELECT password,employeeid FROM data_login WHERE id = ?", login.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer check.Close()
	if check.Next() {
		var hashedpassword string
		var employee_id int
		if err := check.Scan(&hashedpassword, &employee_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(login.Password))
		// println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "รหัสผ่านไม่ถูกต้อง"})
			return
		}
		data, err := db.Query("SELECT  employeeid,name,last_name,nickname,position FROM data_employee WHERE employeeid = ?", employee_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var data2 forms.Employee_profile
		for data.Next() {
			err = data.Scan(&data2.Employee_ID,&data2.Name,&data2.Last_name,&data2.Nickname,&data2.Position)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK,data2)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID ไม่ถูกต้อง"})
		return
	}
}
