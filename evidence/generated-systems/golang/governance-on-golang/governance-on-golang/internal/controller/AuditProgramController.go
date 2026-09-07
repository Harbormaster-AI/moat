package controller

import (
    AuditProgramDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AuditProgramDAO for database creation
//----------------------------------------------------------------------------
func CreateAuditProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditProgram model
	//----------------------------------------------------------------------------
	data := model.AuditProgram{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditProgram model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram data access object to create
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.CreateAuditProgram( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AuditProgramDAO to find the relevant AuditProgram
//----------------------------------------------------------------------------
func GetAuditProgram(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditProgram data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.GetAuditProgram(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AuditProgramDAO for database read of all AuditPrograms
//----------------------------------------------------------------------------
func GetAllAuditProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.GetAllAuditProgram()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AuditProgramDAO for database save
//----------------------------------------------------------------------------
func UpdateAuditProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AuditProgram model
	//----------------------------------------------------------------------------
	var data = model.AuditProgram{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AuditProgram model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.UpdateAuditProgram(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AuditProgramDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAuditProgram(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AuditProgram data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AuditProgramDAO.DeleteAuditProgram(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a AuditProgram
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToAuditProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.AssignOrganizationToAuditProgram(auditProgramId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a AuditProgram
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromAuditProgram( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	auditProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.UnassignOrganizationFromAuditProgram(auditProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more engagementsIds as a Engagements to a AuditProgram
	//----------------------------------------------------------------------------
func AddEngagementsToAuditProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engagementsIds,_ := vars["engagementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.AddEngagementsToAuditProgram(auditProgramId, engagementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more engagementsIds as a Engagements from a AuditProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEngagementsFromAuditProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	auditProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engagementsIds,_ := vars["engagementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AuditProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AuditProgramDAO.RemoveEngagementsFromAuditProgram(auditProgramId, engagementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
