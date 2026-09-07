package controller

import (
    OpportunityDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OpportunityDAO for database creation
//----------------------------------------------------------------------------
func CreateOpportunity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Opportunity model
	//----------------------------------------------------------------------------
	data := model.Opportunity{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Opportunity model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity data access object to create
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.CreateOpportunity( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OpportunityDAO to find the relevant Opportunity
//----------------------------------------------------------------------------
func GetOpportunity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Opportunity data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.GetOpportunity(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OpportunityDAO for database read of all Opportunitys
//----------------------------------------------------------------------------
func GetAllOpportunity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Opportunity data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.GetAllOpportunity()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OpportunityDAO for database save
//----------------------------------------------------------------------------
func UpdateOpportunity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Opportunity model
	//----------------------------------------------------------------------------
	var data = model.Opportunity{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Opportunity model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.UpdateOpportunity(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OpportunityDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOpportunity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Opportunity data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OpportunityDAO.DeleteOpportunity(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Opportunity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToOpportunity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AssignOrganizationToOpportunity(opportunityId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Opportunity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromOpportunity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.UnassignOrganizationFromOpportunity(opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Opportunity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToOpportunity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AssignAccountToOpportunity(opportunityId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Opportunity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromOpportunity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.UnassignAccountFromOpportunity(opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Opportunity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToOpportunity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AssignOwnerToOpportunity(opportunityId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Opportunity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromOpportunity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.UnassignOwnerFromOpportunity(opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more contactsIds as a Contacts to a Opportunity
	//----------------------------------------------------------------------------
func AddContactsToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactsIds,_ := vars["contactsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddContactsToOpportunity(opportunityId, contactsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contactsIds as a Contacts from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContactsFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactsIds,_ := vars["contactsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveContactsFromOpportunity(opportunityId, contactsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more lineItemsIds as a LineItems to a Opportunity
	//----------------------------------------------------------------------------
func AddLineItemsToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemsIds,_ := vars["lineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddLineItemsToOpportunity(opportunityId, lineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more lineItemsIds as a LineItems from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLineItemsFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemsIds,_ := vars["lineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveLineItemsFromOpportunity(opportunityId, lineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more stageHistoryIds as a StageHistory to a Opportunity
	//----------------------------------------------------------------------------
func AddStageHistoryToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	stageHistoryIds,_ := vars["stageHistoryIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddStageHistoryToOpportunity(opportunityId, stageHistoryIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more stageHistoryIds as a StageHistory from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStageHistoryFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	stageHistoryIds,_ := vars["stageHistoryIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveStageHistoryFromOpportunity(opportunityId, stageHistoryIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more quotesIds as a Quotes to a Opportunity
	//----------------------------------------------------------------------------
func AddQuotesToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddQuotesToOpportunity(opportunityId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quotesIds as a Quotes from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuotesFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveQuotesFromOpportunity(opportunityId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a Opportunity
	//----------------------------------------------------------------------------
func AddOrdersToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddOrdersToOpportunity(opportunityId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveOrdersFromOpportunity(opportunityId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a Opportunity
	//----------------------------------------------------------------------------
func AddCampaignsToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddCampaignsToOpportunity(opportunityId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveCampaignsFromOpportunity(opportunityId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a Opportunity
	//----------------------------------------------------------------------------
func AddActivitiesToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddActivitiesToOpportunity(opportunityId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveActivitiesFromOpportunity(opportunityId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more teamsIds as a Teams to a Opportunity
	//----------------------------------------------------------------------------
func AddTeamsToOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.AddTeamsToOpportunity(opportunityId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more teamsIds as a Teams from a Opportunity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTeamsFromOpportunity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	opportunityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Opportunity DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityDAO.RemoveTeamsFromOpportunity(opportunityId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
