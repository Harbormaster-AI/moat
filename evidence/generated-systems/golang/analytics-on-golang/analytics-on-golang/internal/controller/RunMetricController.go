package controller

import (
    RunMetricDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RunMetricDAO for database creation
//----------------------------------------------------------------------------
func CreateRunMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RunMetric model
	//----------------------------------------------------------------------------
	data := model.RunMetric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RunMetric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric data access object to create
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.CreateRunMetric( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RunMetricDAO to find the relevant RunMetric
//----------------------------------------------------------------------------
func GetRunMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RunMetric data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.GetRunMetric(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RunMetricDAO for database read of all RunMetrics
//----------------------------------------------------------------------------
func GetAllRunMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RunMetric data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.GetAllRunMetric()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RunMetricDAO for database save
//----------------------------------------------------------------------------
func UpdateRunMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RunMetric model
	//----------------------------------------------------------------------------
	var data = model.RunMetric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RunMetric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.UpdateRunMetric(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RunMetricDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRunMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RunMetric data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RunMetricDAO.DeleteRunMetric(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a TrainingRun on a RunMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTrainingRunToRunMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunId,_ := strconv.ParseUint( vars["trainingRunId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric DAO
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.AssignTrainingRunToRunMetric(runMetricId, trainingRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TrainingRun on a RunMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTrainingRunFromRunMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric DAO
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.UnassignTrainingRunFromRunMetric(runMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Metric on a RunMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMetricToRunMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricId,_ := strconv.ParseUint( vars["metricId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric DAO
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.AssignMetricToRunMetric(runMetricId, metricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Metric on a RunMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMetricFromRunMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric DAO
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.UnassignMetricFromRunMetric(runMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a RunMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToRunMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric DAO
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.AssignDatasetToRunMetric(runMetricId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a RunMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromRunMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunMetric DAO
	//----------------------------------------------------------------------------
	requestResult := RunMetricDAO.UnassignDatasetFromRunMetric(runMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


