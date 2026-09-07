package controller

import (
    ReplenishmentPolicyDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ReplenishmentPolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ReplenishmentPolicy model
	//----------------------------------------------------------------------------
	data := model.ReplenishmentPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ReplenishmentPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.CreateReplenishmentPolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ReplenishmentPolicyDAO to find the relevant ReplenishmentPolicy
//----------------------------------------------------------------------------
func GetReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ReplenishmentPolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.GetReplenishmentPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ReplenishmentPolicyDAO for database read of all ReplenishmentPolicys
//----------------------------------------------------------------------------
func GetAllReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.GetAllReplenishmentPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ReplenishmentPolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ReplenishmentPolicy model
	//----------------------------------------------------------------------------
	var data = model.ReplenishmentPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ReplenishmentPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.UpdateReplenishmentPolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ReplenishmentPolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ReplenishmentPolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ReplenishmentPolicyDAO.DeleteReplenishmentPolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a ReplenishmentPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	replenishmentPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.AssignSkuToReplenishmentPolicy(replenishmentPolicyId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a ReplenishmentPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromReplenishmentPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	replenishmentPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.UnassignSkuFromReplenishmentPolicy(replenishmentPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a ReplenishmentPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	replenishmentPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.AssignWarehouseToReplenishmentPolicy(replenishmentPolicyId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a ReplenishmentPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromReplenishmentPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	replenishmentPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.UnassignWarehouseFromReplenishmentPolicy(replenishmentPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a ReplenishmentPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToReplenishmentPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	replenishmentPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.AssignLocationToReplenishmentPolicy(replenishmentPolicyId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a ReplenishmentPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromReplenishmentPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	replenishmentPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReplenishmentPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := ReplenishmentPolicyDAO.UnassignLocationFromReplenishmentPolicy(replenishmentPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


