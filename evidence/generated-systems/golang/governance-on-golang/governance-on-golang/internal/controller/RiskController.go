package controller

import (
    RiskDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RiskDAO for database creation
//----------------------------------------------------------------------------
func CreateRisk(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Risk model
	//----------------------------------------------------------------------------
	data := model.Risk{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Risk model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Risk data access object to create
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.CreateRisk( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RiskDAO to find the relevant Risk
//----------------------------------------------------------------------------
func GetRisk(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Risk data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.GetRisk(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RiskDAO for database read of all Risks
//----------------------------------------------------------------------------
func GetAllRisk(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Risk data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.GetAllRisk()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RiskDAO for database save
//----------------------------------------------------------------------------
func UpdateRisk(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Risk model
	//----------------------------------------------------------------------------
	var data = model.Risk{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Risk model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Risk data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.UpdateRisk(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RiskDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRisk(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Risk data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RiskDAO.DeleteRisk(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Risk
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToRisk(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.AssignOrganizationToRisk(riskId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Risk
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromRisk( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.UnassignOrganizationFromRisk(riskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more controlsIds as a Controls to a Risk
	//----------------------------------------------------------------------------
func AddControlsToRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.AddControlsToRisk(riskId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlsIds as a Controls from a Risk
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlsFromRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.RemoveControlsFromRisk(riskId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more assessmentsIds as a Assessments to a Risk
	//----------------------------------------------------------------------------
func AddAssessmentsToRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assessmentsIds,_ := vars["assessmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.AddAssessmentsToRisk(riskId, assessmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assessmentsIds as a Assessments from a Risk
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssessmentsFromRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assessmentsIds,_ := vars["assessmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.RemoveAssessmentsFromRisk(riskId, assessmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more issuesIds as a Issues to a Risk
	//----------------------------------------------------------------------------
func AddIssuesToRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.AddIssuesToRisk(riskId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more issuesIds as a Issues from a Risk
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveIssuesFromRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.RemoveIssuesFromRisk(riskId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more findingsIds as a Findings to a Risk
	//----------------------------------------------------------------------------
func AddFindingsToRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingsIds,_ := vars["findingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.AddFindingsToRisk(riskId, findingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more findingsIds as a Findings from a Risk
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFindingsFromRisk(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	riskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingsIds,_ := vars["findingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Risk DAO
	//----------------------------------------------------------------------------
	requestResult := RiskDAO.RemoveFindingsFromRisk(riskId, findingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
