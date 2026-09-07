package controller

import (
    ThirdPartyAssessmentDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ThirdPartyAssessmentDAO for database creation
//----------------------------------------------------------------------------
func CreateThirdPartyAssessment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdPartyAssessment model
	//----------------------------------------------------------------------------
	data := model.ThirdPartyAssessment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdPartyAssessment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment data access object to create
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.CreateThirdPartyAssessment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ThirdPartyAssessmentDAO to find the relevant ThirdPartyAssessment
//----------------------------------------------------------------------------
func GetThirdPartyAssessment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdPartyAssessment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.GetThirdPartyAssessment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ThirdPartyAssessmentDAO for database read of all ThirdPartyAssessments
//----------------------------------------------------------------------------
func GetAllThirdPartyAssessment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.GetAllThirdPartyAssessment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ThirdPartyAssessmentDAO for database save
//----------------------------------------------------------------------------
func UpdateThirdPartyAssessment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdPartyAssessment model
	//----------------------------------------------------------------------------
	var data = model.ThirdPartyAssessment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdPartyAssessment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.UpdateThirdPartyAssessment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ThirdPartyAssessmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteThirdPartyAssessment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdPartyAssessment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ThirdPartyAssessmentDAO.DeleteThirdPartyAssessment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ThirdParty on a ThirdPartyAssessment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignThirdPartyToThirdPartyAssessment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	thirdPartyAssessmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartyId,_ := strconv.ParseUint( vars["thirdPartyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.AssignThirdPartyToThirdPartyAssessment(thirdPartyAssessmentId, thirdPartyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ThirdParty on a ThirdPartyAssessment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignThirdPartyFromThirdPartyAssessment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	thirdPartyAssessmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.UnassignThirdPartyFromThirdPartyAssessment(thirdPartyAssessmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more issuesIds as a Issues to a ThirdPartyAssessment
	//----------------------------------------------------------------------------
func AddIssuesToThirdPartyAssessment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyAssessmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.AddIssuesToThirdPartyAssessment(thirdPartyAssessmentId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more issuesIds as a Issues from a ThirdPartyAssessment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveIssuesFromThirdPartyAssessment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyAssessmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issuesIds,_ := vars["issuesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdPartyAssessment DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyAssessmentDAO.RemoveIssuesFromThirdPartyAssessment(thirdPartyAssessmentId, issuesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
