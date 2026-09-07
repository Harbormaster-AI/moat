package controller

import (
    PaymentProcessorDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PaymentProcessorDAO for database creation
//----------------------------------------------------------------------------
func CreatePaymentProcessor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentProcessor model
	//----------------------------------------------------------------------------
	data := model.PaymentProcessor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentProcessor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor data access object to create
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.CreatePaymentProcessor( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PaymentProcessorDAO to find the relevant PaymentProcessor
//----------------------------------------------------------------------------
func GetPaymentProcessor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentProcessor data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.GetPaymentProcessor(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PaymentProcessorDAO for database read of all PaymentProcessors
//----------------------------------------------------------------------------
func GetAllPaymentProcessor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.GetAllPaymentProcessor()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PaymentProcessorDAO for database save
//----------------------------------------------------------------------------
func UpdatePaymentProcessor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentProcessor model
	//----------------------------------------------------------------------------
	var data = model.PaymentProcessor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentProcessor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.UpdatePaymentProcessor(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PaymentProcessorDAO for database deletion
//----------------------------------------------------------------------------
func DeletePaymentProcessor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentProcessor data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PaymentProcessorDAO.DeletePaymentProcessor(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more institutionsIds as a Institutions to a PaymentProcessor
	//----------------------------------------------------------------------------
func AddInstitutionsToPaymentProcessor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentProcessorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	institutionsIds,_ := vars["institutionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.AddInstitutionsToPaymentProcessor(paymentProcessorId, institutionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more institutionsIds as a Institutions from a PaymentProcessor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInstitutionsFromPaymentProcessor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentProcessorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	institutionsIds,_ := vars["institutionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.RemoveInstitutionsFromPaymentProcessor(paymentProcessorId, institutionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a PaymentProcessor
	//----------------------------------------------------------------------------
func AddContractsToPaymentProcessor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentProcessorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.AddContractsToPaymentProcessor(paymentProcessorId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a PaymentProcessor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromPaymentProcessor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentProcessorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.RemoveContractsFromPaymentProcessor(paymentProcessorId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more settlementsIds as a Settlements to a PaymentProcessor
	//----------------------------------------------------------------------------
func AddSettlementsToPaymentProcessor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentProcessorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	settlementsIds,_ := vars["settlementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.AddSettlementsToPaymentProcessor(paymentProcessorId, settlementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more settlementsIds as a Settlements from a PaymentProcessor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSettlementsFromPaymentProcessor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentProcessorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	settlementsIds,_ := vars["settlementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentProcessor DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentProcessorDAO.RemoveSettlementsFromPaymentProcessor(paymentProcessorId, settlementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
