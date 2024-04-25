package controllers

import (
	// "database/sql"
	"time"

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

func GetProducts(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM  product AS D1 join productpa AS D2 ON D1.idproduct = D2.idproduct")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	products := []forms.Product{}
	for rows.Next() {
		var product forms.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Code, &product.Desc, &product.Image, &product.Price_Cost, &product.Price, &product.ID, &product.Amount, &product.LowStock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	idproduct := c.Param("idproduct")
	rows, err := db.Query("SELECT * FROM  product AS D1 join productpa AS D2 ON D1.idproduct = D2.idproduct WHERE  D1.idproduct =?", idproduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	products := []forms.Product{}
	for rows.Next() {
		var product forms.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Code, &product.Desc, &product.Image, &product.ID, &product.Price_Cost, &product.Price, &product.Amount, &product.LowStock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var product forms.CreateProduct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO product (name, codeproduct, detail, image, cost_price, price) VALUES (?, ?, ?, ?, ?, ?)", product.Name, product.Code, product.Desc, product.Image, product.Price_Cost, product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("INSERT INTO productpa (idproduct, amount,lowstock) VALUES (?, ?,?)", lastInsertID, 0, product.LowStock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "id": lastInsertID})
}
func DeleteProduct(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	idproduct := c.Param("idproduct")

	result, err := db.Exec("DELETE FROM product WHERE idproduct = ?", idproduct)
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

	result, err = db.Exec("DELETE FROM productpa WHERE idproduct = ?", idproduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
func UpdateProduct(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var product forms.UpdateProduct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE product SET name = ?, codeproduct = ?, detail = ?, image = ?, cost_price = ?, price = ?  WHERE idproduct = ?", product.Name, product.Code, product.Desc, product.Image, product.Price_Cost, product.Price, product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// if rowsAffected == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	// 	return
	// }

	_, err = db.Exec("UPDATE productpa SET lowstock = ?  WHERE idproduct = ?", product.LowStock, product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// rowsAffected, err = result.RowsAffected()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// if rowsAffected == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func AddStockProduct(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var product forms.AddProduct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := db.Query("SELECT product.cost_price ,productpa.amount  FROM  product join productpa on product.idproduct = productpa.idproduct  WHERE  product.idproduct =?", product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	priceandamount := forms.Price_Amount{}
	for rows.Next() {
		err := rows.Scan(&priceandamount.Price, &priceandamount.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// println(((float32(priceandamount.Amount)*priceandamount.Price)+(float32(product.Amount)*product.Cost))/(float32(priceandamount.Amount+product.Amount)))
	// defer rows.Close()
	result, err := db.Exec("UPDATE productpa SET amount = amount + ? WHERE idproduct = ?", product.Amount, product.ID)
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
	_, err = db.Exec("UPDATE product SET cost_price = ? WHERE idproduct = ?", ((float32(priceandamount.Amount)*priceandamount.Price)+(float32(product.Amount)*product.Cost))/(float32(priceandamount.Amount+product.Amount)), product.ID)
	currentTime := time.Now()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("INSERT INTO addstock (datetime, idproduct, cost_price, amount, status) VALUES (?, ?, ?, ?, ?)", currentTime, product.ID, product.Cost, product.Amount, "ADD")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "add updated successfully"})
}

func GetAddRecordProducts(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("select addstock.*,product.name From addstock join product on addstock.idproduct = product.idproduct ORDER BY datetime DESC ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	lists := []forms.RecodeProduct{}
	for rows.Next() {
		var list forms.RecodeProduct
		err := rows.Scan(&list.Date, &list.IDProduct, &list.Cost, &list.Amount, &list.Status,&list.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}		
		lists = append(lists, list)
	}


	c.JSON(http.StatusOK, lists)
}

func EditRecordProducts(c *gin.Context) {
	db, err := server.CreateConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var product forms.EditRecodeProduct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := db.Query("SELECT product.cost_price ,productpa.amount  FROM  product join productpa on product.idproduct = productpa.idproduct  WHERE  product.idproduct =?", product.IDProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	priceandamount := forms.Price_Amount{}
	for rows.Next() {
		err := rows.Scan(&priceandamount.Price, &priceandamount.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	defer rows.Close()
	// currentTime := time.Now()
	_, err = db.Query("UPDATE product SET cost_price = ? WHERE idproduct = ?",(((float32(priceandamount.Amount)*priceandamount.Price)-(float32(product.Amount)*product.Cost))/float32(priceandamount.Amount-product.Amount)),product.IDProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("UPDATE productpa SET amount = amount - ? WHERE idproduct = ?", product.Amount, product.IDProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("UPDATE addstock set status = 'Edit' WHERE datetime = ?", product.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "edit updated successfully"})
}
