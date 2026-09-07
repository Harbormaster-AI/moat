package controller

import (
    AnalyticsWorkspaceDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AnalyticsWorkspaceDAO for database creation
//----------------------------------------------------------------------------
func CreateAnalyticsWorkspace(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AnalyticsWorkspace model
	//----------------------------------------------------------------------------
	data := model.AnalyticsWorkspace{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AnalyticsWorkspace model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace data access object to create
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.CreateAnalyticsWorkspace( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AnalyticsWorkspaceDAO to find the relevant AnalyticsWorkspace
//----------------------------------------------------------------------------
func GetAnalyticsWorkspace(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AnalyticsWorkspace data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.GetAnalyticsWorkspace(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AnalyticsWorkspaceDAO for database read of all AnalyticsWorkspaces
//----------------------------------------------------------------------------
func GetAllAnalyticsWorkspace(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.GetAllAnalyticsWorkspace()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AnalyticsWorkspaceDAO for database save
//----------------------------------------------------------------------------
func UpdateAnalyticsWorkspace(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AnalyticsWorkspace model
	//----------------------------------------------------------------------------
	var data = model.AnalyticsWorkspace{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AnalyticsWorkspace model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.UpdateAnalyticsWorkspace(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AnalyticsWorkspaceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAnalyticsWorkspace(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AnalyticsWorkspace data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AnalyticsWorkspaceDAO.DeleteAnalyticsWorkspace(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddDatasetsToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddDatasetsToAnalyticsWorkspace(analyticsWorkspaceId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveDatasetsFromAnalyticsWorkspace(analyticsWorkspaceId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataSourcesIds as a DataSources to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddDataSourcesToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataSourcesIds,_ := vars["dataSourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddDataSourcesToAnalyticsWorkspace(analyticsWorkspaceId, dataSourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataSourcesIds as a DataSources from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataSourcesFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataSourcesIds,_ := vars["dataSourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveDataSourcesFromAnalyticsWorkspace(analyticsWorkspaceId, dataSourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more pipelinesIds as a Pipelines to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddPipelinesToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddPipelinesToAnalyticsWorkspace(analyticsWorkspaceId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more pipelinesIds as a Pipelines from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePipelinesFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemovePipelinesFromAnalyticsWorkspace(analyticsWorkspaceId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dashboardsIds as a Dashboards to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddDashboardsToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddDashboardsToAnalyticsWorkspace(analyticsWorkspaceId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dashboardsIds as a Dashboards from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDashboardsFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveDashboardsFromAnalyticsWorkspace(analyticsWorkspaceId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddReportsToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddReportsToAnalyticsWorkspace(analyticsWorkspaceId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveReportsFromAnalyticsWorkspace(analyticsWorkspaceId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more notebooksIds as a Notebooks to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddNotebooksToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notebooksIds,_ := vars["notebooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddNotebooksToAnalyticsWorkspace(analyticsWorkspaceId, notebooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more notebooksIds as a Notebooks from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNotebooksFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notebooksIds,_ := vars["notebooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveNotebooksFromAnalyticsWorkspace(analyticsWorkspaceId, notebooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddModelsToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddModelsToAnalyticsWorkspace(analyticsWorkspaceId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveModelsFromAnalyticsWorkspace(analyticsWorkspaceId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featureSetsIds as a FeatureSets to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddFeatureSetsToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddFeatureSetsToAnalyticsWorkspace(analyticsWorkspaceId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featureSetsIds as a FeatureSets from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeatureSetsFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveFeatureSetsFromAnalyticsWorkspace(analyticsWorkspaceId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddPoliciesToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddPoliciesToAnalyticsWorkspace(analyticsWorkspaceId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemovePoliciesFromAnalyticsWorkspace(analyticsWorkspaceId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more lineageNodesIds as a LineageNodes to a AnalyticsWorkspace
	//----------------------------------------------------------------------------
func AddLineageNodesToAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineageNodesIds,_ := vars["lineageNodesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.AddLineageNodesToAnalyticsWorkspace(analyticsWorkspaceId, lineageNodesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more lineageNodesIds as a LineageNodes from a AnalyticsWorkspace
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLineageNodesFromAnalyticsWorkspace(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	analyticsWorkspaceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineageNodesIds,_ := vars["lineageNodesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AnalyticsWorkspace DAO
	//----------------------------------------------------------------------------
	requestResult := AnalyticsWorkspaceDAO.RemoveLineageNodesFromAnalyticsWorkspace(analyticsWorkspaceId, lineageNodesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
