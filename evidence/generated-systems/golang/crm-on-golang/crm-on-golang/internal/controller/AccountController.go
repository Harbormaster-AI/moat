package controller

import (
    AccountDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
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
	// assigns a Organization on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignOrganizationToAccount(accountId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignOrganizationFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ParentAccount on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignParentAccountToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	parentAccountId,_ := strconv.ParseUint( vars["parentAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignParentAccountToAccount(accountId, parentAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ParentAccount on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignParentAccountFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignParentAccountFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignOwnerToAccount(accountId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignOwnerFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Territory on a Account
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTerritoryToAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	territoryId,_ := strconv.ParseUint( vars["territoryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AssignTerritoryToAccount(accountId, territoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Territory on a Account
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTerritoryFromAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.UnassignTerritoryFromAccount(accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more childAccountsIds as a ChildAccounts to a Account
	//----------------------------------------------------------------------------
func AddChildAccountsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childAccountsIds,_ := vars["childAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddChildAccountsToAccount(accountId, childAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more childAccountsIds as a ChildAccounts from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChildAccountsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childAccountsIds,_ := vars["childAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveChildAccountsFromAccount(accountId, childAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contactsIds as a Contacts to a Account
	//----------------------------------------------------------------------------
func AddContactsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactsIds,_ := vars["contactsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddContactsToAccount(accountId, contactsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contactsIds as a Contacts from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContactsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactsIds,_ := vars["contactsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveContactsFromAccount(accountId, contactsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more opportunitiesIds as a Opportunities to a Account
	//----------------------------------------------------------------------------
func AddOpportunitiesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunitiesIds,_ := vars["opportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddOpportunitiesToAccount(accountId, opportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more opportunitiesIds as a Opportunities from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOpportunitiesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunitiesIds,_ := vars["opportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveOpportunitiesFromAccount(accountId, opportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more casesIds as a Cases to a Account
	//----------------------------------------------------------------------------
func AddCasesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	casesIds,_ := vars["casesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddCasesToAccount(accountId, casesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more casesIds as a Cases from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCasesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	casesIds,_ := vars["casesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveCasesFromAccount(accountId, casesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a Account
	//----------------------------------------------------------------------------
func AddActivitiesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddActivitiesToAccount(accountId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveActivitiesFromAccount(accountId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a Account
	//----------------------------------------------------------------------------
func AddCampaignsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddCampaignsToAccount(accountId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveCampaignsFromAccount(accountId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more quotesIds as a Quotes to a Account
	//----------------------------------------------------------------------------
func AddQuotesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddQuotesToAccount(accountId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quotesIds as a Quotes from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuotesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveQuotesFromAccount(accountId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a Account
	//----------------------------------------------------------------------------
func AddOrdersToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddOrdersToAccount(accountId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveOrdersFromAccount(accountId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a Account
	//----------------------------------------------------------------------------
func AddContractsToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddContractsToAccount(accountId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveContractsFromAccount(accountId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more notesIds as a Notes to a Account
	//----------------------------------------------------------------------------
func AddNotesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notesIds,_ := vars["notesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddNotesToAccount(accountId, notesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more notesIds as a Notes from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNotesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notesIds,_ := vars["notesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveNotesFromAccount(accountId, notesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more emailMessagesIds as a EmailMessages to a Account
	//----------------------------------------------------------------------------
func AddEmailMessagesToAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.AddEmailMessagesToAccount(accountId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more emailMessagesIds as a EmailMessages from a Account
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmailMessagesFromAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Account DAO
	//----------------------------------------------------------------------------
	requestResult := AccountDAO.RemoveEmailMessagesFromAccount(accountId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
