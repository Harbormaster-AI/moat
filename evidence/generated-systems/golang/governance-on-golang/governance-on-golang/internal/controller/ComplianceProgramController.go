package controller

import (
    ComplianceProgramDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ComplianceProgramDAO for database creation
//----------------------------------------------------------------------------
func CreateComplianceProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ComplianceProgram model
	//----------------------------------------------------------------------------
	data := model.ComplianceProgram{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ComplianceProgram model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram data access object to create
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.CreateComplianceProgram( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ComplianceProgramDAO to find the relevant ComplianceProgram
//----------------------------------------------------------------------------
func GetComplianceProgram(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ComplianceProgram data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.GetComplianceProgram(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ComplianceProgramDAO for database read of all CompliancePrograms
//----------------------------------------------------------------------------
func GetAllComplianceProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.GetAllComplianceProgram()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ComplianceProgramDAO for database save
//----------------------------------------------------------------------------
func UpdateComplianceProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ComplianceProgram model
	//----------------------------------------------------------------------------
	var data = model.ComplianceProgram{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ComplianceProgram model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.UpdateComplianceProgram(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ComplianceProgramDAO for database deletion
//----------------------------------------------------------------------------
func DeleteComplianceProgram(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ComplianceProgram data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ComplianceProgramDAO.DeleteComplianceProgram(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a ComplianceProgram
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToComplianceProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.AssignOrganizationToComplianceProgram(complianceProgramId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a ComplianceProgram
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromComplianceProgram( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.UnassignOrganizationFromComplianceProgram(complianceProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more requirementsIds as a Requirements to a ComplianceProgram
	//----------------------------------------------------------------------------
func AddRequirementsToComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requirementsIds,_ := vars["requirementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.AddRequirementsToComplianceProgram(complianceProgramId, requirementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more requirementsIds as a Requirements from a ComplianceProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRequirementsFromComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requirementsIds,_ := vars["requirementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.RemoveRequirementsFromComplianceProgram(complianceProgramId, requirementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more controlsIds as a Controls to a ComplianceProgram
	//----------------------------------------------------------------------------
func AddControlsToComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.AddControlsToComplianceProgram(complianceProgramId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlsIds as a Controls from a ComplianceProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlsFromComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.RemoveControlsFromComplianceProgram(complianceProgramId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more attestationsIds as a Attestations to a ComplianceProgram
	//----------------------------------------------------------------------------
func AddAttestationsToComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	attestationsIds,_ := vars["attestationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.AddAttestationsToComplianceProgram(complianceProgramId, attestationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more attestationsIds as a Attestations from a ComplianceProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAttestationsFromComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	attestationsIds,_ := vars["attestationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.RemoveAttestationsFromComplianceProgram(complianceProgramId, attestationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more regulationsIds as a Regulations to a ComplianceProgram
	//----------------------------------------------------------------------------
func AddRegulationsToComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	regulationsIds,_ := vars["regulationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.AddRegulationsToComplianceProgram(complianceProgramId, regulationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more regulationsIds as a Regulations from a ComplianceProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRegulationsFromComplianceProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	complianceProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	regulationsIds,_ := vars["regulationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceProgram DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceProgramDAO.RemoveRegulationsFromComplianceProgram(complianceProgramId, regulationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
