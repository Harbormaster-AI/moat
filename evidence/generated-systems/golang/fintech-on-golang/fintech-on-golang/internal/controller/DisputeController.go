package controller

import (
    DisputeDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DisputeDAO for database creation
//----------------------------------------------------------------------------
func CreateDispute(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dispute model
	//----------------------------------------------------------------------------
	data := model.Dispute{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dispute model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute data access object to create
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.CreateDispute( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DisputeDAO to find the relevant Dispute
//----------------------------------------------------------------------------
func GetDispute(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dispute data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.GetDispute(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DisputeDAO for database read of all Disputes
//----------------------------------------------------------------------------
func GetAllDispute(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Dispute data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.GetAllDispute()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DisputeDAO for database save
//----------------------------------------------------------------------------
func UpdateDispute(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dispute model
	//----------------------------------------------------------------------------
	var data = model.Dispute{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dispute model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UpdateDispute(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DisputeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDispute(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dispute data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DisputeDAO.DeleteDispute(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignTransactionToDispute(disputeId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignTransactionFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Card on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCardToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardId,_ := strconv.ParseUint( vars["cardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignCardToDispute(disputeId, cardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Card on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCardFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignCardFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Merchant on a Dispute
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMerchantToDispute(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	merchantId,_ := strconv.ParseUint( vars["merchantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AssignMerchantToDispute(disputeId, merchantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Merchant on a Dispute
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMerchantFromDispute( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.UnassignMerchantFromDispute(disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more chargebacksIds as a Chargebacks to a Dispute
	//----------------------------------------------------------------------------
func AddChargebacksToDispute(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	chargebacksIds,_ := vars["chargebacksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.AddChargebacksToDispute(disputeId, chargebacksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more chargebacksIds as a Chargebacks from a Dispute
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChargebacksFromDispute(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	disputeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	chargebacksIds,_ := vars["chargebacksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dispute DAO
	//----------------------------------------------------------------------------
	requestResult := DisputeDAO.RemoveChargebacksFromDispute(disputeId, chargebacksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
