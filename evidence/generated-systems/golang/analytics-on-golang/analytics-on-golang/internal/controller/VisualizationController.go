package controller

import (
    VisualizationDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to VisualizationDAO for database creation
//----------------------------------------------------------------------------
func CreateVisualization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Visualization model
	//----------------------------------------------------------------------------
	data := model.Visualization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Visualization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Visualization data access object to create
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.CreateVisualization( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to VisualizationDAO to find the relevant Visualization
//----------------------------------------------------------------------------
func GetVisualization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Visualization data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.GetVisualization(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to VisualizationDAO for database read of all Visualizations
//----------------------------------------------------------------------------
func GetAllVisualization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Visualization data access object to get all
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.GetAllVisualization()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to VisualizationDAO for database save
//----------------------------------------------------------------------------
func UpdateVisualization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Visualization model
	//----------------------------------------------------------------------------
	var data = model.Visualization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Visualization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Visualization data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.UpdateVisualization(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to VisualizationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteVisualization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Visualization data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := VisualizationDAO.DeleteVisualization(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Dashboard on a Visualization
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDashboardToVisualization(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardId,_ := strconv.ParseUint( vars["dashboardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.AssignDashboardToVisualization(visualizationId, dashboardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dashboard on a Visualization
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDashboardFromVisualization( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.UnassignDashboardFromVisualization(visualizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Report on a Visualization
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignReportToVisualization(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportId,_ := strconv.ParseUint( vars["reportId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.AssignReportToVisualization(visualizationId, reportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Report on a Visualization
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignReportFromVisualization( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.UnassignReportFromVisualization(visualizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more metricsIds as a Metrics to a Visualization
	//----------------------------------------------------------------------------
func AddMetricsToVisualization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.AddMetricsToVisualization(visualizationId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more metricsIds as a Metrics from a Visualization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMetricsFromVisualization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.RemoveMetricsFromVisualization(visualizationId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dimensionsIds as a Dimensions to a Visualization
	//----------------------------------------------------------------------------
func AddDimensionsToVisualization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.AddDimensionsToVisualization(visualizationId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dimensionsIds as a Dimensions from a Visualization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDimensionsFromVisualization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.RemoveDimensionsFromVisualization(visualizationId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Visualization
	//----------------------------------------------------------------------------
func AddDatasetsToVisualization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.AddDatasetsToVisualization(visualizationId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Visualization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromVisualization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	visualizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Visualization DAO
	//----------------------------------------------------------------------------
	requestResult := VisualizationDAO.RemoveDatasetsFromVisualization(visualizationId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
