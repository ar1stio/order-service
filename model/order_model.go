package model

type CreateOrder struct {
	Id int `json:"id"`
	BuyerId int `json:"buyer_id"`
	SellerId int `json:"seller_id"`
	BuyerName string `json:"buyer_name"`
	SellerName string `json:"seller_name"`
	DeliverySourceAddress string `json:"delivery_source_address"`
	DeliveryDestinationAddress string `json:"delivery_destination_address"`
	Items int `json:"items"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	TotalPrice int `json:"total_price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateOrder struct {
	Id int `json:"id"`
	BuyerId int `json:"buyer_id"`
	SellerId int `json:"seller_id"`
	BuyerName string `json:"buyer_name"`
	SellerName string `json:"seller_name"`
	DeliverySourceAddress string `json:"delivery_source_address"`
	DeliveryDestinationAddress string `json:"delivery_destination_address"`
	Items int `json:"items"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	TotalPrice int `json:"total_price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Deliveredreq struct {
	Id int `json:"id"`
}

type ShowOrder struct {
	Id int `json:"id"`
	BuyerId int `json:"buyer_id"`
	SellerId int `json:"seller_id"`
	BuyerName string `json:"buyer_name"`
	SellerName string `json:"seller_name"`
	DeliverySourceAddress string `json:"delivery_source_address"`
	DeliveryDestinationAddress string `json:"delivery_destination_address"`
	Status int `json:"status"`
	Items string `json:"items"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	TotalPrice int `json:"total_price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
