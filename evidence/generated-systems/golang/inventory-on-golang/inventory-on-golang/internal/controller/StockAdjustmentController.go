package controller

import (
    StockAdjustmentDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to StockAdjustmentDAO for database creation
//----------------------------------------------------------------------------
func CreateStockAdjustment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StockAdjustment model
	//----------------------------------------------------------------------------
	data := model.StockAdjustment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StockAdjustment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment data access object to create
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.CreateStockAdjustment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to StockAdjustmentDAO to find the relevant StockAdjustment
//----------------------------------------------------------------------------
func GetStockAdjustment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StockAdjustment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.GetStockAdjustment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to StockAdjustmentDAO for database read of all StockAdjustments
//----------------------------------------------------------------------------
func GetAllStockAdjustment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.GetAllStockAdjustment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to StockAdjustmentDAO for database save
//----------------------------------------------------------------------------
func UpdateStockAdjustment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StockAdjustment model
	//----------------------------------------------------------------------------
	var data = model.StockAdjustment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StockAdjustment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.UpdateStockAdjustment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to StockAdjustmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteStockAdjustment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StockAdjustment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := StockAdjustmentDAO.DeleteStockAdjustment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a StockAdjustment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToStockAdjustment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.AssignWarehouseToStockAdjustment(stockAdjustmentId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a StockAdjustment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromStockAdjustment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.UnassignWarehouseFromStockAdjustment(stockAdjustmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more linesIds as a Lines to a StockAdjustment
	//----------------------------------------------------------------------------
func AddLinesToStockAdjustment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockAdjustmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.AddLinesToStockAdjustment(stockAdjustmentId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a StockAdjustment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromStockAdjustment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockAdjustmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.RemoveLinesFromStockAdjustment(stockAdjustmentId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a StockAdjustment
	//----------------------------------------------------------------------------
func AddTransactionsToStockAdjustment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockAdjustmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.AddTransactionsToStockAdjustment(stockAdjustmentId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a StockAdjustment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromStockAdjustment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockAdjustmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustment DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentDAO.RemoveTransactionsFromStockAdjustment(stockAdjustmentId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
