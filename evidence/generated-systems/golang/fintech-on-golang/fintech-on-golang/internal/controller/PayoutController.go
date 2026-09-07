package controller

import (
    PayoutDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PayoutDAO for database creation
//----------------------------------------------------------------------------
func CreatePayout(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Payout model
	//----------------------------------------------------------------------------
	data := model.Payout{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Payout model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Payout data access object to create
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.CreatePayout( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PayoutDAO to find the relevant Payout
//----------------------------------------------------------------------------
func GetPayout(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Payout data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.GetPayout(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PayoutDAO for database read of all Payouts
//----------------------------------------------------------------------------
func GetAllPayout(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Payout data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.GetAllPayout()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PayoutDAO for database save
//----------------------------------------------------------------------------
func UpdatePayout(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Payout model
	//----------------------------------------------------------------------------
	var data = model.Payout{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Payout model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Payout data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.UpdatePayout(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PayoutDAO for database deletion
//----------------------------------------------------------------------------
func DeletePayout(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Payout data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PayoutDAO.DeletePayout(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Merchant on a Payout
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMerchantToPayout(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	merchantId,_ := strconv.ParseUint( vars["merchantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payout DAO
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.AssignMerchantToPayout(payoutId, merchantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Merchant on a Payout
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMerchantFromPayout( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payout DAO
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.UnassignMerchantFromPayout(payoutId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SettlementBatch on a Payout
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSettlementBatchToPayout(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	settlementBatchId,_ := strconv.ParseUint( vars["settlementBatchId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payout DAO
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.AssignSettlementBatchToPayout(payoutId, settlementBatchId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SettlementBatch on a Payout
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSettlementBatchFromPayout( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payout DAO
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.UnassignSettlementBatchFromPayout(payoutId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DestinationAccount on a Payout
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDestinationAccountToPayout(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationAccountId,_ := strconv.ParseUint( vars["destinationAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payout DAO
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.AssignDestinationAccountToPayout(payoutId, destinationAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DestinationAccount on a Payout
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDestinationAccountFromPayout( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payout DAO
	//----------------------------------------------------------------------------
	requestResult := PayoutDAO.UnassignDestinationAccountFromPayout(payoutId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


