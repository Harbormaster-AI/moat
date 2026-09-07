package controller

import (
    ConditionDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ConditionDAO for database creation
//----------------------------------------------------------------------------
func CreateCondition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Condition model
	//----------------------------------------------------------------------------
	data := model.Condition{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Condition model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Condition data access object to create
	//----------------------------------------------------------------------------
	requestResult := ConditionDAO.CreateCondition( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ConditionDAO to find the relevant Condition
//----------------------------------------------------------------------------
func GetCondition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Condition data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConditionDAO.GetCondition(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ConditionDAO for database read of all Conditions
//----------------------------------------------------------------------------
func GetAllCondition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Condition data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ConditionDAO.GetAllCondition()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ConditionDAO for database save
//----------------------------------------------------------------------------
func UpdateCondition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Condition model
	//----------------------------------------------------------------------------
	var data = model.Condition{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Condition model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Condition data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConditionDAO.UpdateCondition(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ConditionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCondition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Condition data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ConditionDAO.DeleteCondition(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a Condition
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToCondition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conditionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Condition DAO
	//----------------------------------------------------------------------------
	requestResult := ConditionDAO.AssignPatientToCondition(conditionId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a Condition
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromCondition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conditionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Condition DAO
	//----------------------------------------------------------------------------
	requestResult := ConditionDAO.UnassignPatientFromCondition(conditionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


