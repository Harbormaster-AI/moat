package controller

import (
    StockKeepingUnitDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to StockKeepingUnitDAO for database creation
//----------------------------------------------------------------------------
func CreateStockKeepingUnit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StockKeepingUnit model
	//----------------------------------------------------------------------------
	data := model.StockKeepingUnit{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StockKeepingUnit model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit data access object to create
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.CreateStockKeepingUnit( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to StockKeepingUnitDAO to find the relevant StockKeepingUnit
//----------------------------------------------------------------------------
func GetStockKeepingUnit(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StockKeepingUnit data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.GetStockKeepingUnit(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to StockKeepingUnitDAO for database read of all StockKeepingUnits
//----------------------------------------------------------------------------
func GetAllStockKeepingUnit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit data access object to get all
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.GetAllStockKeepingUnit()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to StockKeepingUnitDAO for database save
//----------------------------------------------------------------------------
func UpdateStockKeepingUnit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StockKeepingUnit model
	//----------------------------------------------------------------------------
	var data = model.StockKeepingUnit{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StockKeepingUnit model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.UpdateStockKeepingUnit(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to StockKeepingUnitDAO for database deletion
//----------------------------------------------------------------------------
func DeleteStockKeepingUnit(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StockKeepingUnit data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := StockKeepingUnitDAO.DeleteStockKeepingUnit(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a StockKeepingUnit
	//----------------------------------------------------------------------------
func AddInventoryItemsToStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.AddInventoryItemsToStockKeepingUnit(stockKeepingUnitId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a StockKeepingUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.RemoveInventoryItemsFromStockKeepingUnit(stockKeepingUnitId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more uomConversionsIds as a UomConversions to a StockKeepingUnit
	//----------------------------------------------------------------------------
func AddUomConversionsToStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	uomConversionsIds,_ := vars["uomConversionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.AddUomConversionsToStockKeepingUnit(stockKeepingUnitId, uomConversionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more uomConversionsIds as a UomConversions from a StockKeepingUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUomConversionsFromStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	uomConversionsIds,_ := vars["uomConversionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.RemoveUomConversionsFromStockKeepingUnit(stockKeepingUnitId, uomConversionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more replenishmentPoliciesIds as a ReplenishmentPolicies to a StockKeepingUnit
	//----------------------------------------------------------------------------
func AddReplenishmentPoliciesToStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	replenishmentPoliciesIds,_ := vars["replenishmentPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.AddReplenishmentPoliciesToStockKeepingUnit(stockKeepingUnitId, replenishmentPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more replenishmentPoliciesIds as a ReplenishmentPolicies from a StockKeepingUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReplenishmentPoliciesFromStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	replenishmentPoliciesIds,_ := vars["replenishmentPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.RemoveReplenishmentPoliciesFromStockKeepingUnit(stockKeepingUnitId, replenishmentPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more lotsIds as a Lots to a StockKeepingUnit
	//----------------------------------------------------------------------------
func AddLotsToStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotsIds,_ := vars["lotsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.AddLotsToStockKeepingUnit(stockKeepingUnitId, lotsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more lotsIds as a Lots from a StockKeepingUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLotsFromStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotsIds,_ := vars["lotsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.RemoveLotsFromStockKeepingUnit(stockKeepingUnitId, lotsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a StockKeepingUnit
	//----------------------------------------------------------------------------
func AddSerialNumbersToStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.AddSerialNumbersToStockKeepingUnit(stockKeepingUnitId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a StockKeepingUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromStockKeepingUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockKeepingUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockKeepingUnit DAO
	//----------------------------------------------------------------------------
	requestResult := StockKeepingUnitDAO.RemoveSerialNumbersFromStockKeepingUnit(stockKeepingUnitId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
