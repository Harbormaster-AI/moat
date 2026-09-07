package controller

import (
    PredictionDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PredictionDAO for database creation
//----------------------------------------------------------------------------
func CreatePrediction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Prediction model
	//----------------------------------------------------------------------------
	data := model.Prediction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Prediction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction data access object to create
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.CreatePrediction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PredictionDAO to find the relevant Prediction
//----------------------------------------------------------------------------
func GetPrediction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Prediction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.GetPrediction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PredictionDAO for database read of all Predictions
//----------------------------------------------------------------------------
func GetAllPrediction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Prediction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.GetAllPrediction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PredictionDAO for database save
//----------------------------------------------------------------------------
func UpdatePrediction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Prediction model
	//----------------------------------------------------------------------------
	var data = model.Prediction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Prediction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.UpdatePrediction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PredictionDAO for database deletion
//----------------------------------------------------------------------------
func DeletePrediction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Prediction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PredictionDAO.DeletePrediction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Endpoint on a Prediction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEndpointToPrediction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	predictionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	endpointId,_ := strconv.ParseUint( vars["endpointId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction DAO
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.AssignEndpointToPrediction(predictionId, endpointId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Endpoint on a Prediction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEndpointFromPrediction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	predictionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction DAO
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.UnassignEndpointFromPrediction(predictionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ModelVersion on a Prediction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModelVersionToPrediction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	predictionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionId,_ := strconv.ParseUint( vars["modelVersionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction DAO
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.AssignModelVersionToPrediction(predictionId, modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ModelVersion on a Prediction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModelVersionFromPrediction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	predictionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction DAO
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.UnassignModelVersionFromPrediction(predictionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a Prediction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToPrediction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	predictionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction DAO
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.AssignDatasetToPrediction(predictionId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a Prediction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromPrediction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	predictionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Prediction DAO
	//----------------------------------------------------------------------------
	requestResult := PredictionDAO.UnassignDatasetFromPrediction(predictionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


