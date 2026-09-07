package controller

import (
    PurchaseOrderLineDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PurchaseOrderLineDAO for database creation
//----------------------------------------------------------------------------
func CreatePurchaseOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PurchaseOrderLine model
	//----------------------------------------------------------------------------
	data := model.PurchaseOrderLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PurchaseOrderLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.CreatePurchaseOrderLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PurchaseOrderLineDAO to find the relevant PurchaseOrderLine
//----------------------------------------------------------------------------
func GetPurchaseOrderLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PurchaseOrderLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.GetPurchaseOrderLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PurchaseOrderLineDAO for database read of all PurchaseOrderLines
//----------------------------------------------------------------------------
func GetAllPurchaseOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.GetAllPurchaseOrderLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PurchaseOrderLineDAO for database save
//----------------------------------------------------------------------------
func UpdatePurchaseOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PurchaseOrderLine model
	//----------------------------------------------------------------------------
	var data = model.PurchaseOrderLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PurchaseOrderLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.UpdatePurchaseOrderLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PurchaseOrderLineDAO for database deletion
//----------------------------------------------------------------------------
func DeletePurchaseOrderLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PurchaseOrderLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PurchaseOrderLineDAO.DeletePurchaseOrderLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PurchaseOrder on a PurchaseOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPurchaseOrderToPurchaseOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	purchaseOrderId,_ := strconv.ParseUint( vars["purchaseOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.AssignPurchaseOrderToPurchaseOrderLine(purchaseOrderLineId, purchaseOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PurchaseOrder on a PurchaseOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPurchaseOrderFromPurchaseOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.UnassignPurchaseOrderFromPurchaseOrderLine(purchaseOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Item on a PurchaseOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToPurchaseOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.AssignItemToPurchaseOrderLine(purchaseOrderLineId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a PurchaseOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromPurchaseOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderLineDAO.UnassignItemFromPurchaseOrderLine(purchaseOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


