package controller

import (
    PaymentOrderDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PaymentOrderDAO for database creation
//----------------------------------------------------------------------------
func CreatePaymentOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentOrder model
	//----------------------------------------------------------------------------
	data := model.PaymentOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.CreatePaymentOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PaymentOrderDAO to find the relevant PaymentOrder
//----------------------------------------------------------------------------
func GetPaymentOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.GetPaymentOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PaymentOrderDAO for database read of all PaymentOrders
//----------------------------------------------------------------------------
func GetAllPaymentOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.GetAllPaymentOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PaymentOrderDAO for database save
//----------------------------------------------------------------------------
func UpdatePaymentOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentOrder model
	//----------------------------------------------------------------------------
	var data = model.PaymentOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.UpdatePaymentOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PaymentOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeletePaymentOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PaymentOrderDAO.DeletePaymentOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SourceAccount on a PaymentOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSourceAccountToPaymentOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourceAccountId,_ := strconv.ParseUint( vars["sourceAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.AssignSourceAccountToPaymentOrder(paymentOrderId, sourceAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SourceAccount on a PaymentOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSourceAccountFromPaymentOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.UnassignSourceAccountFromPaymentOrder(paymentOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DestinationAccount on a PaymentOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDestinationAccountToPaymentOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationAccountId,_ := strconv.ParseUint( vars["destinationAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.AssignDestinationAccountToPaymentOrder(paymentOrderId, destinationAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DestinationAccount on a PaymentOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDestinationAccountFromPaymentOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.UnassignDestinationAccountFromPaymentOrder(paymentOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Beneficiary on a PaymentOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBeneficiaryToPaymentOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	beneficiaryId,_ := strconv.ParseUint( vars["beneficiaryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.AssignBeneficiaryToPaymentOrder(paymentOrderId, beneficiaryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Beneficiary on a PaymentOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBeneficiaryFromPaymentOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.UnassignBeneficiaryFromPaymentOrder(paymentOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a FxDeal on a PaymentOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFxDealToPaymentOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fxDealId,_ := strconv.ParseUint( vars["fxDealId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.AssignFxDealToPaymentOrder(paymentOrderId, fxDealId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a FxDeal on a PaymentOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFxDealFromPaymentOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.UnassignFxDealFromPaymentOrder(paymentOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a PaymentOrder
	//----------------------------------------------------------------------------
func AddTransactionsToPaymentOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.AddTransactionsToPaymentOrder(paymentOrderId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a PaymentOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromPaymentOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.RemoveTransactionsFromPaymentOrder(paymentOrderId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more feesIds as a Fees to a PaymentOrder
	//----------------------------------------------------------------------------
func AddFeesToPaymentOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feesIds,_ := vars["feesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.AddFeesToPaymentOrder(paymentOrderId, feesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more feesIds as a Fees from a PaymentOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeesFromPaymentOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feesIds,_ := vars["feesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentOrderDAO.RemoveFeesFromPaymentOrder(paymentOrderId, feesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
