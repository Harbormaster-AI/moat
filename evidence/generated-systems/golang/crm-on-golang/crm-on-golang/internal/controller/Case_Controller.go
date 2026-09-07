package controller

import (
    Case_DAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to Case_DAO for database creation
//----------------------------------------------------------------------------
func CreateCase_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Case_ model
	//----------------------------------------------------------------------------
	data := model.Case_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Case_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.CreateCase_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to Case_DAO to find the relevant Case_
//----------------------------------------------------------------------------
func GetCase_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Case_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.GetCase_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to Case_DAO for database read of all Case_s
//----------------------------------------------------------------------------
func GetAllCase_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Case_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.GetAllCase_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to Case_DAO for database save
//----------------------------------------------------------------------------
func UpdateCase_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Case_ model
	//----------------------------------------------------------------------------
	var data = model.Case_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Case_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.UpdateCase_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to Case_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteCase_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Case_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := Case_DAO.DeleteCase_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Case_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToCase_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AssignOrganizationToCase_(case_Id, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Case_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromCase_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.UnassignOrganizationFromCase_(case_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Case_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToCase_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AssignAccountToCase_(case_Id, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Case_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromCase_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.UnassignAccountFromCase_(case_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contact on a Case_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContactToCase_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactId,_ := strconv.ParseUint( vars["contactId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AssignContactToCase_(case_Id, contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contact on a Case_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContactFromCase_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.UnassignContactFromCase_(case_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Case_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToCase_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AssignOwnerToCase_(case_Id, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Case_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromCase_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.UnassignOwnerFromCase_(case_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Team on a Case_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTeamToCase_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamId,_ := strconv.ParseUint( vars["teamId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AssignTeamToCase_(case_Id, teamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Team on a Case_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTeamFromCase_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.UnassignTeamFromCase_(case_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more activitiesIds as a Activities to a Case_
	//----------------------------------------------------------------------------
func AddActivitiesToCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AddActivitiesToCase_(case_Id, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more activitiesIds as a Activities from a Case_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActivitiesFromCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	activitiesIds,_ := vars["activitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.RemoveActivitiesFromCase_(case_Id, activitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more caseCommentsIds as a CaseComments to a Case_
	//----------------------------------------------------------------------------
func AddCaseCommentsToCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	caseCommentsIds,_ := vars["caseCommentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AddCaseCommentsToCase_(case_Id, caseCommentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more caseCommentsIds as a CaseComments from a Case_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCaseCommentsFromCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	caseCommentsIds,_ := vars["caseCommentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.RemoveCaseCommentsFromCase_(case_Id, caseCommentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more emailsIds as a Emails to a Case_
	//----------------------------------------------------------------------------
func AddEmailsToCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailsIds,_ := vars["emailsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AddEmailsToCase_(case_Id, emailsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more emailsIds as a Emails from a Case_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmailsFromCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	emailsIds,_ := vars["emailsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.RemoveEmailsFromCase_(case_Id, emailsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more relatedOpportunitiesIds as a RelatedOpportunities to a Case_
	//----------------------------------------------------------------------------
func AddRelatedOpportunitiesToCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedOpportunitiesIds,_ := vars["relatedOpportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.AddRelatedOpportunitiesToCase_(case_Id, relatedOpportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more relatedOpportunitiesIds as a RelatedOpportunities from a Case_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRelatedOpportunitiesFromCase_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	case_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedOpportunitiesIds,_ := vars["relatedOpportunitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Case_ DAO
	//----------------------------------------------------------------------------
	requestResult := Case_DAO.RemoveRelatedOpportunitiesFromCase_(case_Id, relatedOpportunitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
