package controller

import (
    CampaignDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CampaignDAO for database creation
//----------------------------------------------------------------------------
func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Campaign model
	//----------------------------------------------------------------------------
	data := model.Campaign{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Campaign model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign data access object to create
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.CreateCampaign( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CampaignDAO to find the relevant Campaign
//----------------------------------------------------------------------------
func GetCampaign(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Campaign data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.GetCampaign(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CampaignDAO for database read of all Campaigns
//----------------------------------------------------------------------------
func GetAllCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Campaign data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.GetAllCampaign()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CampaignDAO for database save
//----------------------------------------------------------------------------
func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Campaign model
	//----------------------------------------------------------------------------
	var data = model.Campaign{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Campaign model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.UpdateCampaign(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CampaignDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Campaign data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CampaignDAO.DeleteCampaign(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Campaign
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AssignOrganizationToCampaign(campaignId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Campaign
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromCampaign( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.UnassignOrganizationFromCampaign(campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ParentCampaign on a Campaign
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignParentCampaignToCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	parentCampaignId,_ := strconv.ParseUint( vars["parentCampaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AssignParentCampaignToCampaign(campaignId, parentCampaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ParentCampaign on a Campaign
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignParentCampaignFromCampaign( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.UnassignParentCampaignFromCampaign(campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more childCampaignsIds as a ChildCampaigns to a Campaign
	//----------------------------------------------------------------------------
func AddChildCampaignsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childCampaignsIds,_ := vars["childCampaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddChildCampaignsToCampaign(campaignId, childCampaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more childCampaignsIds as a ChildCampaigns from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChildCampaignsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childCampaignsIds,_ := vars["childCampaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveChildCampaignsFromCampaign(campaignId, childCampaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more membersIds as a Members to a Campaign
	//----------------------------------------------------------------------------
func AddMembersToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	membersIds,_ := vars["membersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddMembersToCampaign(campaignId, membersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more membersIds as a Members from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMembersFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	membersIds,_ := vars["membersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveMembersFromCampaign(campaignId, membersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more opportunitiesIds as a Opportunities to a Campaign
	//----------------------------------------------------------------------------
func AddOpportunitiesToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunitiesIds,_ := vars["opportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddOpportunitiesToCampaign(campaignId, opportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more opportunitiesIds as a Opportunities from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOpportunitiesFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunitiesIds,_ := vars["opportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveOpportunitiesFromCampaign(campaignId, opportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Campaign
	//----------------------------------------------------------------------------
func AddAccountsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddAccountsToCampaign(campaignId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveAccountsFromCampaign(campaignId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more leadsIds as a Leads to a Campaign
	//----------------------------------------------------------------------------
func AddLeadsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leadsIds,_ := vars["leadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddLeadsToCampaign(campaignId, leadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more leadsIds as a Leads from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLeadsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leadsIds,_ := vars["leadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveLeadsFromCampaign(campaignId, leadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contactsIds as a Contacts to a Campaign
	//----------------------------------------------------------------------------
func AddContactsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactsIds,_ := vars["contactsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddContactsToCampaign(campaignId, contactsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contactsIds as a Contacts from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContactsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactsIds,_ := vars["contactsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveContactsFromCampaign(campaignId, contactsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more teamsIds as a Teams to a Campaign
	//----------------------------------------------------------------------------
func AddTeamsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddTeamsToCampaign(campaignId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more teamsIds as a Teams from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTeamsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveTeamsFromCampaign(campaignId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a Campaign
	//----------------------------------------------------------------------------
func AddActivitiesToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddActivitiesToCampaign(campaignId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveActivitiesFromCampaign(campaignId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
