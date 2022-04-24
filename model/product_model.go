package model

type CreateProduct struct { 
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price int `json:"price"`
	SellerId int `json:"seller_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateProduct struct { 
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price int `json:"price"`
	SellerId int `json:"seller_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllProductReq struct { 
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	Price int `json:"price"`
	SellerId int `json:"seller_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ShowSingleProduct struct { 
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price int `json:"price"`
	SellerId int `json:"seller_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ShowProduct struct { 
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	SellerName string `json:"seller_name"`
	Description string `json:"description"`
	Price int `json:"price"`
	SellerId int `json:"seller_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}