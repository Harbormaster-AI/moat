package controller

import (
    FeatureDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FeatureDAO for database creation
//----------------------------------------------------------------------------
func CreateFeature(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Feature model
	//----------------------------------------------------------------------------
	data := model.Feature{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Feature model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Feature data access object to create
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.CreateFeature( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FeatureDAO to find the relevant Feature
//----------------------------------------------------------------------------
func GetFeature(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Feature data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.GetFeature(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FeatureDAO for database read of all Features
//----------------------------------------------------------------------------
func GetAllFeature(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Feature data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.GetAllFeature()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FeatureDAO for database save
//----------------------------------------------------------------------------
func UpdateFeature(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Feature model
	//----------------------------------------------------------------------------
	var data = model.Feature{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Feature model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Feature data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.UpdateFeature(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FeatureDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFeature(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Feature data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FeatureDAO.DeleteFeature(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a FeatureSet on a Feature
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFeatureSetToFeature(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetId,_ := strconv.ParseUint( vars["featureSetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.AssignFeatureSetToFeature(featureId, featureSetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a FeatureSet on a Feature
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFeatureSetFromFeature( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.UnassignFeatureSetFromFeature(featureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more sourceDatasetsIds as a SourceDatasets to a Feature
	//----------------------------------------------------------------------------
func AddSourceDatasetsToFeature(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourceDatasetsIds,_ := vars["sourceDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.AddSourceDatasetsToFeature(featureId, sourceDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more sourceDatasetsIds as a SourceDatasets from a Feature
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSourceDatasetsFromFeature(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourceDatasetsIds,_ := vars["sourceDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.RemoveSourceDatasetsFromFeature(featureId, sourceDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a Feature
	//----------------------------------------------------------------------------
func AddModelsToFeature(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.AddModelsToFeature(featureId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a Feature
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromFeature(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.RemoveModelsFromFeature(featureId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more trainingRunsIds as a TrainingRuns to a Feature
	//----------------------------------------------------------------------------
func AddTrainingRunsToFeature(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunsIds,_ := vars["trainingRunsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.AddTrainingRunsToFeature(featureId, trainingRunsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more trainingRunsIds as a TrainingRuns from a Feature
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTrainingRunsFromFeature(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunsIds,_ := vars["trainingRunsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Feature DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureDAO.RemoveTrainingRunsFromFeature(featureId, trainingRunsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
