package controller

import (
    PolicyCoverageDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PolicyCoverageDAO for database creation
//----------------------------------------------------------------------------
func CreatePolicyCoverage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PolicyCoverage model
	//----------------------------------------------------------------------------
	data := model.PolicyCoverage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PolicyCoverage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage data access object to create
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.CreatePolicyCoverage( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PolicyCoverageDAO to find the relevant PolicyCoverage
//----------------------------------------------------------------------------
func GetPolicyCoverage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PolicyCoverage data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.GetPolicyCoverage(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PolicyCoverageDAO for database read of all PolicyCoverages
//----------------------------------------------------------------------------
func GetAllPolicyCoverage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.GetAllPolicyCoverage()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PolicyCoverageDAO for database save
//----------------------------------------------------------------------------
func UpdatePolicyCoverage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PolicyCoverage model
	//----------------------------------------------------------------------------
	var data = model.PolicyCoverage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PolicyCoverage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.UpdatePolicyCoverage(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PolicyCoverageDAO for database deletion
//----------------------------------------------------------------------------
func DeletePolicyCoverage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PolicyCoverage data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PolicyCoverageDAO.DeletePolicyCoverage(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Policy on a PolicyCoverage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToPolicyCoverage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyCoverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.AssignPolicyToPolicyCoverage(policyCoverageId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a PolicyCoverage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromPolicyCoverage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyCoverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.UnassignPolicyFromPolicyCoverage(policyCoverageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more insuredObjectsIds as a InsuredObjects to a PolicyCoverage
	//----------------------------------------------------------------------------
func AddInsuredObjectsToPolicyCoverage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyCoverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectsIds,_ := vars["insuredObjectsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.AddInsuredObjectsToPolicyCoverage(policyCoverageId, insuredObjectsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more insuredObjectsIds as a InsuredObjects from a PolicyCoverage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInsuredObjectsFromPolicyCoverage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyCoverageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectsIds,_ := vars["insuredObjectsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PolicyCoverage DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyCoverageDAO.RemoveInsuredObjectsFromPolicyCoverage(policyCoverageId, insuredObjectsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
