package controller

import (
    ProcedureDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProcedureDAO for database creation
//----------------------------------------------------------------------------
func CreateProcedure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Procedure model
	//----------------------------------------------------------------------------
	data := model.Procedure{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Procedure model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.CreateProcedure( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProcedureDAO to find the relevant Procedure
//----------------------------------------------------------------------------
func GetProcedure(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Procedure data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.GetProcedure(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProcedureDAO for database read of all Procedures
//----------------------------------------------------------------------------
func GetAllProcedure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Procedure data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.GetAllProcedure()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProcedureDAO for database save
//----------------------------------------------------------------------------
func UpdateProcedure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Procedure model
	//----------------------------------------------------------------------------
	var data = model.Procedure{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Procedure model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.UpdateProcedure(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProcedureDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProcedure(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Procedure data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProcedureDAO.DeleteProcedure(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a Procedure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToProcedure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.AssignEncounterToProcedure(procedureId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a Procedure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromProcedure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.UnassignEncounterFromProcedure(procedureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Performer on a Procedure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPerformerToProcedure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performerId,_ := strconv.ParseUint( vars["performerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.AssignPerformerToProcedure(procedureId, performerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Performer on a Procedure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPerformerFromProcedure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.UnassignPerformerFromProcedure(procedureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ProcedureOrder on a Procedure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProcedureOrderToProcedure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	procedureOrderId,_ := strconv.ParseUint( vars["procedureOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.AssignProcedureOrderToProcedure(procedureId, procedureOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ProcedureOrder on a Procedure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProcedureOrderFromProcedure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Procedure DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureDAO.UnassignProcedureOrderFromProcedure(procedureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


