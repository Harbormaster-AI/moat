package controller

import (
    ReportDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ReportDAO for database creation
//----------------------------------------------------------------------------
func CreateReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Report model
	//----------------------------------------------------------------------------
	data := model.Report{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Report model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Report data access object to create
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.CreateReport( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ReportDAO to find the relevant Report
//----------------------------------------------------------------------------
func GetReport(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Report data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.GetReport(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ReportDAO for database read of all Reports
//----------------------------------------------------------------------------
func GetAllReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Report data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.GetAllReport()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ReportDAO for database save
//----------------------------------------------------------------------------
func UpdateReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Report model
	//----------------------------------------------------------------------------
	var data = model.Report{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Report model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Report data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.UpdateReport(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ReportDAO for database deletion
//----------------------------------------------------------------------------
func DeleteReport(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Report data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ReportDAO.DeleteReport(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a Report
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AssignWorkspaceToReport(reportId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a Report
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.UnassignWorkspaceFromReport(reportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more visualizationsIds as a Visualizations to a Report
	//----------------------------------------------------------------------------
func AddVisualizationsToReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	visualizationsIds,_ := vars["visualizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AddVisualizationsToReport(reportId, visualizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more visualizationsIds as a Visualizations from a Report
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVisualizationsFromReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	visualizationsIds,_ := vars["visualizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.RemoveVisualizationsFromReport(reportId, visualizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Report
	//----------------------------------------------------------------------------
func AddDatasetsToReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AddDatasetsToReport(reportId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Report
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.RemoveDatasetsFromReport(reportId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more semanticModelsIds as a SemanticModels to a Report
	//----------------------------------------------------------------------------
func AddSemanticModelsToReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelsIds,_ := vars["semanticModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AddSemanticModelsToReport(reportId, semanticModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more semanticModelsIds as a SemanticModels from a Report
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSemanticModelsFromReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelsIds,_ := vars["semanticModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.RemoveSemanticModelsFromReport(reportId, semanticModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more queriesIds as a Queries to a Report
	//----------------------------------------------------------------------------
func AddQueriesToReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	queriesIds,_ := vars["queriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AddQueriesToReport(reportId, queriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more queriesIds as a Queries from a Report
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQueriesFromReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	queriesIds,_ := vars["queriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.RemoveQueriesFromReport(reportId, queriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tagsIds as a Tags to a Report
	//----------------------------------------------------------------------------
func AddTagsToReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AddTagsToReport(reportId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tagsIds as a Tags from a Report
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTagsFromReport(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.RemoveTagsFromReport(reportId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
