package controller

import (
    AttestationDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AttestationDAO for database creation
//----------------------------------------------------------------------------
func CreateAttestation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Attestation model
	//----------------------------------------------------------------------------
	data := model.Attestation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Attestation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation data access object to create
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.CreateAttestation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AttestationDAO to find the relevant Attestation
//----------------------------------------------------------------------------
func GetAttestation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Attestation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.GetAttestation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AttestationDAO for database read of all Attestations
//----------------------------------------------------------------------------
func GetAllAttestation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Attestation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.GetAllAttestation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AttestationDAO for database save
//----------------------------------------------------------------------------
func UpdateAttestation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Attestation model
	//----------------------------------------------------------------------------
	var data = model.Attestation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Attestation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.UpdateAttestation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AttestationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAttestation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Attestation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AttestationDAO.DeleteAttestation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Control on a Attestation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignControlToAttestation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	attestationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlId,_ := strconv.ParseUint( vars["controlId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation DAO
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.AssignControlToAttestation(attestationId, controlId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Control on a Attestation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignControlFromAttestation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	attestationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation DAO
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.UnassignControlFromAttestation(attestationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Policy on a Attestation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToAttestation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	attestationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation DAO
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.AssignPolicyToAttestation(attestationId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Attestation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromAttestation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	attestationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation DAO
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.UnassignPolicyFromAttestation(attestationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ComplianceProgram on a Attestation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignComplianceProgramToAttestation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	attestationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	complianceProgramId,_ := strconv.ParseUint( vars["complianceProgramId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation DAO
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.AssignComplianceProgramToAttestation(attestationId, complianceProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ComplianceProgram on a Attestation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignComplianceProgramFromAttestation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	attestationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Attestation DAO
	//----------------------------------------------------------------------------
	requestResult := AttestationDAO.UnassignComplianceProgramFromAttestation(attestationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


