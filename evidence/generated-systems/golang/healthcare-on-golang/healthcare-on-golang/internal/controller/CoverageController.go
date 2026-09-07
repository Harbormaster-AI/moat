package controller

import (
    CoverageDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CoverageDAO for database creation
//----------------------------------------------------------------------------
func CreateCoverage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Coverage model
	//----------------------------------------------------------------------------
	data := model.Coverage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Coverage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Coverage data access object to create
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.CreateCoverage( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CoverageDAO to find the relevant Coverage
//----------------------------------------------------------------------------
func GetCoverage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Coverage data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.GetCoverage(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CoverageDAO for database read of all Coverages
//----------------------------------------------------------------------------
func GetAllCoverage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Coverage data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.GetAllCoverage()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CoverageDAO for database save
//----------------------------------------------------------------------------
func UpdateCoverage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Coverage model
	//----------------------------------------------------------------------------
	var data = model.Coverage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Coverage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Coverage data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.UpdateCoverage(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CoverageDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCoverage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Coverage data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CoverageDAO.DeleteCoverage(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a Coverage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToCoverage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.AssignPatientToCoverage(coverageId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a Coverage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromCoverage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.UnassignPatientFromCoverage(coverageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plan on a Coverage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlanToCoverage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	planId,_ := strconv.ParseUint( vars["planId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.AssignPlanToCoverage(coverageId, planId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plan on a Coverage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlanFromCoverage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.UnassignPlanFromCoverage(coverageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a Coverage
	//----------------------------------------------------------------------------
func AddClaimsToCoverage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.AddClaimsToCoverage(coverageId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a Coverage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromCoverage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.RemoveClaimsFromCoverage(coverageId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more authorizationsIds as a Authorizations to a Coverage
	//----------------------------------------------------------------------------
func AddAuthorizationsToCoverage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	authorizationsIds,_ := vars["authorizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.AddAuthorizationsToCoverage(coverageId, authorizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more authorizationsIds as a Authorizations from a Coverage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAuthorizationsFromCoverage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	coverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	authorizationsIds,_ := vars["authorizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Coverage DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDAO.RemoveAuthorizationsFromCoverage(coverageId, authorizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
