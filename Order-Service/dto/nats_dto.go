package dto

type (
	NatsPurchaseOrder struct {
		UserId  string
		Product Product
	}
	NatsUser struct {
		ID       string `json:"id"`
		FullName string `json:"name"`
		Email    string `json:"email"`
		Address  string `json:"address"`
	}
)
