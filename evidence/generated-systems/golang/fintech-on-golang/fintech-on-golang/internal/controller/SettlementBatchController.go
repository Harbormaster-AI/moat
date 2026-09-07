package controller

import (
    SettlementBatchDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SettlementBatchDAO for database creation
//----------------------------------------------------------------------------
func CreateSettlementBatch(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SettlementBatch model
	//----------------------------------------------------------------------------
	data := model.SettlementBatch{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SettlementBatch model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch data access object to create
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.CreateSettlementBatch( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SettlementBatchDAO to find the relevant SettlementBatch
//----------------------------------------------------------------------------
func GetSettlementBatch(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SettlementBatch data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.GetSettlementBatch(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SettlementBatchDAO for database read of all SettlementBatchs
//----------------------------------------------------------------------------
func GetAllSettlementBatch(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.GetAllSettlementBatch()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SettlementBatchDAO for database save
//----------------------------------------------------------------------------
func UpdateSettlementBatch(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SettlementBatch model
	//----------------------------------------------------------------------------
	var data = model.SettlementBatch{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SettlementBatch model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.UpdateSettlementBatch(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SettlementBatchDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSettlementBatch(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SettlementBatch data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SettlementBatchDAO.DeleteSettlementBatch(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Processor on a SettlementBatch
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProcessorToSettlementBatch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processorId,_ := strconv.ParseUint( vars["processorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.AssignProcessorToSettlementBatch(settlementBatchId, processorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Processor on a SettlementBatch
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProcessorFromSettlementBatch( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.UnassignProcessorFromSettlementBatch(settlementBatchId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Merchant on a SettlementBatch
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMerchantToSettlementBatch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	merchantId,_ := strconv.ParseUint( vars["merchantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.AssignMerchantToSettlementBatch(settlementBatchId, merchantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Merchant on a SettlementBatch
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMerchantFromSettlementBatch( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.UnassignMerchantFromSettlementBatch(settlementBatchId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more payoutsIds as a Payouts to a SettlementBatch
	//----------------------------------------------------------------------------
func AddPayoutsToSettlementBatch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payoutsIds,_ := vars["payoutsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.AddPayoutsToSettlementBatch(settlementBatchId, payoutsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more payoutsIds as a Payouts from a SettlementBatch
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePayoutsFromSettlementBatch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payoutsIds,_ := vars["payoutsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.RemovePayoutsFromSettlementBatch(settlementBatchId, payoutsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a SettlementBatch
	//----------------------------------------------------------------------------
func AddTransactionsToSettlementBatch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.AddTransactionsToSettlementBatch(settlementBatchId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a SettlementBatch
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromSettlementBatch(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	settlementBatchId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SettlementBatch DAO
	//----------------------------------------------------------------------------
	requestResult := SettlementBatchDAO.RemoveTransactionsFromSettlementBatch(settlementBatchId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
