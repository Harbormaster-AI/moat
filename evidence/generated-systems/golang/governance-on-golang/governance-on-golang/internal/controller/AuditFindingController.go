package controller

import (
    AuditFindingDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AuditFindingDAO for database creation
//----------------------------------------------------------------------------
func CreateAuditFinding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditFinding model
	//----------------------------------------------------------------------------
	data := model.AuditFinding{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditFinding model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding data access object to create
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.CreateAuditFinding( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AuditFindingDAO to find the relevant AuditFinding
//----------------------------------------------------------------------------
func GetAuditFinding(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditFinding data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.GetAuditFinding(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AuditFindingDAO for database read of all AuditFindings
//----------------------------------------------------------------------------
func GetAllAuditFinding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.GetAllAuditFinding()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AuditFindingDAO for database save
//----------------------------------------------------------------------------
func UpdateAuditFinding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditFinding model
	//----------------------------------------------------------------------------
	var data = model.AuditFinding{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditFinding model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.UpdateAuditFinding(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AuditFindingDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAuditFinding(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditFinding data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AuditFindingDAO.DeleteAuditFinding(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Engagement on a AuditFinding
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEngagementToAuditFinding(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engagementId,_ := strconv.ParseUint( vars["engagementId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.AssignEngagementToAuditFinding(auditFindingId, engagementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Engagement on a AuditFinding
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEngagementFromAuditFinding( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.UnassignEngagementFromAuditFinding(auditFindingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Workpaper on a AuditFinding
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkpaperToAuditFinding(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workpaperId,_ := strconv.ParseUint( vars["workpaperId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.AssignWorkpaperToAuditFinding(auditFindingId, workpaperId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workpaper on a AuditFinding
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkpaperFromAuditFinding( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.UnassignWorkpaperFromAuditFinding(auditFindingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more correctiveActionsIds as a CorrectiveActions to a AuditFinding
	//----------------------------------------------------------------------------
func AddCorrectiveActionsToAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.AddCorrectiveActionsToAuditFinding(auditFindingId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more correctiveActionsIds as a CorrectiveActions from a AuditFinding
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.RemoveCorrectiveActionsFromAuditFinding(auditFindingId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more relatedRisksIds as a RelatedRisks to a AuditFinding
	//----------------------------------------------------------------------------
func AddRelatedRisksToAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedRisksIds,_ := vars["relatedRisksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.AddRelatedRisksToAuditFinding(auditFindingId, relatedRisksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more relatedRisksIds as a RelatedRisks from a AuditFinding
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRelatedRisksFromAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedRisksIds,_ := vars["relatedRisksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.RemoveRelatedRisksFromAuditFinding(auditFindingId, relatedRisksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more relatedControlsIds as a RelatedControls to a AuditFinding
	//----------------------------------------------------------------------------
func AddRelatedControlsToAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedControlsIds,_ := vars["relatedControlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.AddRelatedControlsToAuditFinding(auditFindingId, relatedControlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more relatedControlsIds as a RelatedControls from a AuditFinding
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRelatedControlsFromAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedControlsIds,_ := vars["relatedControlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.RemoveRelatedControlsFromAuditFinding(auditFindingId, relatedControlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more issuesIds as a Issues to a AuditFinding
	//----------------------------------------------------------------------------
func AddIssuesToAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.AddIssuesToAuditFinding(auditFindingId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more issuesIds as a Issues from a AuditFinding
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveIssuesFromAuditFinding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditFindingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditFinding DAO
	//----------------------------------------------------------------------------
	requestResult := AuditFindingDAO.RemoveIssuesFromAuditFinding(auditFindingId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
