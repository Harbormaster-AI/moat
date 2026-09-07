package controller

import (
    DashboardDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DashboardDAO for database creation
//----------------------------------------------------------------------------
func CreateDashboard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dashboard model
	//----------------------------------------------------------------------------
	data := model.Dashboard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dashboard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard data access object to create
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.CreateDashboard( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DashboardDAO to find the relevant Dashboard
//----------------------------------------------------------------------------
func GetDashboard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dashboard data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.GetDashboard(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DashboardDAO for database read of all Dashboards
//----------------------------------------------------------------------------
func GetAllDashboard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Dashboard data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.GetAllDashboard()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DashboardDAO for database save
//----------------------------------------------------------------------------
func UpdateDashboard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dashboard model
	//----------------------------------------------------------------------------
	var data = model.Dashboard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dashboard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.UpdateDashboard(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DashboardDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDashboard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dashboard data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DashboardDAO.DeleteDashboard(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a Dashboard
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToDashboard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AssignWorkspaceToDashboard(dashboardId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a Dashboard
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromDashboard( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.UnassignWorkspaceFromDashboard(dashboardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more visualizationsIds as a Visualizations to a Dashboard
	//----------------------------------------------------------------------------
func AddVisualizationsToDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	visualizationsIds,_ := vars["visualizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AddVisualizationsToDashboard(dashboardId, visualizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more visualizationsIds as a Visualizations from a Dashboard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVisualizationsFromDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	visualizationsIds,_ := vars["visualizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.RemoveVisualizationsFromDashboard(dashboardId, visualizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a Dashboard
	//----------------------------------------------------------------------------
func AddReportsToDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AddReportsToDashboard(dashboardId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a Dashboard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.RemoveReportsFromDashboard(dashboardId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Dashboard
	//----------------------------------------------------------------------------
func AddDatasetsToDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AddDatasetsToDashboard(dashboardId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Dashboard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.RemoveDatasetsFromDashboard(dashboardId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a Dashboard
	//----------------------------------------------------------------------------
func AddAlertsToDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AddAlertsToDashboard(dashboardId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a Dashboard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.RemoveAlertsFromDashboard(dashboardId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more queriesIds as a Queries to a Dashboard
	//----------------------------------------------------------------------------
func AddQueriesToDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	queriesIds,_ := vars["queriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AddQueriesToDashboard(dashboardId, queriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more queriesIds as a Queries from a Dashboard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQueriesFromDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	queriesIds,_ := vars["queriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.RemoveQueriesFromDashboard(dashboardId, queriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tagsIds as a Tags to a Dashboard
	//----------------------------------------------------------------------------
func AddTagsToDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.AddTagsToDashboard(dashboardId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tagsIds as a Tags from a Dashboard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTagsFromDashboard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dashboardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dashboard DAO
	//----------------------------------------------------------------------------
	requestResult := DashboardDAO.RemoveTagsFromDashboard(dashboardId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
