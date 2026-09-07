package controller

import (
    MerchantDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MerchantDAO for database creation
//----------------------------------------------------------------------------
func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Merchant model
	//----------------------------------------------------------------------------
	data := model.Merchant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Merchant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Merchant data access object to create
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.CreateMerchant( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MerchantDAO to find the relevant Merchant
//----------------------------------------------------------------------------
func GetMerchant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Merchant data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.GetMerchant(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MerchantDAO for database read of all Merchants
//----------------------------------------------------------------------------
func GetAllMerchant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Merchant data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.GetAllMerchant()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MerchantDAO for database save
//----------------------------------------------------------------------------
func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Merchant model
	//----------------------------------------------------------------------------
	var data = model.Merchant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Merchant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Merchant data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.UpdateMerchant(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MerchantDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Merchant data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MerchantDAO.DeleteMerchant(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more terminalsIds as a Terminals to a Merchant
	//----------------------------------------------------------------------------
func AddTerminalsToMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	terminalsIds,_ := vars["terminalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.AddTerminalsToMerchant(merchantId, terminalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more terminalsIds as a Terminals from a Merchant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTerminalsFromMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	terminalsIds,_ := vars["terminalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.RemoveTerminalsFromMerchant(merchantId, terminalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentContractsIds as a PaymentContracts to a Merchant
	//----------------------------------------------------------------------------
func AddPaymentContractsToMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentContractsIds,_ := vars["paymentContractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.AddPaymentContractsToMerchant(merchantId, paymentContractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentContractsIds as a PaymentContracts from a Merchant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentContractsFromMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentContractsIds,_ := vars["paymentContractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.RemovePaymentContractsFromMerchant(merchantId, paymentContractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more payoutsIds as a Payouts to a Merchant
	//----------------------------------------------------------------------------
func AddPayoutsToMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payoutsIds,_ := vars["payoutsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.AddPayoutsToMerchant(merchantId, payoutsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more payoutsIds as a Payouts from a Merchant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePayoutsFromMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payoutsIds,_ := vars["payoutsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.RemovePayoutsFromMerchant(merchantId, payoutsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more settlementsIds as a Settlements to a Merchant
	//----------------------------------------------------------------------------
func AddSettlementsToMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	settlementsIds,_ := vars["settlementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.AddSettlementsToMerchant(merchantId, settlementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more settlementsIds as a Settlements from a Merchant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSettlementsFromMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	settlementsIds,_ := vars["settlementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.RemoveSettlementsFromMerchant(merchantId, settlementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more disputesIds as a Disputes to a Merchant
	//----------------------------------------------------------------------------
func AddDisputesToMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.AddDisputesToMerchant(merchantId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more disputesIds as a Disputes from a Merchant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDisputesFromMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.RemoveDisputesFromMerchant(merchantId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more invoicesIds as a Invoices to a Merchant
	//----------------------------------------------------------------------------
func AddInvoicesToMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoicesIds,_ := vars["invoicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.AddInvoicesToMerchant(merchantId, invoicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more invoicesIds as a Invoices from a Merchant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInvoicesFromMerchant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	merchantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoicesIds,_ := vars["invoicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Merchant DAO
	//----------------------------------------------------------------------------
	requestResult := MerchantDAO.RemoveInvoicesFromMerchant(merchantId, invoicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
