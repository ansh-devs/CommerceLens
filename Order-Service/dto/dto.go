package dto

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
		Status          string `json:"status"`
		ShippingAddress string `json:"shipping_address"`
	}

	PlaceOrderReq struct {
		ProductID string `json:"product_id"`
		UserID    string `json:"user_id"`
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
	PlaceOrderResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		//Order   Order  `json:"order"`
	}
	GetOrderResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Order   Order  `json:"order"`
	}
	CancelOrderResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	GetAllUserOrdersResp struct {
		Status  string  `json:"status"`
		Message string  `json:"message"`
		Orders  []Order `json:"user_orders"`
	}

	Product struct {
		ID          string `json:"id"`
		ProductName string `json:"product_name"`
		Description string `json:"description"`
		Price       string `json:"price"`
	}
)
