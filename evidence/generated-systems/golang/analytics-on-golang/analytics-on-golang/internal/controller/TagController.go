package controller

import (
    TagDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TagDAO for database creation
//----------------------------------------------------------------------------
func CreateTag(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Tag model
	//----------------------------------------------------------------------------
	data := model.Tag{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Tag model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Tag data access object to create
	//----------------------------------------------------------------------------
	requestResult := TagDAO.CreateTag( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TagDAO to find the relevant Tag
//----------------------------------------------------------------------------
func GetTag(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Tag data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TagDAO.GetTag(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TagDAO for database read of all Tags
//----------------------------------------------------------------------------
func GetAllTag(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Tag data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TagDAO.GetAllTag()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TagDAO for database save
//----------------------------------------------------------------------------
func UpdateTag(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Tag model
	//----------------------------------------------------------------------------
	var data = model.Tag{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Tag model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Tag data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TagDAO.UpdateTag(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TagDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTag(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Tag data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TagDAO.DeleteTag(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Tag
	//----------------------------------------------------------------------------
func AddDatasetsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddDatasetsToTag(tagId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveDatasetsFromTag(tagId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a Tag
	//----------------------------------------------------------------------------
func AddModelsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddModelsToTag(tagId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveModelsFromTag(tagId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelVersionsIds as a ModelVersions to a Tag
	//----------------------------------------------------------------------------
func AddModelVersionsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionsIds,_ := vars["modelVersionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddModelVersionsToTag(tagId, modelVersionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelVersionsIds as a ModelVersions from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelVersionsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionsIds,_ := vars["modelVersionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveModelVersionsFromTag(tagId, modelVersionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dashboardsIds as a Dashboards to a Tag
	//----------------------------------------------------------------------------
func AddDashboardsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddDashboardsToTag(tagId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dashboardsIds as a Dashboards from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDashboardsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveDashboardsFromTag(tagId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a Tag
	//----------------------------------------------------------------------------
func AddReportsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddReportsToTag(tagId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveReportsFromTag(tagId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featureSetsIds as a FeatureSets to a Tag
	//----------------------------------------------------------------------------
func AddFeatureSetsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddFeatureSetsToTag(tagId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featureSetsIds as a FeatureSets from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeatureSetsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveFeatureSetsFromTag(tagId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more metricsIds as a Metrics to a Tag
	//----------------------------------------------------------------------------
func AddMetricsToTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.AddMetricsToTag(tagId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more metricsIds as a Metrics from a Tag
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMetricsFromTag(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tagId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tag DAO
	//----------------------------------------------------------------------------
	requestResult := TagDAO.RemoveMetricsFromTag(tagId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
