package model

type LoginBuyerReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginBuyerRes struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	Token string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateBuyer struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBuyer struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}