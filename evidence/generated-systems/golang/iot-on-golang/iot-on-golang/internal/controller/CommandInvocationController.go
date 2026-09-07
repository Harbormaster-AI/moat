package controller

import (
    CommandInvocationDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CommandInvocationDAO for database creation
//----------------------------------------------------------------------------
func CreateCommandInvocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CommandInvocation model
	//----------------------------------------------------------------------------
	data := model.CommandInvocation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CommandInvocation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation data access object to create
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.CreateCommandInvocation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CommandInvocationDAO to find the relevant CommandInvocation
//----------------------------------------------------------------------------
func GetCommandInvocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CommandInvocation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.GetCommandInvocation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CommandInvocationDAO for database read of all CommandInvocations
//----------------------------------------------------------------------------
func GetAllCommandInvocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.GetAllCommandInvocation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CommandInvocationDAO for database save
//----------------------------------------------------------------------------
func UpdateCommandInvocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CommandInvocation model
	//----------------------------------------------------------------------------
	var data = model.CommandInvocation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CommandInvocation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.UpdateCommandInvocation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CommandInvocationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCommandInvocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CommandInvocation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CommandInvocationDAO.DeleteCommandInvocation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a CommandInvocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToCommandInvocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.AssignDeviceToCommandInvocation(commandInvocationId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a CommandInvocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromCommandInvocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.UnassignDeviceFromCommandInvocation(commandInvocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CommandDefinition on a CommandInvocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCommandDefinitionToCommandInvocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandDefinitionId,_ := strconv.ParseUint( vars["commandDefinitionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.AssignCommandDefinitionToCommandInvocation(commandInvocationId, commandDefinitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CommandDefinition on a CommandInvocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCommandDefinitionFromCommandInvocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.UnassignCommandDefinitionFromCommandInvocation(commandInvocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Actuator on a CommandInvocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignActuatorToCommandInvocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	actuatorId,_ := strconv.ParseUint( vars["actuatorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.AssignActuatorToCommandInvocation(commandInvocationId, actuatorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Actuator on a CommandInvocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignActuatorFromCommandInvocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.UnassignActuatorFromCommandInvocation(commandInvocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a User on a CommandInvocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignUserToCommandInvocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	userId,_ := strconv.ParseUint( vars["userId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.AssignUserToCommandInvocation(commandInvocationId, userId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a User on a CommandInvocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignUserFromCommandInvocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	commandInvocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CommandInvocation DAO
	//----------------------------------------------------------------------------
	requestResult := CommandInvocationDAO.UnassignUserFromCommandInvocation(commandInvocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


