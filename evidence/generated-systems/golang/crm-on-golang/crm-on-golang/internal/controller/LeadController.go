package controller

import (
    LeadDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LeadDAO for database creation
//----------------------------------------------------------------------------
func CreateLead(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Lead model
	//----------------------------------------------------------------------------
	data := model.Lead{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Lead model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Lead data access object to create
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.CreateLead( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LeadDAO to find the relevant Lead
//----------------------------------------------------------------------------
func GetLead(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Lead data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.GetLead(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LeadDAO for database read of all Leads
//----------------------------------------------------------------------------
func GetAllLead(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Lead data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.GetAllLead()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LeadDAO for database save
//----------------------------------------------------------------------------
func UpdateLead(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Lead model
	//----------------------------------------------------------------------------
	var data = model.Lead{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Lead model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Lead data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.UpdateLead(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LeadDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLead(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Lead data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LeadDAO.DeleteLead(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Lead
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToLead(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AssignOrganizationToLead(leadId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Lead
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromLead( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.UnassignOrganizationFromLead(leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Lead
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToLead(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AssignOwnerToLead(leadId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Lead
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromLead( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.UnassignOwnerFromLead(leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ConvertedAccount on a Lead
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConvertedAccountToLead(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	convertedAccountId,_ := strconv.ParseUint( vars["convertedAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AssignConvertedAccountToLead(leadId, convertedAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConvertedAccount on a Lead
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConvertedAccountFromLead( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.UnassignConvertedAccountFromLead(leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ConvertedContact on a Lead
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConvertedContactToLead(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	convertedContactId,_ := strconv.ParseUint( vars["convertedContactId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AssignConvertedContactToLead(leadId, convertedContactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConvertedContact on a Lead
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConvertedContactFromLead( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.UnassignConvertedContactFromLead(leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ConvertedOpportunity on a Lead
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConvertedOpportunityToLead(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	convertedOpportunityId,_ := strconv.ParseUint( vars["convertedOpportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AssignConvertedOpportunityToLead(leadId, convertedOpportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConvertedOpportunity on a Lead
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConvertedOpportunityFromLead( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.UnassignConvertedOpportunityFromLead(leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a Lead
	//----------------------------------------------------------------------------
func AddActivitiesToLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AddActivitiesToLead(leadId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a Lead
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.RemoveActivitiesFromLead(leadId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a Lead
	//----------------------------------------------------------------------------
func AddCampaignsToLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AddCampaignsToLead(leadId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a Lead
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.RemoveCampaignsFromLead(leadId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more notesIds as a Notes to a Lead
	//----------------------------------------------------------------------------
func AddNotesToLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notesIds,_ := vars["notesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AddNotesToLead(leadId, notesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more notesIds as a Notes from a Lead
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNotesFromLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notesIds,_ := vars["notesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.RemoveNotesFromLead(leadId, notesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more emailMessagesIds as a EmailMessages to a Lead
	//----------------------------------------------------------------------------
func AddEmailMessagesToLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.AddEmailMessagesToLead(leadId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more emailMessagesIds as a EmailMessages from a Lead
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmailMessagesFromLead(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lead DAO
	//----------------------------------------------------------------------------
	requestResult := LeadDAO.RemoveEmailMessagesFromLead(leadId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
