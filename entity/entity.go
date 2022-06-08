package entity

type FindValueInfo struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type CashInfo struct{
	Cash_1000 int `json:"cash_1000"`
	Cash_500 int `json:"cash_500"`
	Cash_100 int `json:"cash_100"`
	Cash_50 int `json:"cash_50"`
	Cash_20 int `json:"cash_20"`
	Coin_10 int `json:"coin_10"`
	Coin_5 int `json:"coin_5"`
	Coin_1 int `json:"coin_1"`
	Coin_025 int `json:"coin_025"`
}

type ResponseText struct{
	StatusCode int `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct{
	StatusCode int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}