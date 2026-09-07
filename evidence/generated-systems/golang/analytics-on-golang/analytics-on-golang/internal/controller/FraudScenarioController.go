package controller

import (
    FraudScenarioDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FraudScenarioDAO for database creation
//----------------------------------------------------------------------------
func CreateFraudScenario(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FraudScenario model
	//----------------------------------------------------------------------------
	data := model.FraudScenario{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FraudScenario model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario data access object to create
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.CreateFraudScenario( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FraudScenarioDAO to find the relevant FraudScenario
//----------------------------------------------------------------------------
func GetFraudScenario(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FraudScenario data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.GetFraudScenario(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FraudScenarioDAO for database read of all FraudScenarios
//----------------------------------------------------------------------------
func GetAllFraudScenario(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.GetAllFraudScenario()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FraudScenarioDAO for database save
//----------------------------------------------------------------------------
func UpdateFraudScenario(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FraudScenario model
	//----------------------------------------------------------------------------
	var data = model.FraudScenario{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FraudScenario model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.UpdateFraudScenario(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FraudScenarioDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFraudScenario(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FraudScenario data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FraudScenarioDAO.DeleteFraudScenario(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a FraudScenario
	//----------------------------------------------------------------------------
func AddModelsToFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.AddModelsToFraudScenario(fraudScenarioId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a FraudScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.RemoveModelsFromFraudScenario(fraudScenarioId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a FraudScenario
	//----------------------------------------------------------------------------
func AddDatasetsToFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.AddDatasetsToFraudScenario(fraudScenarioId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a FraudScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.RemoveDatasetsFromFraudScenario(fraudScenarioId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a FraudScenario
	//----------------------------------------------------------------------------
func AddAlertsToFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.AddAlertsToFraudScenario(fraudScenarioId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a FraudScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.RemoveAlertsFromFraudScenario(fraudScenarioId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more signalsIds as a Signals to a FraudScenario
	//----------------------------------------------------------------------------
func AddSignalsToFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	signalsIds,_ := vars["signalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.AddSignalsToFraudScenario(fraudScenarioId, signalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more signalsIds as a Signals from a FraudScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSignalsFromFraudScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fraudScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	signalsIds,_ := vars["signalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FraudScenario DAO
	//----------------------------------------------------------------------------
	requestResult := FraudScenarioDAO.RemoveSignalsFromFraudScenario(fraudScenarioId, signalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
