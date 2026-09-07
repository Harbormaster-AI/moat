package controller

import (
    AuditEngagementDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AuditEngagementDAO for database creation
//----------------------------------------------------------------------------
func CreateAuditEngagement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditEngagement model
	//----------------------------------------------------------------------------
	data := model.AuditEngagement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditEngagement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement data access object to create
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.CreateAuditEngagement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AuditEngagementDAO to find the relevant AuditEngagement
//----------------------------------------------------------------------------
func GetAuditEngagement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditEngagement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.GetAuditEngagement(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AuditEngagementDAO for database read of all AuditEngagements
//----------------------------------------------------------------------------
func GetAllAuditEngagement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.GetAllAuditEngagement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AuditEngagementDAO for database save
//----------------------------------------------------------------------------
func UpdateAuditEngagement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditEngagement model
	//----------------------------------------------------------------------------
	var data = model.AuditEngagement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditEngagement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.UpdateAuditEngagement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AuditEngagementDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAuditEngagement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditEngagement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AuditEngagementDAO.DeleteAuditEngagement(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a AuditProgram on a AuditEngagement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAuditProgramToAuditEngagement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	auditProgramId,_ := strconv.ParseUint( vars["auditProgramId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.AssignAuditProgramToAuditEngagement(auditEngagementId, auditProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AuditProgram on a AuditEngagement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAuditProgramFromAuditEngagement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.UnassignAuditProgramFromAuditEngagement(auditEngagementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more businessUnitsIds as a BusinessUnits to a AuditEngagement
	//----------------------------------------------------------------------------
func AddBusinessUnitsToAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitsIds,_ := vars["businessUnitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.AddBusinessUnitsToAuditEngagement(auditEngagementId, businessUnitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more businessUnitsIds as a BusinessUnits from a AuditEngagement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBusinessUnitsFromAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitsIds,_ := vars["businessUnitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.RemoveBusinessUnitsFromAuditEngagement(auditEngagementId, businessUnitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more controlTestsIds as a ControlTests to a AuditEngagement
	//----------------------------------------------------------------------------
func AddControlTestsToAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlTestsIds,_ := vars["controlTestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.AddControlTestsToAuditEngagement(auditEngagementId, controlTestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlTestsIds as a ControlTests from a AuditEngagement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlTestsFromAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlTestsIds,_ := vars["controlTestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.RemoveControlTestsFromAuditEngagement(auditEngagementId, controlTestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more workpapersIds as a Workpapers to a AuditEngagement
	//----------------------------------------------------------------------------
func AddWorkpapersToAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workpapersIds,_ := vars["workpapersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.AddWorkpapersToAuditEngagement(auditEngagementId, workpapersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workpapersIds as a Workpapers from a AuditEngagement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkpapersFromAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workpapersIds,_ := vars["workpapersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.RemoveWorkpapersFromAuditEngagement(auditEngagementId, workpapersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more findingsIds as a Findings to a AuditEngagement
	//----------------------------------------------------------------------------
func AddFindingsToAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingsIds,_ := vars["findingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.AddFindingsToAuditEngagement(auditEngagementId, findingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more findingsIds as a Findings from a AuditEngagement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFindingsFromAuditEngagement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditEngagementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingsIds,_ := vars["findingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditEngagement DAO
	//----------------------------------------------------------------------------
	requestResult := AuditEngagementDAO.RemoveFindingsFromAuditEngagement(auditEngagementId, findingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
