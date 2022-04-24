package model

type LoginSellerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSellerRes struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	AlamatPickup string `json:"alamat_pickup"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateSeller struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	AlamatPickup string `json:"alamat_pickup"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UpdateSeller struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	AlamatPickup string `json:"alamat_pickup"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
