package controller

import (
    AlertDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AlertDAO for database creation
//----------------------------------------------------------------------------
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Alert model
	//----------------------------------------------------------------------------
	data := model.Alert{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Alert model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Alert data access object to create
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.CreateAlert( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AlertDAO to find the relevant Alert
//----------------------------------------------------------------------------
func GetAlert(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Alert data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.GetAlert(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AlertDAO for database read of all Alerts
//----------------------------------------------------------------------------
func GetAllAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Alert data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.GetAllAlert()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AlertDAO for database save
//----------------------------------------------------------------------------
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Alert model
	//----------------------------------------------------------------------------
	var data = model.Alert{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Alert model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Alert data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.UpdateAlert(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AlertDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAlert(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Alert data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AlertDAO.DeleteAlert(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Metric on a Alert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMetricToAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricId,_ := strconv.ParseUint( vars["metricId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.AssignMetricToAlert(alertId, metricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Metric on a Alert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMetricFromAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.UnassignMetricFromAlert(alertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dashboard on a Alert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDashboardToAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardId,_ := strconv.ParseUint( vars["dashboardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.AssignDashboardToAlert(alertId, dashboardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dashboard on a Alert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDashboardFromAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.UnassignDashboardFromAlert(alertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a Alert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.AssignDatasetToAlert(alertId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a Alert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.UnassignDatasetFromAlert(alertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Rule on a Alert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRuleToAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ruleId,_ := strconv.ParseUint( vars["ruleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.AssignRuleToAlert(alertId, ruleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Rule on a Alert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRuleFromAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.UnassignRuleFromAlert(alertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more anomaliesIds as a Anomalies to a Alert
	//----------------------------------------------------------------------------
func AddAnomaliesToAlert(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	anomaliesIds,_ := vars["anomaliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.AddAnomaliesToAlert(alertId, anomaliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more anomaliesIds as a Anomalies from a Alert
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAnomaliesFromAlert(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	anomaliesIds,_ := vars["anomaliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.RemoveAnomaliesFromAlert(alertId, anomaliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more subscribersIds as a Subscribers to a Alert
	//----------------------------------------------------------------------------
func AddSubscribersToAlert(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	subscribersIds,_ := vars["subscribersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.AddSubscribersToAlert(alertId, subscribersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more subscribersIds as a Subscribers from a Alert
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSubscribersFromAlert(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	subscribersIds,_ := vars["subscribersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Alert DAO
	//----------------------------------------------------------------------------
	requestResult := AlertDAO.RemoveSubscribersFromAlert(alertId, subscribersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
