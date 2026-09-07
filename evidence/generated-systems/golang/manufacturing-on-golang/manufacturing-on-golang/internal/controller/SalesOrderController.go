package controller

import (
    SalesOrderDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SalesOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateSalesOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesOrder model
	//----------------------------------------------------------------------------
	data := model.SalesOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.CreateSalesOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SalesOrderDAO to find the relevant SalesOrder
//----------------------------------------------------------------------------
func GetSalesOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.GetSalesOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SalesOrderDAO for database read of all SalesOrders
//----------------------------------------------------------------------------
func GetAllSalesOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.GetAllSalesOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SalesOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateSalesOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesOrder model
	//----------------------------------------------------------------------------
	var data = model.SalesOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.UpdateSalesOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SalesOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSalesOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SalesOrderDAO.DeleteSalesOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a SalesOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToSalesOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.AssignCustomerToSalesOrder(salesOrderId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a SalesOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromSalesOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.UnassignCustomerFromSalesOrder(salesOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plant on a SalesOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToSalesOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.AssignPlantToSalesOrder(salesOrderId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a SalesOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromSalesOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.UnassignPlantFromSalesOrder(salesOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more linesIds as a Lines to a SalesOrder
	//----------------------------------------------------------------------------
func AddLinesToSalesOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.AddLinesToSalesOrder(salesOrderId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a SalesOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromSalesOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.RemoveLinesFromSalesOrder(salesOrderId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more workOrdersIds as a WorkOrders to a SalesOrder
	//----------------------------------------------------------------------------
func AddWorkOrdersToSalesOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.AddWorkOrdersToSalesOrder(salesOrderId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workOrdersIds as a WorkOrders from a SalesOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkOrdersFromSalesOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesOrder DAO
	//----------------------------------------------------------------------------
	requestResult := SalesOrderDAO.RemoveWorkOrdersFromSalesOrder(salesOrderId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
