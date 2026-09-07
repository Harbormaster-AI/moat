package controller

import (
    RecommendationScenarioDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RecommendationScenarioDAO for database creation
//----------------------------------------------------------------------------
func CreateRecommendationScenario(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RecommendationScenario model
	//----------------------------------------------------------------------------
	data := model.RecommendationScenario{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RecommendationScenario model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario data access object to create
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.CreateRecommendationScenario( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RecommendationScenarioDAO to find the relevant RecommendationScenario
//----------------------------------------------------------------------------
func GetRecommendationScenario(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RecommendationScenario data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.GetRecommendationScenario(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RecommendationScenarioDAO for database read of all RecommendationScenarios
//----------------------------------------------------------------------------
func GetAllRecommendationScenario(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.GetAllRecommendationScenario()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RecommendationScenarioDAO for database save
//----------------------------------------------------------------------------
func UpdateRecommendationScenario(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RecommendationScenario model
	//----------------------------------------------------------------------------
	var data = model.RecommendationScenario{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RecommendationScenario model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.UpdateRecommendationScenario(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RecommendationScenarioDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRecommendationScenario(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RecommendationScenario data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RecommendationScenarioDAO.DeleteRecommendationScenario(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a RecommendationScenario
	//----------------------------------------------------------------------------
func AddModelsToRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.AddModelsToRecommendationScenario(recommendationScenarioId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a RecommendationScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.RemoveModelsFromRecommendationScenario(recommendationScenarioId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a RecommendationScenario
	//----------------------------------------------------------------------------
func AddDatasetsToRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.AddDatasetsToRecommendationScenario(recommendationScenarioId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a RecommendationScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.RemoveDatasetsFromRecommendationScenario(recommendationScenarioId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more experimentsIds as a Experiments to a RecommendationScenario
	//----------------------------------------------------------------------------
func AddExperimentsToRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.AddExperimentsToRecommendationScenario(recommendationScenarioId, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more experimentsIds as a Experiments from a RecommendationScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExperimentsFromRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.RemoveExperimentsFromRecommendationScenario(recommendationScenarioId, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a RecommendationScenario
	//----------------------------------------------------------------------------
func AddAlertsToRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.AddAlertsToRecommendationScenario(recommendationScenarioId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a RecommendationScenario
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromRecommendationScenario(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recommendationScenarioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecommendationScenario DAO
	//----------------------------------------------------------------------------
	requestResult := RecommendationScenarioDAO.RemoveAlertsFromRecommendationScenario(recommendationScenarioId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
