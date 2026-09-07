package controller

import (
    ControlDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ControlDAO for database creation
//----------------------------------------------------------------------------
func CreateControl(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Control model
	//----------------------------------------------------------------------------
	data := model.Control{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Control model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Control data access object to create
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.CreateControl( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ControlDAO to find the relevant Control
//----------------------------------------------------------------------------
func GetControl(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Control data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.GetControl(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ControlDAO for database read of all Controls
//----------------------------------------------------------------------------
func GetAllControl(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Control data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.GetAllControl()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ControlDAO for database save
//----------------------------------------------------------------------------
func UpdateControl(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Control model
	//----------------------------------------------------------------------------
	var data = model.Control{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Control model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Control data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.UpdateControl(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ControlDAO for database deletion
//----------------------------------------------------------------------------
func DeleteControl(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Control data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ControlDAO.DeleteControl(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Policy on a Control
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToControl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AssignPolicyToControl(controlId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Control
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromControl( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.UnassignPolicyFromControl(controlId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more controlTestsIds as a ControlTests to a Control
	//----------------------------------------------------------------------------
func AddControlTestsToControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlTestsIds,_ := vars["controlTestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AddControlTestsToControl(controlId, controlTestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlTestsIds as a ControlTests from a Control
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlTestsFromControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlTestsIds,_ := vars["controlTestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.RemoveControlTestsFromControl(controlId, controlTestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more evidenceIds as a Evidence to a Control
	//----------------------------------------------------------------------------
func AddEvidenceToControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evidenceIds,_ := vars["evidenceIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AddEvidenceToControl(controlId, evidenceIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more evidenceIds as a Evidence from a Control
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEvidenceFromControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evidenceIds,_ := vars["evidenceIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.RemoveEvidenceFromControl(controlId, evidenceIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more risksIds as a Risks to a Control
	//----------------------------------------------------------------------------
func AddRisksToControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	risksIds,_ := vars["risksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AddRisksToControl(controlId, risksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more risksIds as a Risks from a Control
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRisksFromControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	risksIds,_ := vars["risksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.RemoveRisksFromControl(controlId, risksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more obligationsIds as a Obligations to a Control
	//----------------------------------------------------------------------------
func AddObligationsToControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AddObligationsToControl(controlId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more obligationsIds as a Obligations from a Control
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObligationsFromControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.RemoveObligationsFromControl(controlId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more proceduresIds as a Procedures to a Control
	//----------------------------------------------------------------------------
func AddProceduresToControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AddProceduresToControl(controlId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more proceduresIds as a Procedures from a Control
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProceduresFromControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.RemoveProceduresFromControl(controlId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more issuesIds as a Issues to a Control
	//----------------------------------------------------------------------------
func AddIssuesToControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.AddIssuesToControl(controlId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more issuesIds as a Issues from a Control
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveIssuesFromControl(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Control DAO
	//----------------------------------------------------------------------------
	requestResult := ControlDAO.RemoveIssuesFromControl(controlId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
