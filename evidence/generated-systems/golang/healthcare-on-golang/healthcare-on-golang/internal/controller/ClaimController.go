package controller

import (
    ClaimDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ClaimDAO for database creation
//----------------------------------------------------------------------------
func CreateClaim(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Claim model
	//----------------------------------------------------------------------------
	data := model.Claim{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Claim model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Claim data access object to create
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.CreateClaim( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ClaimDAO to find the relevant Claim
//----------------------------------------------------------------------------
func GetClaim(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Claim data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.GetClaim(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ClaimDAO for database read of all Claims
//----------------------------------------------------------------------------
func GetAllClaim(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Claim data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.GetAllClaim()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ClaimDAO for database save
//----------------------------------------------------------------------------
func UpdateClaim(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Claim model
	//----------------------------------------------------------------------------
	var data = model.Claim{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Claim model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Claim data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UpdateClaim(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ClaimDAO for database deletion
//----------------------------------------------------------------------------
func DeleteClaim(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Claim data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ClaimDAO.DeleteClaim(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignPatientToClaim(claimId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignPatientFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Coverage on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCoverageToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coverageId,_ := strconv.ParseUint( vars["coverageId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignCoverageToClaim(claimId, coverageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Coverage on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCoverageFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignCoverageFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignEncounterToClaim(claimId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignEncounterFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Payer on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPayerToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payerId,_ := strconv.ParseUint( vars["payerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignPayerToClaim(claimId, payerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Payer on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPayerFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignPayerFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more invoicesIds as a Invoices to a Claim
	//----------------------------------------------------------------------------
func AddInvoicesToClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoicesIds,_ := vars["invoicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AddInvoicesToClaim(claimId, invoicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more invoicesIds as a Invoices from a Claim
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInvoicesFromClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoicesIds,_ := vars["invoicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.RemoveInvoicesFromClaim(claimId, invoicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
