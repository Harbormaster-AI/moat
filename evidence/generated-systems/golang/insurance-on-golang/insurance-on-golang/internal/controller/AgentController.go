package controller

import (
    AgentDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AgentDAO for database creation
//----------------------------------------------------------------------------
func CreateAgent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Agent model
	//----------------------------------------------------------------------------
	data := model.Agent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Agent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Agent data access object to create
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.CreateAgent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AgentDAO to find the relevant Agent
//----------------------------------------------------------------------------
func GetAgent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Agent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.GetAgent(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AgentDAO for database read of all Agents
//----------------------------------------------------------------------------
func GetAllAgent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Agent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.GetAllAgent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AgentDAO for database save
//----------------------------------------------------------------------------
func UpdateAgent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Agent model
	//----------------------------------------------------------------------------
	var data = model.Agent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Agent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Agent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.UpdateAgent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AgentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAgent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Agent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AgentDAO.DeleteAgent(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Distributor on a Agent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDistributorToAgent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	agentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	distributorId,_ := strconv.ParseUint( vars["distributorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Agent DAO
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.AssignDistributorToAgent(agentId, distributorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Distributor on a Agent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDistributorFromAgent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	agentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Agent DAO
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.UnassignDistributorFromAgent(agentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a Agent
	//----------------------------------------------------------------------------
func AddPoliciesToAgent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agent DAO
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.AddPoliciesToAgent(agentId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a Agent
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromAgent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agent DAO
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.RemovePoliciesFromAgent(agentId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more customersIds as a Customers to a Agent
	//----------------------------------------------------------------------------
func AddCustomersToAgent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agent DAO
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.AddCustomersToAgent(agentId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more customersIds as a Customers from a Agent
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCustomersFromAgent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agent DAO
	//----------------------------------------------------------------------------
	requestResult := AgentDAO.RemoveCustomersFromAgent(agentId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
