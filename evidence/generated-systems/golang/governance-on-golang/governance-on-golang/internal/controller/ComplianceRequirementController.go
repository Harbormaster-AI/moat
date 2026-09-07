package controller

import (
    ComplianceRequirementDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ComplianceRequirementDAO for database creation
//----------------------------------------------------------------------------
func CreateComplianceRequirement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ComplianceRequirement model
	//----------------------------------------------------------------------------
	data := model.ComplianceRequirement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ComplianceRequirement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement data access object to create
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.CreateComplianceRequirement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ComplianceRequirementDAO to find the relevant ComplianceRequirement
//----------------------------------------------------------------------------
func GetComplianceRequirement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ComplianceRequirement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.GetComplianceRequirement(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ComplianceRequirementDAO for database read of all ComplianceRequirements
//----------------------------------------------------------------------------
func GetAllComplianceRequirement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.GetAllComplianceRequirement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ComplianceRequirementDAO for database save
//----------------------------------------------------------------------------
func UpdateComplianceRequirement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ComplianceRequirement model
	//----------------------------------------------------------------------------
	var data = model.ComplianceRequirement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ComplianceRequirement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.UpdateComplianceRequirement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ComplianceRequirementDAO for database deletion
//----------------------------------------------------------------------------
func DeleteComplianceRequirement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ComplianceRequirement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ComplianceRequirementDAO.DeleteComplianceRequirement(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ComplianceProgram on a ComplianceRequirement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignComplianceProgramToComplianceRequirement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	complianceProgramId,_ := strconv.ParseUint( vars["complianceProgramId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.AssignComplianceProgramToComplianceRequirement(complianceRequirementId, complianceProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ComplianceProgram on a ComplianceRequirement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignComplianceProgramFromComplianceRequirement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.UnassignComplianceProgramFromComplianceRequirement(complianceRequirementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a ComplianceRequirement
	//----------------------------------------------------------------------------
func AddPoliciesToComplianceRequirement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.AddPoliciesToComplianceRequirement(complianceRequirementId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a ComplianceRequirement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromComplianceRequirement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.RemovePoliciesFromComplianceRequirement(complianceRequirementId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more controlsIds as a Controls to a ComplianceRequirement
	//----------------------------------------------------------------------------
func AddControlsToComplianceRequirement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.AddControlsToComplianceRequirement(complianceRequirementId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlsIds as a Controls from a ComplianceRequirement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlsFromComplianceRequirement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.RemoveControlsFromComplianceRequirement(complianceRequirementId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more obligationsIds as a Obligations to a ComplianceRequirement
	//----------------------------------------------------------------------------
func AddObligationsToComplianceRequirement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.AddObligationsToComplianceRequirement(complianceRequirementId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more obligationsIds as a Obligations from a ComplianceRequirement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObligationsFromComplianceRequirement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceRequirementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceRequirement DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceRequirementDAO.RemoveObligationsFromComplianceRequirement(complianceRequirementId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
