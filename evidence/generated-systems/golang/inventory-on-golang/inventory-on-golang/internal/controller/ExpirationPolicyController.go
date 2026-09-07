package controller

import (
    ExpirationPolicyDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExpirationPolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateExpirationPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExpirationPolicy model
	//----------------------------------------------------------------------------
	data := model.ExpirationPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExpirationPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.CreateExpirationPolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExpirationPolicyDAO to find the relevant ExpirationPolicy
//----------------------------------------------------------------------------
func GetExpirationPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExpirationPolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.GetExpirationPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExpirationPolicyDAO for database read of all ExpirationPolicys
//----------------------------------------------------------------------------
func GetAllExpirationPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.GetAllExpirationPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExpirationPolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateExpirationPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExpirationPolicy model
	//----------------------------------------------------------------------------
	var data = model.ExpirationPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExpirationPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.UpdateExpirationPolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExpirationPolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExpirationPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExpirationPolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExpirationPolicyDAO.DeleteExpirationPolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a ExpirationPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToExpirationPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	expirationPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.AssignSkuToExpirationPolicy(expirationPolicyId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a ExpirationPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromExpirationPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	expirationPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.UnassignSkuFromExpirationPolicy(expirationPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a ExpirationPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToExpirationPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	expirationPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.AssignWarehouseToExpirationPolicy(expirationPolicyId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a ExpirationPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromExpirationPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	expirationPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExpirationPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ExpirationPolicyDAO.UnassignWarehouseFromExpirationPolicy(expirationPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


