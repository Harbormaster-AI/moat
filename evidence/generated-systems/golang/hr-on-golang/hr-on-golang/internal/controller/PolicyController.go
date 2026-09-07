package controller

import (
    PolicyDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PolicyDAO for database creation
//----------------------------------------------------------------------------
func CreatePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Policy model
	//----------------------------------------------------------------------------
	data := model.Policy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Policy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Policy data access object to create
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.CreatePolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PolicyDAO to find the relevant Policy
//----------------------------------------------------------------------------
func GetPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Policy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.GetPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PolicyDAO for database read of all Policys
//----------------------------------------------------------------------------
func GetAllPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Policy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.GetAllPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PolicyDAO for database save
//----------------------------------------------------------------------------
func UpdatePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Policy model
	//----------------------------------------------------------------------------
	var data = model.Policy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Policy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Policy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UpdatePolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeletePolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Policy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PolicyDAO.DeletePolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Policy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AssignOrganizationToPolicy(policyId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Policy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UnassignOrganizationFromPolicy(policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more acknowledgementsIds as a Acknowledgements to a Policy
	//----------------------------------------------------------------------------
func AddAcknowledgementsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	acknowledgementsIds,_ := vars["acknowledgementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddAcknowledgementsToPolicy(policyId, acknowledgementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more acknowledgementsIds as a Acknowledgements from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAcknowledgementsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	acknowledgementsIds,_ := vars["acknowledgementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveAcknowledgementsFromPolicy(policyId, acknowledgementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
