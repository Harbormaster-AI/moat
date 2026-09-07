package controller

import (
    ClaimPaymentDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ClaimPaymentDAO for database creation
//----------------------------------------------------------------------------
func CreateClaimPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ClaimPayment model
	//----------------------------------------------------------------------------
	data := model.ClaimPayment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ClaimPayment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment data access object to create
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.CreateClaimPayment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ClaimPaymentDAO to find the relevant ClaimPayment
//----------------------------------------------------------------------------
func GetClaimPayment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ClaimPayment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.GetClaimPayment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ClaimPaymentDAO for database read of all ClaimPayments
//----------------------------------------------------------------------------
func GetAllClaimPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.GetAllClaimPayment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ClaimPaymentDAO for database save
//----------------------------------------------------------------------------
func UpdateClaimPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ClaimPayment model
	//----------------------------------------------------------------------------
	var data = model.ClaimPayment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ClaimPayment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.UpdateClaimPayment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ClaimPaymentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteClaimPayment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ClaimPayment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ClaimPaymentDAO.DeleteClaimPayment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Claim on a ClaimPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClaimToClaimPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimId,_ := strconv.ParseUint( vars["claimId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.AssignClaimToClaimPayment(claimPaymentId, claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Claim on a ClaimPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClaimFromClaimPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.UnassignClaimFromClaimPayment(claimPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Exposure on a ClaimPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExposureToClaimPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exposureId,_ := strconv.ParseUint( vars["exposureId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.AssignExposureToClaimPayment(claimPaymentId, exposureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Exposure on a ClaimPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExposureFromClaimPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.UnassignExposureFromClaimPayment(claimPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Beneficiary on a ClaimPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBeneficiaryToClaimPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	beneficiaryId,_ := strconv.ParseUint( vars["beneficiaryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.AssignBeneficiaryToClaimPayment(claimPaymentId, beneficiaryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Beneficiary on a ClaimPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBeneficiaryFromClaimPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.UnassignBeneficiaryFromClaimPayment(claimPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ServiceProvider on a ClaimPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignServiceProviderToClaimPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serviceProviderId,_ := strconv.ParseUint( vars["serviceProviderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.AssignServiceProviderToClaimPayment(claimPaymentId, serviceProviderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ServiceProvider on a ClaimPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignServiceProviderFromClaimPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.UnassignServiceProviderFromClaimPayment(claimPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Customer on a ClaimPayment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToClaimPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.AssignCustomerToClaimPayment(claimPaymentId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a ClaimPayment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromClaimPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimPaymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimPayment DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimPaymentDAO.UnassignCustomerFromClaimPayment(claimPaymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


