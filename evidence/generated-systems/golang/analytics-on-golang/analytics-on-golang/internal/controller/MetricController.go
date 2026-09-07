package controller

import (
    MetricDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MetricDAO for database creation
//----------------------------------------------------------------------------
func CreateMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Metric model
	//----------------------------------------------------------------------------
	data := model.Metric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Metric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Metric data access object to create
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.CreateMetric( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MetricDAO to find the relevant Metric
//----------------------------------------------------------------------------
func GetMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Metric data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.GetMetric(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MetricDAO for database read of all Metrics
//----------------------------------------------------------------------------
func GetAllMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Metric data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.GetAllMetric()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MetricDAO for database save
//----------------------------------------------------------------------------
func UpdateMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Metric model
	//----------------------------------------------------------------------------
	var data = model.Metric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Metric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Metric data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.UpdateMetric(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MetricDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Metric data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MetricDAO.DeleteMetric(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SemanticModel on a Metric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSemanticModelToMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelId,_ := strconv.ParseUint( vars["semanticModelId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.AssignSemanticModelToMetric(metricId, semanticModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SemanticModel on a Metric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSemanticModelFromMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.UnassignSemanticModelFromMetric(metricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Metric
	//----------------------------------------------------------------------------
func AddDatasetsToMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.AddDatasetsToMetric(metricId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Metric
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.RemoveDatasetsFromMetric(metricId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more glossaryTermsIds as a GlossaryTerms to a Metric
	//----------------------------------------------------------------------------
func AddGlossaryTermsToMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.AddGlossaryTermsToMetric(metricId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more glossaryTermsIds as a GlossaryTerms from a Metric
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.RemoveGlossaryTermsFromMetric(metricId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a Metric
	//----------------------------------------------------------------------------
func AddAlertsToMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.AddAlertsToMetric(metricId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a Metric
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.RemoveAlertsFromMetric(metricId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more visualizationsIds as a Visualizations to a Metric
	//----------------------------------------------------------------------------
func AddVisualizationsToMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	visualizationsIds,_ := vars["visualizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.AddVisualizationsToMetric(metricId, visualizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more visualizationsIds as a Visualizations from a Metric
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVisualizationsFromMetric(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	metricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	visualizationsIds,_ := vars["visualizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Metric DAO
	//----------------------------------------------------------------------------
	requestResult := MetricDAO.RemoveVisualizationsFromMetric(metricId, visualizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
