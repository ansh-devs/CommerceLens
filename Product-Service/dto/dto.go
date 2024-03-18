package dto

type (
	Product struct {
		ID          string `json:"id"`
		ProductName string `json:"product_name"`
		Description string `json:"description"`
		Price       string `json:"price"`
	}

	GetProductReq struct {
		ProductID string `json:"product_id"`
	}
	GetProductResp struct {
		Status  string  `json:"status"`
		Message string  `json:"message"`
		Product Product `json:"product"`
	}
	GetAllProductsResp struct {
		Status  string    `json:"status"`
		Message string    `json:"message"`
		Product []Product `json:"products"`
	}

	PurchaseOrderReq struct {
		UserAccessToken string `json:"token"`
		ProductID       string `json:"product_id"`
	}
)
