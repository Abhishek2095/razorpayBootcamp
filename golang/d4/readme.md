# ecommerce backend

## Installation
- clone the repository 
- create a database in mysql with name `ecommerce`
- configure db parameters(Host, Port, Password etc) in `Config/database.go`
- cd to project directory 
- and run
```
go run main.go
```

## Availabe Endpoints
Method | Endpoint | Description
--- | --- | ---
`POST`   | *`/product`* | adds a product
`PATCH`  | *`/product`* | patches/updates a product
`GET`    | *`/product`* | return all products
`GET`    | *`/product/:id`* | returns product with :id
`GET`    | *`/order/order-history`* | return all orders made
`POST`   | *`/order`* | make an order
`GET`    | *`/order/:id`* | retrieve order :id
`GET`    | *`/order/order-history/:id`* | retrieve order history of customer :id
`POST`   | *`/order/multiple`* | make multiple ordes
`POST`   | *`/customer`* | add a customer
`GET`    | *`/customer/:id`* | view customer :id