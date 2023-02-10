package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sellerapp/models"
	"sellerapp/utils"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDB() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", utils.DB_USER, utils.DB_PASSWORD, utils.DB_HOST, utils.DB_PORT, utils.DB_NAME)
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(true)
	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&models.Order{}, &models.Item{})
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	// Creates new order by inserting records in the `orders` and `items` table
	db.Create(&order)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []models.Order
	db.Preload("Items").Order("updated_at DESC").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) { // Get order by order_id
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["order_id"]

	var order models.Order
	db.Preload("Items").First(&order, inputOrderID)
	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) { //Only updates status currently according to description
	var updatedOrder models.Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["order_id"]

	var order models.Order
	db.Preload("Items").First(&order, inputOrderID)

	order.Status = updatedOrder.Status

	db.Save(&order)
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputOrderID := params["order_id"]
	// Convert `order_id` string param to uint64
	id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("order_id = ?", idToDelete).Delete(&models.Item{})  // Delete item row from items table
	db.Where("order_id = ?", idToDelete).Delete(&models.Order{}) // Soft delete order by updating deleted_at column in orders table
	w.WriteHeader(http.StatusNoContent)
}
