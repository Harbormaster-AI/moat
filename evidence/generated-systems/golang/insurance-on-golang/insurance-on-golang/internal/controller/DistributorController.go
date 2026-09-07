package controller

import (
    DistributorDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DistributorDAO for database creation
//----------------------------------------------------------------------------
func CreateDistributor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Distributor model
	//----------------------------------------------------------------------------
	data := model.Distributor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Distributor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Distributor data access object to create
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.CreateDistributor( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DistributorDAO to find the relevant Distributor
//----------------------------------------------------------------------------
func GetDistributor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Distributor data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.GetDistributor(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DistributorDAO for database read of all Distributors
//----------------------------------------------------------------------------
func GetAllDistributor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Distributor data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.GetAllDistributor()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DistributorDAO for database save
//----------------------------------------------------------------------------
func UpdateDistributor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Distributor model
	//----------------------------------------------------------------------------
	var data = model.Distributor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Distributor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Distributor data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.UpdateDistributor(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DistributorDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDistributor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Distributor data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DistributorDAO.DeleteDistributor(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more insurersIds as a Insurers to a Distributor
	//----------------------------------------------------------------------------
func AddInsurersToDistributor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	distributorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insurersIds,_ := vars["insurersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Distributor DAO
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.AddInsurersToDistributor(distributorId, insurersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more insurersIds as a Insurers from a Distributor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInsurersFromDistributor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	distributorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insurersIds,_ := vars["insurersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Distributor DAO
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.RemoveInsurersFromDistributor(distributorId, insurersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more agentsIds as a Agents to a Distributor
	//----------------------------------------------------------------------------
func AddAgentsToDistributor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	distributorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agentsIds,_ := vars["agentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Distributor DAO
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.AddAgentsToDistributor(distributorId, agentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more agentsIds as a Agents from a Distributor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAgentsFromDistributor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	distributorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agentsIds,_ := vars["agentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Distributor DAO
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.RemoveAgentsFromDistributor(distributorId, agentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a Distributor
	//----------------------------------------------------------------------------
func AddPoliciesToDistributor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	distributorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Distributor DAO
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.AddPoliciesToDistributor(distributorId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a Distributor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromDistributor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	distributorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Distributor DAO
	//----------------------------------------------------------------------------
	requestResult := DistributorDAO.RemovePoliciesFromDistributor(distributorId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
