package controller

import (
    BillingAccountDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BillingAccountDAO for database creation
//----------------------------------------------------------------------------
func CreateBillingAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BillingAccount model
	//----------------------------------------------------------------------------
	data := model.BillingAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BillingAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount data access object to create
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.CreateBillingAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BillingAccountDAO to find the relevant BillingAccount
//----------------------------------------------------------------------------
func GetBillingAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BillingAccount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.GetBillingAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BillingAccountDAO for database read of all BillingAccounts
//----------------------------------------------------------------------------
func GetAllBillingAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.GetAllBillingAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BillingAccountDAO for database save
//----------------------------------------------------------------------------
func UpdateBillingAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BillingAccount model
	//----------------------------------------------------------------------------
	var data = model.BillingAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BillingAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.UpdateBillingAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BillingAccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBillingAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BillingAccount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BillingAccountDAO.DeleteBillingAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a BillingAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToBillingAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.AssignCustomerToBillingAccount(billingAccountId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a BillingAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromBillingAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.UnassignCustomerFromBillingAccount(billingAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a BillingAccount
	//----------------------------------------------------------------------------
func AddPoliciesToBillingAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.AddPoliciesToBillingAccount(billingAccountId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a BillingAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromBillingAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.RemovePoliciesFromBillingAccount(billingAccountId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more invoicesIds as a Invoices to a BillingAccount
	//----------------------------------------------------------------------------
func AddInvoicesToBillingAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoicesIds,_ := vars["invoicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.AddInvoicesToBillingAccount(billingAccountId, invoicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more invoicesIds as a Invoices from a BillingAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInvoicesFromBillingAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoicesIds,_ := vars["invoicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.RemoveInvoicesFromBillingAccount(billingAccountId, invoicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentsIds as a Payments to a BillingAccount
	//----------------------------------------------------------------------------
func AddPaymentsToBillingAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.AddPaymentsToBillingAccount(billingAccountId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentsIds as a Payments from a BillingAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentsFromBillingAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingAccount DAO
	//----------------------------------------------------------------------------
	requestResult := BillingAccountDAO.RemovePaymentsFromBillingAccount(billingAccountId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
