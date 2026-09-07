package controller

import (
    ClaimDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
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
	// assigns a Policy on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignPolicyToClaim(claimId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignPolicyFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignCustomerToClaim(claimId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignCustomerFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Adjuster on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdjusterToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adjusterId,_ := strconv.ParseUint( vars["adjusterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignAdjusterToClaim(claimId, adjusterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Adjuster on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdjusterFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignAdjusterFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Incident on a Claim
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignIncidentToClaim(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	incidentId,_ := strconv.ParseUint( vars["incidentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AssignIncidentToClaim(claimId, incidentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Incident on a Claim
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignIncidentFromClaim( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.UnassignIncidentFromClaim(claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more exposuresIds as a Exposures to a Claim
	//----------------------------------------------------------------------------
func AddExposuresToClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exposuresIds,_ := vars["exposuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AddExposuresToClaim(claimId, exposuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more exposuresIds as a Exposures from a Claim
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExposuresFromClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exposuresIds,_ := vars["exposuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.RemoveExposuresFromClaim(claimId, exposuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reservesIds as a Reserves to a Claim
	//----------------------------------------------------------------------------
func AddReservesToClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservesIds,_ := vars["reservesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AddReservesToClaim(claimId, reservesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reservesIds as a Reserves from a Claim
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReservesFromClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservesIds,_ := vars["reservesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.RemoveReservesFromClaim(claimId, reservesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more claimPaymentsIds as a ClaimPayments to a Claim
	//----------------------------------------------------------------------------
func AddClaimPaymentsToClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimPaymentsIds,_ := vars["claimPaymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AddClaimPaymentsToClaim(claimId, claimPaymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimPaymentsIds as a ClaimPayments from a Claim
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimPaymentsFromClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimPaymentsIds,_ := vars["claimPaymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.RemoveClaimPaymentsFromClaim(claimId, claimPaymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more serviceProvidersIds as a ServiceProviders to a Claim
	//----------------------------------------------------------------------------
func AddServiceProvidersToClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serviceProvidersIds,_ := vars["serviceProvidersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AddServiceProvidersToClaim(claimId, serviceProvidersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serviceProvidersIds as a ServiceProviders from a Claim
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveServiceProvidersFromClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serviceProvidersIds,_ := vars["serviceProvidersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.RemoveServiceProvidersFromClaim(claimId, serviceProvidersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more subrogationsIds as a Subrogations to a Claim
	//----------------------------------------------------------------------------
func AddSubrogationsToClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	subrogationsIds,_ := vars["subrogationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.AddSubrogationsToClaim(claimId, subrogationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more subrogationsIds as a Subrogations from a Claim
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSubrogationsFromClaim(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	claimId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	subrogationsIds,_ := vars["subrogationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Claim DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimDAO.RemoveSubrogationsFromClaim(claimId, subrogationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
