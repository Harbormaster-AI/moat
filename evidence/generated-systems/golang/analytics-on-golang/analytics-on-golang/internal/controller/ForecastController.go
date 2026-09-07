package controller

import (
    ForecastDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ForecastDAO for database creation
//----------------------------------------------------------------------------
func CreateForecast(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Forecast model
	//----------------------------------------------------------------------------
	data := model.Forecast{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Forecast model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Forecast data access object to create
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.CreateForecast( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ForecastDAO to find the relevant Forecast
//----------------------------------------------------------------------------
func GetForecast(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Forecast data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.GetForecast(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ForecastDAO for database read of all Forecasts
//----------------------------------------------------------------------------
func GetAllForecast(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Forecast data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.GetAllForecast()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ForecastDAO for database save
//----------------------------------------------------------------------------
func UpdateForecast(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Forecast model
	//----------------------------------------------------------------------------
	var data = model.Forecast{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Forecast model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Forecast data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.UpdateForecast(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ForecastDAO for database deletion
//----------------------------------------------------------------------------
func DeleteForecast(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Forecast data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ForecastDAO.DeleteForecast(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ModelVersion on a Forecast
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModelVersionToForecast(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionId,_ := strconv.ParseUint( vars["modelVersionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.AssignModelVersionToForecast(forecastId, modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ModelVersion on a Forecast
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModelVersionFromForecast( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.UnassignModelVersionFromForecast(forecastId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TimeSeries on a Forecast
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTimeSeriesToForecast(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timeSeriesId,_ := strconv.ParseUint( vars["timeSeriesId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.AssignTimeSeriesToForecast(forecastId, timeSeriesId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TimeSeries on a Forecast
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTimeSeriesFromForecast( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.UnassignTimeSeriesFromForecast(forecastId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Forecast
	//----------------------------------------------------------------------------
func AddDatasetsToForecast(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.AddDatasetsToForecast(forecastId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Forecast
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromForecast(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.RemoveDatasetsFromForecast(forecastId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
