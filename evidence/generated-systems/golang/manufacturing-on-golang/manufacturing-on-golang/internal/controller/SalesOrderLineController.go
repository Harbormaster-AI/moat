package controller

import (
    SalesOrderLineDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SalesOrderLineDAO for database creation
//----------------------------------------------------------------------------
func CreateSalesOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesOrderLine model
	//----------------------------------------------------------------------------
	data := model.SalesOrderLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesOrderLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.CreateSalesOrderLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SalesOrderLineDAO to find the relevant SalesOrderLine
//----------------------------------------------------------------------------
func GetSalesOrderLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesOrderLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.GetSalesOrderLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SalesOrderLineDAO for database read of all SalesOrderLines
//----------------------------------------------------------------------------
func GetAllSalesOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.GetAllSalesOrderLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SalesOrderLineDAO for database save
//----------------------------------------------------------------------------
func UpdateSalesOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesOrderLine model
	//----------------------------------------------------------------------------
	var data = model.SalesOrderLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesOrderLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.UpdateSalesOrderLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SalesOrderLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSalesOrderLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesOrderLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SalesOrderLineDAO.DeleteSalesOrderLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SalesOrder on a SalesOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSalesOrderToSalesOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesOrderId,_ := strconv.ParseUint( vars["salesOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.AssignSalesOrderToSalesOrderLine(salesOrderLineId, salesOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SalesOrder on a SalesOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSalesOrderFromSalesOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.UnassignSalesOrderFromSalesOrderLine(salesOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Item on a SalesOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToSalesOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.AssignItemToSalesOrderLine(salesOrderLineId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a SalesOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromSalesOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderLineDAO.UnassignItemFromSalesOrderLine(salesOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


