package forms

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Desc         string `json:"description"`
	Image        string `json:"image"`
	Price_Cost   float32    `json:"cost_price"`
	Price		 float32    `json:"price"`
	Amount       int    `json:"amount"`
	LowStock	int `json:"low_stock"`
}

type CreateProduct struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Desc         string `json:"description"`
	Image        string `json:"image"`
	Price_Cost   float32    `json:"cost_price"`
	Price		 float32    `json:"price"`
	LowStock	int `json:"low_stock"`
}

type UpdateProduct struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Desc         string `json:"description"`
	Image        string `json:"image"`
	Price_Cost   float32    `json:"cost_price"`
	Price		 float32    `json:"price"`
	LowStock	int `json:"low_stock"`
}

type AddProduct struct {
	ID           int    `json:"id"`
	Amount       int    `json:"amount"`
	Cost		 float32	`json:"cost"`
}
type Price_Amount struct{
	Price float32 `json:"price"`
	Amount int `json:"amount"`
}

type RecodeProduct struct{
	Date string `json:"date"`
	IDProduct int `json:"idproduct"`
	Cost float32 `json:"cost"`
	Amount int `json:"amount"`
	Status string `json:"status"`
	Name string `json:"name"`

}

type EditRecodeProduct struct{
	Date string `json:"date"`
	IDProduct int `json:"idproduct"`
	Cost float32 `json:"cost"`
	Amount int `json:"amount"`
}