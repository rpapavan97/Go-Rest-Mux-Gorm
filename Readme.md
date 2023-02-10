# SellerApp - B02 - Build an order management service (Advanced)

Steps to setup SellerApp B02 on local machine

# Pre-installation steps
1. Clone the SellerApp
2. Install GoLang
3. Install Mysql
4. Change the DB_PASSWORD value in constants.go file according to your settings
5. Open mysql console and connect to the DB server
6. Create database sellerapp (CREATE DATABASE sellerapp;)

# Start Golang server
go run main.go

# Create order 
Curl command - 
```
curl -H 'Content-Type: application/json' -d '{"status": "PENDING_INVOICE","items": [{"description": "IPhone 10X","quantity": 1,"price": 2534.56}],"total": 2534.56,"currency_unit": "USD"}' -X POST http://localhost:8080/orders
```

# Get Orders by updated_at DESC
Curl command - 
```
curl http://localhost:8080/orders
```

# Get order by order_id
Curl command - 
```
curl http://localhost:8080/orders/1
```

# Update order by order_id (Currently updates only status of order)
Curl command - 
```
curl -H 'Content-Type: application/json' -d '{"status": "COMPLETED"}' -X PUT http://localhost:8080/orders/1
```

# Delete order by order_id
Curl command - 
```
curl -X DELETE http://localhost:8080/orders/1
```