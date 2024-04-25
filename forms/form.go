package forms

type Profit struct {
	Date           string    `json:"date"`
	Profit       float32    `json:"profit"`
}

type Profitincome struct {
	Income_Month          float32    `json:"income_month"`
	Income_Year       float32    `json:"income_year"`
	Profit_Month          float32    `json:"profit_month"`
	Profit_Year       float32    `json:"profit_year"`
}
