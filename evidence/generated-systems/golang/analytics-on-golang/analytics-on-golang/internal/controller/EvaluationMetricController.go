package controller

import (
    EvaluationMetricDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EvaluationMetricDAO for database creation
//----------------------------------------------------------------------------
func CreateEvaluationMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EvaluationMetric model
	//----------------------------------------------------------------------------
	data := model.EvaluationMetric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EvaluationMetric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric data access object to create
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.CreateEvaluationMetric( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EvaluationMetricDAO to find the relevant EvaluationMetric
//----------------------------------------------------------------------------
func GetEvaluationMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EvaluationMetric data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.GetEvaluationMetric(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EvaluationMetricDAO for database read of all EvaluationMetrics
//----------------------------------------------------------------------------
func GetAllEvaluationMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.GetAllEvaluationMetric()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EvaluationMetricDAO for database save
//----------------------------------------------------------------------------
func UpdateEvaluationMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EvaluationMetric model
	//----------------------------------------------------------------------------
	var data = model.EvaluationMetric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EvaluationMetric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.UpdateEvaluationMetric(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EvaluationMetricDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEvaluationMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EvaluationMetric data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EvaluationMetricDAO.DeleteEvaluationMetric(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ModelVersion on a EvaluationMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModelVersionToEvaluationMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evaluationMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionId,_ := strconv.ParseUint( vars["modelVersionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric DAO
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.AssignModelVersionToEvaluationMetric(evaluationMetricId, modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ModelVersion on a EvaluationMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModelVersionFromEvaluationMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evaluationMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric DAO
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.UnassignModelVersionFromEvaluationMetric(evaluationMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Metric on a EvaluationMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMetricToEvaluationMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evaluationMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricId,_ := strconv.ParseUint( vars["metricId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric DAO
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.AssignMetricToEvaluationMetric(evaluationMetricId, metricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Metric on a EvaluationMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMetricFromEvaluationMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evaluationMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric DAO
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.UnassignMetricFromEvaluationMetric(evaluationMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a EvaluationMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToEvaluationMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evaluationMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric DAO
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.AssignDatasetToEvaluationMetric(evaluationMetricId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a EvaluationMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromEvaluationMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evaluationMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EvaluationMetric DAO
	//----------------------------------------------------------------------------
	requestResult := EvaluationMetricDAO.UnassignDatasetFromEvaluationMetric(evaluationMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


