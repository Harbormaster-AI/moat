package controller

import (
    OrderItemDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OrderItemDAO for database creation
//----------------------------------------------------------------------------
func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OrderItem model
	//----------------------------------------------------------------------------
	data := model.OrderItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OrderItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.CreateOrderItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OrderItemDAO to find the relevant OrderItem
//----------------------------------------------------------------------------
func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]
	
	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	//----------------------------------------------------------------------------
	// Delegate to the OrderItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.GetOrderItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OrderItemDAO for database read of all OrderItems
//----------------------------------------------------------------------------
func GetAllOrderItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the OrderItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.GetAllOrderItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OrderItemDAO for database save
//----------------------------------------------------------------------------
func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OrderItem model
	//----------------------------------------------------------------------------
	var data = model.OrderItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OrderItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.UpdateOrderItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OrderItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]

	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OrderItemDAO.DeleteOrderItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Order on a OrderItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderToOrderItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderId,_ := strconv.ParseUint( vars["orderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem DAO
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.AssignOrderToOrderItem(orderItemId, orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Order on a OrderItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderFromOrderItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem DAO
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.UnassignOrderFromOrderItem(orderItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a OrderItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToOrderItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem DAO
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.AssignProductToOrderItem(orderItemId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a OrderItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromOrderItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem DAO
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.UnassignProductFromOrderItem(orderItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PriceBookEntry on a OrderItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPriceBookEntryToOrderItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookEntryId,_ := strconv.ParseUint( vars["priceBookEntryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem DAO
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.AssignPriceBookEntryToOrderItem(orderItemId, priceBookEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PriceBookEntry on a OrderItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPriceBookEntryFromOrderItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OrderItem DAO
	//----------------------------------------------------------------------------
	requestResult := OrderItemDAO.UnassignPriceBookEntryFromOrderItem(orderItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


