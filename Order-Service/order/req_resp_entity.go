package order

type (
	Order struct {
		ID              string `json:"id"`
		ProductID       string `json:"product_id"`
		UserID          string `json:"user_id"`
		TotalCost       string `json:"total_cost"`
		Username        string `json:"username"`
		ProductName     string `json:"product_name"`
		Description     string `json:"description"`
		Price           string `json:"price"`
		ShippingAddress string `json:"shipping_address"`
	}

	PlaceOrderReq struct {
		ProductID   string `json:"product_id"`
		AccessToken string `json:"access_token"`
	}
	GetOrderReq struct {
		OrderID string `json:"order_id"`
	}
	CancelOrderReq struct {
		OrderID string `json:"order_id"`
	}
	GetAllUserOrdersReq struct {
		UserID string `json:"user_id"`
	}
	PlaceOrderResp struct{}
	GetOrderResp   struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Order   Order  `json:"order"`
	}
	CancelOrderResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	GetAllUserOrdersResp struct{}
)
