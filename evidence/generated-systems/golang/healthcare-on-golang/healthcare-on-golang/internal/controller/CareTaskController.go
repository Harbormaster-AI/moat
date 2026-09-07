package controller

import (
    CareTaskDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CareTaskDAO for database creation
//----------------------------------------------------------------------------
func CreateCareTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CareTask model
	//----------------------------------------------------------------------------
	data := model.CareTask{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CareTask model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask data access object to create
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.CreateCareTask( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CareTaskDAO to find the relevant CareTask
//----------------------------------------------------------------------------
func GetCareTask(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CareTask data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.GetCareTask(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CareTaskDAO for database read of all CareTasks
//----------------------------------------------------------------------------
func GetAllCareTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CareTask data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.GetAllCareTask()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CareTaskDAO for database save
//----------------------------------------------------------------------------
func UpdateCareTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CareTask model
	//----------------------------------------------------------------------------
	var data = model.CareTask{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CareTask model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.UpdateCareTask(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CareTaskDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCareTask(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CareTask data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CareTaskDAO.DeleteCareTask(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CarePlan on a CareTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCarePlanToCareTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	carePlanId,_ := strconv.ParseUint( vars["carePlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask DAO
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.AssignCarePlanToCareTask(careTaskId, carePlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CarePlan on a CareTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCarePlanFromCareTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask DAO
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.UnassignCarePlanFromCareTask(careTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AssignedTo on a CareTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAssignedToToCareTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignedToId,_ := strconv.ParseUint( vars["assignedToId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask DAO
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.AssignAssignedToToCareTask(careTaskId, assignedToId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AssignedTo on a CareTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAssignedToFromCareTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask DAO
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.UnassignAssignedToFromCareTask(careTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a CareTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToCareTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask DAO
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.AssignEncounterToCareTask(careTaskId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a CareTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromCareTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTask DAO
	//----------------------------------------------------------------------------
	requestResult := CareTaskDAO.UnassignEncounterFromCareTask(careTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


