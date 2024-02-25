package dto

type (
	User struct {
		ID       string `json:"id"`
		FullName string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
		Address  string `json:"address"`
	}

	UserWithoutPassword struct {
		ID       string `json:"id"`
		FullName string `json:"name"`
		Email    string `json:"email"`
		Address  string `json:"address"`
	}

	GetUserDetailsResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		User    User   `json:"user"`
	}

	GetUserDetailsRequest struct {
		AccessToken string `json:"token"`
	}

	LoginUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUserResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Token   string `json:"token"`
	}

	RegisterUserRequest struct {
		ID       string `json:"id"`
		FullName string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Address  string `json:"address"`
	}

	RegisterUserResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		User    User   `json:"user"`
	}
)
