package controller

import (
    NoteDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to NoteDAO for database creation
//----------------------------------------------------------------------------
func CreateNote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Note model
	//----------------------------------------------------------------------------
	data := model.Note{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Note model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Note data access object to create
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.CreateNote( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to NoteDAO to find the relevant Note
//----------------------------------------------------------------------------
func GetNote(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Note data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.GetNote(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to NoteDAO for database read of all Notes
//----------------------------------------------------------------------------
func GetAllNote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Note data access object to get all
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.GetAllNote()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to NoteDAO for database save
//----------------------------------------------------------------------------
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Note model
	//----------------------------------------------------------------------------
	var data = model.Note{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Note model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Note data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UpdateNote(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to NoteDAO for database deletion
//----------------------------------------------------------------------------
func DeleteNote(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Note data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := NoteDAO.DeleteNote(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignOrganizationToNote(noteId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignOrganizationFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignOwnerToNote(noteId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignOwnerFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignAccountToNote(noteId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignAccountFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contact on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContactToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactId,_ := strconv.ParseUint( vars["contactId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignContactToNote(noteId, contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contact on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContactFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignContactFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Opportunity on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityId,_ := strconv.ParseUint( vars["opportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignOpportunityToNote(noteId, opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Opportunity on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignOpportunityFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Case on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCaseToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	caseId,_ := strconv.ParseUint( vars["caseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignCaseToNote(noteId, caseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Case on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCaseFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignCaseFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lead on a Note
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLeadToNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leadId,_ := strconv.ParseUint( vars["leadId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.AssignLeadToNote(noteId, leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lead on a Note
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLeadFromNote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	noteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Note DAO
	//----------------------------------------------------------------------------
	requestResult := NoteDAO.UnassignLeadFromNote(noteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


