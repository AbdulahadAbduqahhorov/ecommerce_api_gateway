package models


type Order struct{
     Id string `json:"id"`   
     CustomerName string `json:"customer_name" binding:"required"`
     CustomerAddress string `json:"customer_address" binding:"required"`
     CustomerPhone string `json:"customer_phone" binding:"required"`
     TotalPrice int `json:"total_price" binding:"required"`
}
type OrderItem struct{
	ProductId string `json:"product_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}
type CreateOrderModel struct {
	CustomerName string `json:"customer_name" binding:"required"`
	CustomerAddress string `json:"customer_address" binding:"required"`
	CustomerPhone string `json:"customer_phone" binding:"required"`
	OrderItems []OrderItem `json:"order_items"`
}

// type GetOrderByIdModel struct {

// }
type GetAllOrderModel struct {
	Orders []Order `json:"products"`
	Count     int32           `json:"count"`
}

