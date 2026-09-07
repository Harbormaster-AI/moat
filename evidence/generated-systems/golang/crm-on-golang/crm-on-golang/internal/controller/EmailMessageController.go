package controller

import (
    EmailMessageDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EmailMessageDAO for database creation
//----------------------------------------------------------------------------
func CreateEmailMessage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EmailMessage model
	//----------------------------------------------------------------------------
	data := model.EmailMessage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EmailMessage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage data access object to create
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.CreateEmailMessage( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EmailMessageDAO to find the relevant EmailMessage
//----------------------------------------------------------------------------
func GetEmailMessage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EmailMessage data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.GetEmailMessage(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EmailMessageDAO for database read of all EmailMessages
//----------------------------------------------------------------------------
func GetAllEmailMessage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.GetAllEmailMessage()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EmailMessageDAO for database save
//----------------------------------------------------------------------------
func UpdateEmailMessage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EmailMessage model
	//----------------------------------------------------------------------------
	var data = model.EmailMessage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EmailMessage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UpdateEmailMessage(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EmailMessageDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEmailMessage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EmailMessage data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EmailMessageDAO.DeleteEmailMessage(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignOrganizationToEmailMessage(emailMessageId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignOrganizationFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignOwnerToEmailMessage(emailMessageId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignOwnerFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignAccountToEmailMessage(emailMessageId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignAccountFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contact on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContactToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactId,_ := strconv.ParseUint( vars["contactId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignContactToEmailMessage(emailMessageId, contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contact on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContactFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignContactFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lead on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLeadToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leadId,_ := strconv.ParseUint( vars["leadId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignLeadToEmailMessage(emailMessageId, leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lead on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLeadFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignLeadFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Case on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCaseToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	caseId,_ := strconv.ParseUint( vars["caseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignCaseToEmailMessage(emailMessageId, caseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Case on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCaseFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignCaseFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Opportunity on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityId,_ := strconv.ParseUint( vars["opportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignOpportunityToEmailMessage(emailMessageId, opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Opportunity on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignOpportunityFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a EmailMessage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToEmailMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.AssignCampaignToEmailMessage(emailMessageId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a EmailMessage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromEmailMessage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	emailMessageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmailMessage DAO
	//----------------------------------------------------------------------------
	requestResult := EmailMessageDAO.UnassignCampaignFromEmailMessage(emailMessageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


