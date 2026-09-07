package controller

import (
    ContactDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ContactDAO for database creation
//----------------------------------------------------------------------------
func CreateContact(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Contact model
	//----------------------------------------------------------------------------
	data := model.Contact{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Contact model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Contact data access object to create
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.CreateContact( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ContactDAO to find the relevant Contact
//----------------------------------------------------------------------------
func GetContact(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Contact data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.GetContact(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ContactDAO for database read of all Contacts
//----------------------------------------------------------------------------
func GetAllContact(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Contact data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.GetAllContact()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ContactDAO for database save
//----------------------------------------------------------------------------
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Contact model
	//----------------------------------------------------------------------------
	var data = model.Contact{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Contact model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Contact data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.UpdateContact(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ContactDAO for database deletion
//----------------------------------------------------------------------------
func DeleteContact(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Contact data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ContactDAO.DeleteContact(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Contact
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AssignOrganizationToContact(contactId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Contact
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromContact( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.UnassignOrganizationFromContact(contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Contact
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AssignAccountToContact(contactId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Contact
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromContact( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.UnassignAccountFromContact(contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Contact
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AssignOwnerToContact(contactId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Contact
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromContact( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.UnassignOwnerFromContact(contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a Contact
	//----------------------------------------------------------------------------
func AddActivitiesToContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AddActivitiesToContact(contactId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a Contact
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.RemoveActivitiesFromContact(contactId, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more opportunitiesIds as a Opportunities to a Contact
	//----------------------------------------------------------------------------
func AddOpportunitiesToContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunitiesIds,_ := vars["opportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AddOpportunitiesToContact(contactId, opportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more opportunitiesIds as a Opportunities from a Contact
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOpportunitiesFromContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunitiesIds,_ := vars["opportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.RemoveOpportunitiesFromContact(contactId, opportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more casesIds as a Cases to a Contact
	//----------------------------------------------------------------------------
func AddCasesToContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	casesIds,_ := vars["casesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AddCasesToContact(contactId, casesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more casesIds as a Cases from a Contact
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCasesFromContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	casesIds,_ := vars["casesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.RemoveCasesFromContact(contactId, casesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a Contact
	//----------------------------------------------------------------------------
func AddCampaignsToContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AddCampaignsToContact(contactId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a Contact
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.RemoveCampaignsFromContact(contactId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more notesIds as a Notes to a Contact
	//----------------------------------------------------------------------------
func AddNotesToContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notesIds,_ := vars["notesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AddNotesToContact(contactId, notesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more notesIds as a Notes from a Contact
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNotesFromContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notesIds,_ := vars["notesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.RemoveNotesFromContact(contactId, notesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more emailMessagesIds as a EmailMessages to a Contact
	//----------------------------------------------------------------------------
func AddEmailMessagesToContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.AddEmailMessagesToContact(contactId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more emailMessagesIds as a EmailMessages from a Contact
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmailMessagesFromContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contactId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailMessagesIds,_ := vars["emailMessagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contact DAO
	//----------------------------------------------------------------------------
	requestResult := ContactDAO.RemoveEmailMessagesFromContact(contactId, emailMessagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
