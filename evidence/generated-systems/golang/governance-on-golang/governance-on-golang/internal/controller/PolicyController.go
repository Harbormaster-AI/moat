package controller

import (
    PolicyDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
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
	// adds one or more ownersIds as a Owners to a Policy
	//----------------------------------------------------------------------------
func AddOwnersToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownersIds,_ := vars["ownersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddOwnersToPolicy(policyId, ownersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownersIds as a Owners from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnersFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownersIds,_ := vars["ownersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveOwnersFromPolicy(policyId, ownersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more relatedRequirementsIds as a RelatedRequirements to a Policy
	//----------------------------------------------------------------------------
func AddRelatedRequirementsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedRequirementsIds,_ := vars["relatedRequirementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddRelatedRequirementsToPolicy(policyId, relatedRequirementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more relatedRequirementsIds as a RelatedRequirements from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRelatedRequirementsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedRequirementsIds,_ := vars["relatedRequirementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveRelatedRequirementsFromPolicy(policyId, relatedRequirementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more controlsIds as a Controls to a Policy
	//----------------------------------------------------------------------------
func AddControlsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddControlsToPolicy(policyId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlsIds as a Controls from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveControlsFromPolicy(policyId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more proceduresIds as a Procedures to a Policy
	//----------------------------------------------------------------------------
func AddProceduresToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddProceduresToPolicy(policyId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more proceduresIds as a Procedures from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProceduresFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveProceduresFromPolicy(policyId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more exceptionsIds as a Exceptions to a Policy
	//----------------------------------------------------------------------------
func AddExceptionsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exceptionsIds,_ := vars["exceptionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddExceptionsToPolicy(policyId, exceptionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more exceptionsIds as a Exceptions from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExceptionsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exceptionsIds,_ := vars["exceptionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveExceptionsFromPolicy(policyId, exceptionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more attestationsIds as a Attestations to a Policy
	//----------------------------------------------------------------------------
func AddAttestationsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	attestationsIds,_ := vars["attestationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddAttestationsToPolicy(policyId, attestationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more attestationsIds as a Attestations from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAttestationsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	attestationsIds,_ := vars["attestationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveAttestationsFromPolicy(policyId, attestationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
