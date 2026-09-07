package controller

import (
    InboundShipmentDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InboundShipmentDAO for database creation
//----------------------------------------------------------------------------
func CreateInboundShipment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InboundShipment model
	//----------------------------------------------------------------------------
	data := model.InboundShipment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InboundShipment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment data access object to create
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.CreateInboundShipment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InboundShipmentDAO to find the relevant InboundShipment
//----------------------------------------------------------------------------
func GetInboundShipment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InboundShipment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.GetInboundShipment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InboundShipmentDAO for database read of all InboundShipments
//----------------------------------------------------------------------------
func GetAllInboundShipment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.GetAllInboundShipment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InboundShipmentDAO for database save
//----------------------------------------------------------------------------
func UpdateInboundShipment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InboundShipment model
	//----------------------------------------------------------------------------
	var data = model.InboundShipment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InboundShipment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.UpdateInboundShipment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InboundShipmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInboundShipment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InboundShipment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InboundShipmentDAO.DeleteInboundShipment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a InboundShipment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToInboundShipment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.AssignWarehouseToInboundShipment(inboundShipmentId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a InboundShipment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromInboundShipment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.UnassignWarehouseFromInboundShipment(inboundShipmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more linesIds as a Lines to a InboundShipment
	//----------------------------------------------------------------------------
func AddLinesToInboundShipment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inboundShipmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.AddLinesToInboundShipment(inboundShipmentId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a InboundShipment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromInboundShipment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inboundShipmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.RemoveLinesFromInboundShipment(inboundShipmentId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a InboundShipment
	//----------------------------------------------------------------------------
func AddTransactionsToInboundShipment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inboundShipmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.AddTransactionsToInboundShipment(inboundShipmentId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a InboundShipment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromInboundShipment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inboundShipmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipment DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentDAO.RemoveTransactionsFromInboundShipment(inboundShipmentId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
