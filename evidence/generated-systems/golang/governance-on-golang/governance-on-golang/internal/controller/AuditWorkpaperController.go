package controller

import (
    AuditWorkpaperDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AuditWorkpaperDAO for database creation
//----------------------------------------------------------------------------
func CreateAuditWorkpaper(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditWorkpaper model
	//----------------------------------------------------------------------------
	data := model.AuditWorkpaper{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditWorkpaper model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper data access object to create
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.CreateAuditWorkpaper( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AuditWorkpaperDAO to find the relevant AuditWorkpaper
//----------------------------------------------------------------------------
func GetAuditWorkpaper(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditWorkpaper data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.GetAuditWorkpaper(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AuditWorkpaperDAO for database read of all AuditWorkpapers
//----------------------------------------------------------------------------
func GetAllAuditWorkpaper(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.GetAllAuditWorkpaper()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AuditWorkpaperDAO for database save
//----------------------------------------------------------------------------
func UpdateAuditWorkpaper(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditWorkpaper model
	//----------------------------------------------------------------------------
	var data = model.AuditWorkpaper{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditWorkpaper model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.UpdateAuditWorkpaper(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AuditWorkpaperDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAuditWorkpaper(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditWorkpaper data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AuditWorkpaperDAO.DeleteAuditWorkpaper(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Engagement on a AuditWorkpaper
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEngagementToAuditWorkpaper(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditWorkpaperId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engagementId,_ := strconv.ParseUint( vars["engagementId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper DAO
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.AssignEngagementToAuditWorkpaper(auditWorkpaperId, engagementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Engagement on a AuditWorkpaper
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEngagementFromAuditWorkpaper( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditWorkpaperId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper DAO
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.UnassignEngagementFromAuditWorkpaper(auditWorkpaperId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more evidenceIds as a Evidence to a AuditWorkpaper
	//----------------------------------------------------------------------------
func AddEvidenceToAuditWorkpaper(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditWorkpaperId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evidenceIds,_ := vars["evidenceIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper DAO
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.AddEvidenceToAuditWorkpaper(auditWorkpaperId, evidenceIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more evidenceIds as a Evidence from a AuditWorkpaper
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEvidenceFromAuditWorkpaper(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditWorkpaperId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evidenceIds,_ := vars["evidenceIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper DAO
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.RemoveEvidenceFromAuditWorkpaper(auditWorkpaperId, evidenceIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more findingsIds as a Findings to a AuditWorkpaper
	//----------------------------------------------------------------------------
func AddFindingsToAuditWorkpaper(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditWorkpaperId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingsIds,_ := vars["findingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper DAO
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.AddFindingsToAuditWorkpaper(auditWorkpaperId, findingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more findingsIds as a Findings from a AuditWorkpaper
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFindingsFromAuditWorkpaper(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditWorkpaperId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingsIds,_ := vars["findingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditWorkpaper DAO
	//----------------------------------------------------------------------------
	requestResult := AuditWorkpaperDAO.RemoveFindingsFromAuditWorkpaper(auditWorkpaperId, findingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
