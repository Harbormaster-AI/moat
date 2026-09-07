package controller

import (
    ForecastLineDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ForecastLineDAO for database creation
//----------------------------------------------------------------------------
func CreateForecastLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ForecastLine model
	//----------------------------------------------------------------------------
	data := model.ForecastLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ForecastLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.CreateForecastLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ForecastLineDAO to find the relevant ForecastLine
//----------------------------------------------------------------------------
func GetForecastLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ForecastLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.GetForecastLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ForecastLineDAO for database read of all ForecastLines
//----------------------------------------------------------------------------
func GetAllForecastLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.GetAllForecastLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ForecastLineDAO for database save
//----------------------------------------------------------------------------
func UpdateForecastLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ForecastLine model
	//----------------------------------------------------------------------------
	var data = model.ForecastLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ForecastLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.UpdateForecastLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ForecastLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteForecastLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ForecastLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ForecastLineDAO.DeleteForecastLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Forecast on a ForecastLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignForecastToForecastLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	forecastId,_ := strconv.ParseUint( vars["forecastId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.AssignForecastToForecastLine(forecastLineId, forecastId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Forecast on a ForecastLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignForecastFromForecastLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.UnassignForecastFromForecastLine(forecastLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Item on a ForecastLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToForecastLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.AssignItemToForecastLine(forecastLineId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a ForecastLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromForecastLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	forecastLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ForecastLine DAO
	//----------------------------------------------------------------------------
	requestResult := ForecastLineDAO.UnassignItemFromForecastLine(forecastLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


