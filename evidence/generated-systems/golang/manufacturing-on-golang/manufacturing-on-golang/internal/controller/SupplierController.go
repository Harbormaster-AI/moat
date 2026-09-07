package controller

import (
    SupplierDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SupplierDAO for database creation
//----------------------------------------------------------------------------
func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Supplier model
	//----------------------------------------------------------------------------
	data := model.Supplier{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Supplier model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Supplier data access object to create
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.CreateSupplier( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SupplierDAO to find the relevant Supplier
//----------------------------------------------------------------------------
func GetSupplier(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Supplier data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.GetSupplier(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SupplierDAO for database read of all Suppliers
//----------------------------------------------------------------------------
func GetAllSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Supplier data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.GetAllSupplier()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SupplierDAO for database save
//----------------------------------------------------------------------------
func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Supplier model
	//----------------------------------------------------------------------------
	var data = model.Supplier{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Supplier model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Supplier data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.UpdateSupplier(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SupplierDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Supplier data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SupplierDAO.DeleteSupplier(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more enterprisesIds as a Enterprises to a Supplier
	//----------------------------------------------------------------------------
func AddEnterprisesToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enterprisesIds,_ := vars["enterprisesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddEnterprisesToSupplier(supplierId, enterprisesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more enterprisesIds as a Enterprises from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEnterprisesFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enterprisesIds,_ := vars["enterprisesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveEnterprisesFromSupplier(supplierId, enterprisesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more itemsIds as a Items to a Supplier
	//----------------------------------------------------------------------------
func AddItemsToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddItemsToSupplier(supplierId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more itemsIds as a Items from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveItemsFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveItemsFromSupplier(supplierId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more purchaseOrdersIds as a PurchaseOrders to a Supplier
	//----------------------------------------------------------------------------
func AddPurchaseOrdersToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	purchaseOrdersIds,_ := vars["purchaseOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddPurchaseOrdersToSupplier(supplierId, purchaseOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more purchaseOrdersIds as a PurchaseOrders from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePurchaseOrdersFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	purchaseOrdersIds,_ := vars["purchaseOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemovePurchaseOrdersFromSupplier(supplierId, purchaseOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
