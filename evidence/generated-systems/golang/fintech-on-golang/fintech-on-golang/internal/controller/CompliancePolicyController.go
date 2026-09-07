package controller

import (
    CompliancePolicyDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CompliancePolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateCompliancePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CompliancePolicy model
	//----------------------------------------------------------------------------
	data := model.CompliancePolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CompliancePolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CompliancePolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := CompliancePolicyDAO.CreateCompliancePolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CompliancePolicyDAO to find the relevant CompliancePolicy
//----------------------------------------------------------------------------
func GetCompliancePolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CompliancePolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompliancePolicyDAO.GetCompliancePolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CompliancePolicyDAO for database read of all CompliancePolicys
//----------------------------------------------------------------------------
func GetAllCompliancePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CompliancePolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CompliancePolicyDAO.GetAllCompliancePolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CompliancePolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateCompliancePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CompliancePolicy model
	//----------------------------------------------------------------------------
	var data = model.CompliancePolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CompliancePolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CompliancePolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompliancePolicyDAO.UpdateCompliancePolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CompliancePolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCompliancePolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CompliancePolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CompliancePolicyDAO.DeleteCompliancePolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Institution on a CompliancePolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInstitutionToCompliancePolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	compliancePolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	institutionId,_ := strconv.ParseUint( vars["institutionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompliancePolicy DAO
	//----------------------------------------------------------------------------
	requestResult := CompliancePolicyDAO.AssignInstitutionToCompliancePolicy(compliancePolicyId, institutionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Institution on a CompliancePolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInstitutionFromCompliancePolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	compliancePolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompliancePolicy DAO
	//----------------------------------------------------------------------------
	requestResult := CompliancePolicyDAO.UnassignInstitutionFromCompliancePolicy(compliancePolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


