package controller

import (
    UserDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UserDAO for database creation
//----------------------------------------------------------------------------
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty User model
	//----------------------------------------------------------------------------
	data := model.User{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a User model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the User data access object to create
	//----------------------------------------------------------------------------
	requestResult := UserDAO.CreateUser( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UserDAO to find the relevant User
//----------------------------------------------------------------------------
func GetUser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the User data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UserDAO.GetUser(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UserDAO for database read of all Users
//----------------------------------------------------------------------------
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the User data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UserDAO.GetAllUser()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UserDAO for database save
//----------------------------------------------------------------------------
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty User model
	//----------------------------------------------------------------------------
	var data = model.User{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a User model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the User data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UserDAO.UpdateUser(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UserDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the User data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UserDAO.DeleteUser(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a User
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AssignOrganizationToUser(userId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a User
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromUser( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.UnassignOrganizationFromUser(userId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more teamsIds as a Teams to a User
	//----------------------------------------------------------------------------
func AddTeamsToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddTeamsToUser(userId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more teamsIds as a Teams from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTeamsFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveTeamsFromUser(userId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a User
	//----------------------------------------------------------------------------
func AddActivitiesToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddActivitiesToUser(userId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveActivitiesFromUser(userId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ownedAccountsIds as a OwnedAccounts to a User
	//----------------------------------------------------------------------------
func AddOwnedAccountsToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedAccountsIds,_ := vars["ownedAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddOwnedAccountsToUser(userId, ownedAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownedAccountsIds as a OwnedAccounts from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnedAccountsFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedAccountsIds,_ := vars["ownedAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveOwnedAccountsFromUser(userId, ownedAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ownedLeadsIds as a OwnedLeads to a User
	//----------------------------------------------------------------------------
func AddOwnedLeadsToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedLeadsIds,_ := vars["ownedLeadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddOwnedLeadsToUser(userId, ownedLeadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownedLeadsIds as a OwnedLeads from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnedLeadsFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedLeadsIds,_ := vars["ownedLeadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveOwnedLeadsFromUser(userId, ownedLeadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ownedOpportunitiesIds as a OwnedOpportunities to a User
	//----------------------------------------------------------------------------
func AddOwnedOpportunitiesToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedOpportunitiesIds,_ := vars["ownedOpportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddOwnedOpportunitiesToUser(userId, ownedOpportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownedOpportunitiesIds as a OwnedOpportunities from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnedOpportunitiesFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedOpportunitiesIds,_ := vars["ownedOpportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveOwnedOpportunitiesFromUser(userId, ownedOpportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ownedCasesIds as a OwnedCases to a User
	//----------------------------------------------------------------------------
func AddOwnedCasesToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedCasesIds,_ := vars["ownedCasesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddOwnedCasesToUser(userId, ownedCasesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownedCasesIds as a OwnedCases from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnedCasesFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedCasesIds,_ := vars["ownedCasesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveOwnedCasesFromUser(userId, ownedCasesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more quotesIds as a Quotes to a User
	//----------------------------------------------------------------------------
func AddQuotesToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddQuotesToUser(userId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quotesIds as a Quotes from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuotesFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveQuotesFromUser(userId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a User
	//----------------------------------------------------------------------------
func AddOrdersToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddOrdersToUser(userId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveOrdersFromUser(userId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a User
	//----------------------------------------------------------------------------
func AddContractsToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddContractsToUser(userId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveContractsFromUser(userId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more emailMessagesIds as a EmailMessages to a User
	//----------------------------------------------------------------------------
func AddEmailMessagesToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddEmailMessagesToUser(userId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more emailMessagesIds as a EmailMessages from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmailMessagesFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveEmailMessagesFromUser(userId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
