package controller

import (
    TransactionDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TransactionDAO for database creation
//----------------------------------------------------------------------------
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Transaction model
	//----------------------------------------------------------------------------
	data := model.Transaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Transaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction data access object to create
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.CreateTransaction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TransactionDAO to find the relevant Transaction
//----------------------------------------------------------------------------
func GetTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Transaction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.GetTransaction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TransactionDAO for database read of all Transactions
//----------------------------------------------------------------------------
func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Transaction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.GetAllTransaction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TransactionDAO for database save
//----------------------------------------------------------------------------
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Transaction model
	//----------------------------------------------------------------------------
	var data = model.Transaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Transaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UpdateTransaction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TransactionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Transaction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TransactionDAO.DeleteTransaction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Account on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignAccountToTransaction(transactionId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignAccountFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Wallet on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWalletToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	walletId,_ := strconv.ParseUint( vars["walletId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignWalletToTransaction(transactionId, walletId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Wallet on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWalletFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignWalletFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PaymentOrder on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPaymentOrderToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentOrderId,_ := strconv.ParseUint( vars["paymentOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignPaymentOrderToTransaction(transactionId, paymentOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PaymentOrder on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPaymentOrderFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignPaymentOrderFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Merchant on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMerchantToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	merchantId,_ := strconv.ParseUint( vars["merchantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignMerchantToTransaction(transactionId, merchantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Merchant on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMerchantFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignMerchantFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Card on a Transaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCardToTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardId,_ := strconv.ParseUint( vars["cardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AssignCardToTransaction(transactionId, cardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Card on a Transaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCardFromTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.UnassignCardFromTransaction(transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more relatedTransactionsIds as a RelatedTransactions to a Transaction
	//----------------------------------------------------------------------------
func AddRelatedTransactionsToTransaction(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedTransactionsIds,_ := vars["relatedTransactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AddRelatedTransactionsToTransaction(transactionId, relatedTransactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more relatedTransactionsIds as a RelatedTransactions from a Transaction
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRelatedTransactionsFromTransaction(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedTransactionsIds,_ := vars["relatedTransactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.RemoveRelatedTransactionsFromTransaction(transactionId, relatedTransactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a Transaction
	//----------------------------------------------------------------------------
func AddAlertsToTransaction(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.AddAlertsToTransaction(transactionId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a Transaction
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromTransaction(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Transaction DAO
	//----------------------------------------------------------------------------
	requestResult := TransactionDAO.RemoveAlertsFromTransaction(transactionId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
