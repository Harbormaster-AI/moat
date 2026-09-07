package controller

import (
    FraudSignalDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FraudSignalDAO for database creation
//----------------------------------------------------------------------------
func CreateFraudSignal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FraudSignal model
	//----------------------------------------------------------------------------
	data := model.FraudSignal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FraudSignal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal data access object to create
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.CreateFraudSignal( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FraudSignalDAO to find the relevant FraudSignal
//----------------------------------------------------------------------------
func GetFraudSignal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FraudSignal data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.GetFraudSignal(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FraudSignalDAO for database read of all FraudSignals
//----------------------------------------------------------------------------
func GetAllFraudSignal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.GetAllFraudSignal()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FraudSignalDAO for database save
//----------------------------------------------------------------------------
func UpdateFraudSignal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FraudSignal model
	//----------------------------------------------------------------------------
	var data = model.FraudSignal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FraudSignal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.UpdateFraudSignal(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FraudSignalDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFraudSignal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FraudSignal data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FraudSignalDAO.DeleteFraudSignal(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Scenario on a FraudSignal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignScenarioToFraudSignal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fraudSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	scenarioId,_ := strconv.ParseUint( vars["scenarioId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal DAO
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.AssignScenarioToFraudSignal(fraudSignalId, scenarioId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Scenario on a FraudSignal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignScenarioFromFraudSignal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fraudSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal DAO
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.UnassignScenarioFromFraudSignal(fraudSignalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dataset on a FraudSignal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDatasetToFraudSignal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fraudSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetId,_ := strconv.ParseUint( vars["datasetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal DAO
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.AssignDatasetToFraudSignal(fraudSignalId, datasetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dataset on a FraudSignal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDatasetFromFraudSignal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fraudSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal DAO
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.UnassignDatasetFromFraudSignal(fraudSignalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ModelVersion on a FraudSignal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModelVersionToFraudSignal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fraudSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionId,_ := strconv.ParseUint( vars["modelVersionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal DAO
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.AssignModelVersionToFraudSignal(fraudSignalId, modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ModelVersion on a FraudSignal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModelVersionFromFraudSignal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fraudSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FraudSignal DAO
	//----------------------------------------------------------------------------
	requestResult := FraudSignalDAO.UnassignModelVersionFromFraudSignal(fraudSignalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


