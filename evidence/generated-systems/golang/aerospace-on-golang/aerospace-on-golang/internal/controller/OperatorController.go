package controller

import (
    OperatorDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OperatorDAO for database creation
//----------------------------------------------------------------------------
func CreateOperator(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Operator model
	//----------------------------------------------------------------------------
	data := model.Operator{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Operator model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Operator data access object to create
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.CreateOperator( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OperatorDAO to find the relevant Operator
//----------------------------------------------------------------------------
func GetOperator(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Operator data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.GetOperator(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OperatorDAO for database read of all Operators
//----------------------------------------------------------------------------
func GetAllOperator(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Operator data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.GetAllOperator()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OperatorDAO for database save
//----------------------------------------------------------------------------
func UpdateOperator(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Operator model
	//----------------------------------------------------------------------------
	var data = model.Operator{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Operator model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Operator data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.UpdateOperator(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OperatorDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOperator(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Operator data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OperatorDAO.DeleteOperator(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SalesRegion on a Operator
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSalesRegionToOperator(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operatorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesRegionId,_ := strconv.ParseUint( vars["salesRegionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operator DAO
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.AssignSalesRegionToOperator(operatorId, salesRegionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SalesRegion on a Operator
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSalesRegionFromOperator( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operatorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operator DAO
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.UnassignSalesRegionFromOperator(operatorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more aircraftOrdersIds as a AircraftOrders to a Operator
	//----------------------------------------------------------------------------
func AddAircraftOrdersToOperator(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	operatorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftOrdersIds,_ := vars["aircraftOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Operator DAO
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.AddAircraftOrdersToOperator(operatorId, aircraftOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more aircraftOrdersIds as a AircraftOrders from a Operator
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAircraftOrdersFromOperator(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	operatorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftOrdersIds,_ := vars["aircraftOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Operator DAO
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.RemoveAircraftOrdersFromOperator(operatorId, aircraftOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more operatedAircraftIds as a OperatedAircraft to a Operator
	//----------------------------------------------------------------------------
func AddOperatedAircraftToOperator(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	operatorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatedAircraftIds,_ := vars["operatedAircraftIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Operator DAO
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.AddOperatedAircraftToOperator(operatorId, operatedAircraftIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more operatedAircraftIds as a OperatedAircraft from a Operator
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOperatedAircraftFromOperator(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	operatorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatedAircraftIds,_ := vars["operatedAircraftIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Operator DAO
	//----------------------------------------------------------------------------
	requestResult := OperatorDAO.RemoveOperatedAircraftFromOperator(operatorId, operatedAircraftIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
