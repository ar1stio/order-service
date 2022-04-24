# API Spec
# Gokomodo backend test

# Go Test Add Order
go test -run ^TestOrderController_Create$ order-service/controller

# seller-service

For buyer

<!-- Request : -->
<!-- - Header : -->
<!-- - X-Api-Key : "your secret api key" -->

## Register Buyer

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/buyer/register`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "email":"string",
    "name":"string",
    "password":"string",
    "alamat_pengiriman":"string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Update Buyer

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/buyer/update`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "id":"integer",
    "email":"string",
    "name":"string",
    "password":"string",
    "alamat_pengiriman":"string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Login Buyer

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/buyer/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "email":"string",
    "password":"string",
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data": {
        "id": "integer",
        "email": "string",
        "name": "string",
        "password": "string",
        "alamat_pengiriman": "string",
        "token": "string",
        "created_at": "string",
        "updated_at": "string"
    }
}
```

# seller-service

For seller

<!-- Request : -->
<!-- - Header : -->
<!-- - X-Api-Key : "your secret api key" -->

## Register Seller

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/seller/register`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "email":"string",
    "name":"string",
    "password":"string",
    "alamat_pickup":"string"
}
```
Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```
## Update Seller

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/seller/update`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "id":"integer",
    "email":"string",
    "name":"string",
    "password":"string",
    "alamat_pickup":"string"
}
```
Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```
## Login Seller

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/seller/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "email":"string",
    "password":"string",
}
```
Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id": "integer",
        "email": "string",
        "name": "string",
        "password": "string",
        "alamat_pickup": "string",
        "token": "string",
        "created_at": "string",
        "updated_at": "string"
    }
}
```

# product-service

For buyer and seller

<!-- Request : -->
<!-- - Header : -->
<!-- - X-Api-Key : "your secret api key" -->

## Add Product

Request :
- Method : POST
- Endpoint : `localhost:3800/product-service/seller-product/register`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "product_name":"string",
    "description":"string",
    "price":"integer",
    "seller_id":"integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Update Product

Request :
- Method : POST
- Endpoint : `localhost:3800/product-service/seller-product/update`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "id":"integer",
    "product_name":"string",
    "description":"string",
    "price":"integer",
    "seller_id":"integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Single Product fin by id

Request :
- Method : GET
- Endpoint : `localhost:3800/product-service/product/:id`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id": "integer",
        "product_name": "string",
        "description": "string",
        "price": "integer",
        "seller_id": "integer",
        "created_at": "string",
        "updated_at": "string"
    }
}
```

## All product find by filter

Request :
- Method : POST
- Endpoint : `api/member/find`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id":"integer",
    "product_name":"string",
    "price":"integer",
    "seller_id":"integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
            "id": "integer",
            "product_name": "string",
            "seller_name": "string",
            "description": "string",
            "price": "integer",
            "seller_id": "integer",
            "created_at": "string",
            "updated_at": "string"
        }
}
```

# order-service

For buyer and seller

<!-- Request : -->
<!-- - Header : -->
<!-- - X-Api-Key : "your secret api key" -->

## Add Order

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/buyer/order`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "id":"integer",
    "buyer_id":"integer",
    "seller_id":"integer",
    "buyer_name":"string",
    "seller_name":"string",
    "delivery_source_address":"string",
    "delivery_destination_address":"string",
    "status":"integer",
    "items":"string",
    "quantity":"integer",
    "price":"integer",
    "total_price":"integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Order Accepted

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/delivered`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "id":"integer",
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## All Order Product 

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/show-order-product`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id":"integer",
    "buyer_id":"integer",
    "seller_id":"integer",
    "buyer_name":"string",
    "seller_name":"string",
    "delivery_source_address":"string",
    "delivery_destination_address":"string",
    "status":"integer",
    "items":"string",
    "quantity":"integer",
    "price":"integer",
    "total_price":"integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
                {
            "id":"integer",
            "buyer_id":"integer",
            "seller_id":"integer",
            "buyer_name":"string",
            "seller_name":"string",
            "delivery_source_address":"string",
            "delivery_destination_address":"string",
            "status":"integer",
            "items":"string",
            "quantity":"integer",
            "price":"integer",
            "total_price":"integer"
        }
    ]
}
```

## All Order List

Request :
- Method : POST
- Endpoint : `localhost:3800/order-service/show-order-list`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "id":"integer",
    "buyer_id":"integer",
    "seller_id":"integer",
    "buyer_name":"string",
    "seller_name":"string",
    "delivery_source_address":"string",
    "delivery_destination_address":"string",
    "status":"integer",
    "items":"string",
    "quantity":"integer",
    "price":"integer",
    "total_price":"integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
                {
            "id":"integer",
            "buyer_id":"integer",
            "seller_id":"integer",
            "buyer_name":"string",
            "seller_name":"string",
            "delivery_source_address":"string",
            "delivery_destination_address":"string",
            "status":"integer",
            "items":"string",
            "quantity":"integer",
            "price":"integer",
            "total_price":"integer"
        }
    ]
}
```
