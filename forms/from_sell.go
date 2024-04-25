package forms

type SellProduct struct {
	IDProduct      int    `json:"idproduct"`
	Amount         int `json:"amount"`
	Price_Cost		 float32    `json:"cost_price"`
	Price		 float32    `json:"price"`
}

type Histony struct {
	Name      string    `json:"name"`
	Amount         int `json:"amount"`
	Price_Total		 float32    `json:"price_total"`
}