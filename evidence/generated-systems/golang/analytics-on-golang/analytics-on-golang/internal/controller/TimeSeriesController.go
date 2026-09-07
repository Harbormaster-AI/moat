package controller

import (
    TimeSeriesDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TimeSeriesDAO for database creation
//----------------------------------------------------------------------------
func CreateTimeSeries(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TimeSeries model
	//----------------------------------------------------------------------------
	data := model.TimeSeries{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TimeSeries model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries data access object to create
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.CreateTimeSeries( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TimeSeriesDAO to find the relevant TimeSeries
//----------------------------------------------------------------------------
func GetTimeSeries(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TimeSeries data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.GetTimeSeries(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TimeSeriesDAO for database read of all TimeSeriess
//----------------------------------------------------------------------------
func GetAllTimeSeries(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.GetAllTimeSeries()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TimeSeriesDAO for database save
//----------------------------------------------------------------------------
func UpdateTimeSeries(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TimeSeries model
	//----------------------------------------------------------------------------
	var data = model.TimeSeries{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TimeSeries model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.UpdateTimeSeries(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TimeSeriesDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTimeSeries(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TimeSeries data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TimeSeriesDAO.DeleteTimeSeries(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a TimeSeries
	//----------------------------------------------------------------------------
func AddDatasetsToTimeSeries(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timeSeriesId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries DAO
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.AddDatasetsToTimeSeries(timeSeriesId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a TimeSeries
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromTimeSeries(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timeSeriesId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries DAO
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.RemoveDatasetsFromTimeSeries(timeSeriesId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more forecastsIds as a Forecasts to a TimeSeries
	//----------------------------------------------------------------------------
func AddForecastsToTimeSeries(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timeSeriesId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	forecastsIds,_ := vars["forecastsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries DAO
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.AddForecastsToTimeSeries(timeSeriesId, forecastsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more forecastsIds as a Forecasts from a TimeSeries
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveForecastsFromTimeSeries(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timeSeriesId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	forecastsIds,_ := vars["forecastsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries DAO
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.RemoveForecastsFromTimeSeries(timeSeriesId, forecastsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more anomaliesIds as a Anomalies to a TimeSeries
	//----------------------------------------------------------------------------
func AddAnomaliesToTimeSeries(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timeSeriesId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	anomaliesIds,_ := vars["anomaliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries DAO
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.AddAnomaliesToTimeSeries(timeSeriesId, anomaliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more anomaliesIds as a Anomalies from a TimeSeries
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAnomaliesFromTimeSeries(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timeSeriesId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	anomaliesIds,_ := vars["anomaliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TimeSeries DAO
	//----------------------------------------------------------------------------
	requestResult := TimeSeriesDAO.RemoveAnomaliesFromTimeSeries(timeSeriesId, anomaliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
