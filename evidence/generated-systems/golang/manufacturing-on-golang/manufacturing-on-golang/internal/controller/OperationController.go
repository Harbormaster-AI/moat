package controller

import (
    OperationDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OperationDAO for database creation
//----------------------------------------------------------------------------
func CreateOperation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Operation model
	//----------------------------------------------------------------------------
	data := model.Operation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Operation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Operation data access object to create
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.CreateOperation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OperationDAO to find the relevant Operation
//----------------------------------------------------------------------------
func GetOperation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Operation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.GetOperation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OperationDAO for database read of all Operations
//----------------------------------------------------------------------------
func GetAllOperation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Operation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.GetAllOperation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OperationDAO for database save
//----------------------------------------------------------------------------
func UpdateOperation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Operation model
	//----------------------------------------------------------------------------
	var data = model.Operation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Operation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Operation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.UpdateOperation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OperationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOperation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Operation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OperationDAO.DeleteOperation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Routing on a Operation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRoutingToOperation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	routingId,_ := strconv.ParseUint( vars["routingId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operation DAO
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.AssignRoutingToOperation(operationId, routingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Routing on a Operation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRoutingFromOperation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operation DAO
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.UnassignRoutingFromOperation(operationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkCenter on a Operation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkCenterToOperation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCenterId,_ := strconv.ParseUint( vars["workCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operation DAO
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.AssignWorkCenterToOperation(operationId, workCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkCenter on a Operation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkCenterFromOperation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operation DAO
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.UnassignWorkCenterFromOperation(operationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InspectionPlan on a Operation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInspectionPlanToOperation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inspectionPlanId,_ := strconv.ParseUint( vars["inspectionPlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operation DAO
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.AssignInspectionPlanToOperation(operationId, inspectionPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InspectionPlan on a Operation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInspectionPlanFromOperation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	operationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Operation DAO
	//----------------------------------------------------------------------------
	requestResult := OperationDAO.UnassignInspectionPlanFromOperation(operationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


