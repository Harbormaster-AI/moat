package controller

import (
    QualityCheckDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to QualityCheckDAO for database creation
//----------------------------------------------------------------------------
func CreateQualityCheck(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QualityCheck model
	//----------------------------------------------------------------------------
	data := model.QualityCheck{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QualityCheck model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck data access object to create
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.CreateQualityCheck( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to QualityCheckDAO to find the relevant QualityCheck
//----------------------------------------------------------------------------
func GetQualityCheck(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QualityCheck data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.GetQualityCheck(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to QualityCheckDAO for database read of all QualityChecks
//----------------------------------------------------------------------------
func GetAllQualityCheck(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck data access object to get all
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.GetAllQualityCheck()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to QualityCheckDAO for database save
//----------------------------------------------------------------------------
func UpdateQualityCheck(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QualityCheck model
	//----------------------------------------------------------------------------
	var data = model.QualityCheck{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QualityCheck model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.UpdateQualityCheck(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to QualityCheckDAO for database deletion
//----------------------------------------------------------------------------
func DeleteQualityCheck(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QualityCheck data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := QualityCheckDAO.DeleteQualityCheck(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Rule on a QualityCheck
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRuleToQualityCheck(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualityCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ruleId,_ := strconv.ParseUint( vars["ruleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck DAO
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.AssignRuleToQualityCheck(qualityCheckId, ruleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Rule on a QualityCheck
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRuleFromQualityCheck( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualityCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck DAO
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.UnassignRuleFromQualityCheck(qualityCheckId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a QualityCheck
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToQualityCheck(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualityCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck DAO
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.AssignDatasetToQualityCheck(qualityCheckId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a QualityCheck
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromQualityCheck( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualityCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualityCheck DAO
	//----------------------------------------------------------------------------
	requestResult := QualityCheckDAO.UnassignDatasetFromQualityCheck(qualityCheckId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


