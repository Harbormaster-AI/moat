package controller

import (
    AnomalyDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AnomalyDAO for database creation
//----------------------------------------------------------------------------
func CreateAnomaly(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Anomaly model
	//----------------------------------------------------------------------------
	data := model.Anomaly{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Anomaly model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly data access object to create
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.CreateAnomaly( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AnomalyDAO to find the relevant Anomaly
//----------------------------------------------------------------------------
func GetAnomaly(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Anomaly data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.GetAnomaly(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AnomalyDAO for database read of all Anomalys
//----------------------------------------------------------------------------
func GetAllAnomaly(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Anomaly data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.GetAllAnomaly()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AnomalyDAO for database save
//----------------------------------------------------------------------------
func UpdateAnomaly(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Anomaly model
	//----------------------------------------------------------------------------
	var data = model.Anomaly{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Anomaly model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.UpdateAnomaly(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AnomalyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAnomaly(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Anomaly data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AnomalyDAO.DeleteAnomaly(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a TimeSeries on a Anomaly
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTimeSeriesToAnomaly(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	anomalyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timeSeriesId,_ := strconv.ParseUint( vars["timeSeriesId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly DAO
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.AssignTimeSeriesToAnomaly(anomalyId, timeSeriesId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TimeSeries on a Anomaly
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTimeSeriesFromAnomaly( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	anomalyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly DAO
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.UnassignTimeSeriesFromAnomaly(anomalyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Alert on a Anomaly
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAlertToAnomaly(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	anomalyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertId,_ := strconv.ParseUint( vars["alertId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly DAO
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.AssignAlertToAnomaly(anomalyId, alertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Alert on a Anomaly
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAlertFromAnomaly( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	anomalyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly DAO
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.UnassignAlertFromAnomaly(anomalyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a Anomaly
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToAnomaly(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	anomalyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly DAO
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.AssignDatasetToAnomaly(anomalyId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a Anomaly
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromAnomaly( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	anomalyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Anomaly DAO
	//----------------------------------------------------------------------------
	requestResult := AnomalyDAO.UnassignDatasetFromAnomaly(anomalyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


