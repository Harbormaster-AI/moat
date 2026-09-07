package controller

import (
    QualityRuleDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to QualityRuleDAO for database creation
//----------------------------------------------------------------------------
func CreateQualityRule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QualityRule model
	//----------------------------------------------------------------------------
	data := model.QualityRule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QualityRule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QualityRule data access object to create
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.CreateQualityRule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to QualityRuleDAO to find the relevant QualityRule
//----------------------------------------------------------------------------
func GetQualityRule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QualityRule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.GetQualityRule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to QualityRuleDAO for database read of all QualityRules
//----------------------------------------------------------------------------
func GetAllQualityRule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the QualityRule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.GetAllQualityRule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to QualityRuleDAO for database save
//----------------------------------------------------------------------------
func UpdateQualityRule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QualityRule model
	//----------------------------------------------------------------------------
	var data = model.QualityRule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QualityRule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QualityRule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.UpdateQualityRule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to QualityRuleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteQualityRule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QualityRule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := QualityRuleDAO.DeleteQualityRule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a QualityRule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToQualityRule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualityRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualityRule DAO
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.AssignDatasetToQualityRule(qualityRuleId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a QualityRule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromQualityRule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualityRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualityRule DAO
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.UnassignDatasetFromQualityRule(qualityRuleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more checksIds as a Checks to a QualityRule
	//----------------------------------------------------------------------------
func AddChecksToQualityRule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	qualityRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	checksIds,_ := vars["checksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the QualityRule DAO
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.AddChecksToQualityRule(qualityRuleId, checksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more checksIds as a Checks from a QualityRule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChecksFromQualityRule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	qualityRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	checksIds,_ := vars["checksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the QualityRule DAO
	//----------------------------------------------------------------------------
	requestResult := QualityRuleDAO.RemoveChecksFromQualityRule(qualityRuleId, checksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
