package controller

import (
    AccountDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AccountDAO for database creation
//----------------------------------------------------------------------------
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Account model
	//----------------------------------------------------------------------------
	data := model.Account{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Account model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Account data access object to create
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.CreateAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AccountDAO to find the relevant Account
//----------------------------------------------------------------------------
func GetAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Account data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.GetAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AccountDAO for database read of all Accounts
//----------------------------------------------------------------------------
func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Account data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.GetAllAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AccountDAO for database save
//----------------------------------------------------------------------------
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Account model
	//----------------------------------------------------------------------------
	var data = model.Account{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Account model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Account data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UpdateAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Account data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AccountDAO.DeleteAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignCustomerToAccount(accountId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignCustomerFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Institution on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInstitutionToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	institutionId,_ := strconv.ParseUint( vars["institutionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignInstitutionToAccount(accountId, institutionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Institution on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInstitutionFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignInstitutionFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a Account
	//----------------------------------------------------------------------------
func AddTransactionsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddTransactionsToAccount(accountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveTransactionsFromAccount(accountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more cardsIds as a Cards to a Account
	//----------------------------------------------------------------------------
func AddCardsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardsIds,_ := vars["cardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddCardsToAccount(accountId, cardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more cardsIds as a Cards from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCardsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardsIds,_ := vars["cardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveCardsFromAccount(accountId, cardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more statementsIds as a Statements to a Account
	//----------------------------------------------------------------------------
func AddStatementsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	statementsIds,_ := vars["statementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddStatementsToAccount(accountId, statementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more statementsIds as a Statements from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStatementsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	statementsIds,_ := vars["statementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveStatementsFromAccount(accountId, statementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more mandatesIds as a Mandates to a Account
	//----------------------------------------------------------------------------
func AddMandatesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mandatesIds,_ := vars["mandatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddMandatesToAccount(accountId, mandatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more mandatesIds as a Mandates from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMandatesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mandatesIds,_ := vars["mandatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveMandatesFromAccount(accountId, mandatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
