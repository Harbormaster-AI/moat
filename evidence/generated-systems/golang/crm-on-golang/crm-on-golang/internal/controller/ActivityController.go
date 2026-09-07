package controller

import (
    ActivityDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ActivityDAO for database creation
//----------------------------------------------------------------------------
func CreateActivity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Activity model
	//----------------------------------------------------------------------------
	data := model.Activity{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Activity model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Activity data access object to create
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.CreateActivity( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ActivityDAO to find the relevant Activity
//----------------------------------------------------------------------------
func GetActivity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Activity data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.GetActivity(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ActivityDAO for database read of all Activitys
//----------------------------------------------------------------------------
func GetAllActivity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Activity data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.GetAllActivity()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ActivityDAO for database save
//----------------------------------------------------------------------------
func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Activity model
	//----------------------------------------------------------------------------
	var data = model.Activity{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Activity model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Activity data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UpdateActivity(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ActivityDAO for database deletion
//----------------------------------------------------------------------------
func DeleteActivity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Activity data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ActivityDAO.DeleteActivity(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignOrganizationToActivity(activityId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignOrganizationFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignOwnerToActivity(activityId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignOwnerFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignAccountToActivity(activityId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignAccountFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contact on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContactToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactId,_ := strconv.ParseUint( vars["contactId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignContactToActivity(activityId, contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contact on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContactFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignContactFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lead on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLeadToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leadId,_ := strconv.ParseUint( vars["leadId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignLeadToActivity(activityId, leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lead on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLeadFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignLeadFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Opportunity on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityId,_ := strconv.ParseUint( vars["opportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignOpportunityToActivity(activityId, opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Opportunity on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignOpportunityFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Case on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCaseToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	caseId,_ := strconv.ParseUint( vars["caseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignCaseToActivity(activityId, caseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Case on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCaseFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignCaseFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a Activity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.AssignCampaignToActivity(activityId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a Activity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	activityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Activity DAO
	//----------------------------------------------------------------------------
	requestResult := ActivityDAO.UnassignCampaignFromActivity(activityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


