package controller

import (
    ForecastDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
	// adds one or more linesIds as a Lines to a Forecast
	//----------------------------------------------------------------------------
func AddLinesToForecast(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.AddLinesToForecast(forecastId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a Forecast
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromForecast(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	forecastId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Forecast DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastDAO.RemoveLinesFromForecast(forecastId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
