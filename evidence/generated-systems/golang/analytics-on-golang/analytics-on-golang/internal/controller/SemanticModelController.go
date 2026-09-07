package controller

import (
    SemanticModelDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SemanticModelDAO for database creation
//----------------------------------------------------------------------------
func CreateSemanticModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SemanticModel model
	//----------------------------------------------------------------------------
	data := model.SemanticModel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SemanticModel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel data access object to create
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.CreateSemanticModel( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SemanticModelDAO to find the relevant SemanticModel
//----------------------------------------------------------------------------
func GetSemanticModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SemanticModel data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.GetSemanticModel(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SemanticModelDAO for database read of all SemanticModels
//----------------------------------------------------------------------------
func GetAllSemanticModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.GetAllSemanticModel()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SemanticModelDAO for database save
//----------------------------------------------------------------------------
func UpdateSemanticModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SemanticModel model
	//----------------------------------------------------------------------------
	var data = model.SemanticModel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SemanticModel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.UpdateSemanticModel(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SemanticModelDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSemanticModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SemanticModel data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SemanticModelDAO.DeleteSemanticModel(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a SemanticModel
	//----------------------------------------------------------------------------
func AddDatasetsToSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.AddDatasetsToSemanticModel(semanticModelId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a SemanticModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.RemoveDatasetsFromSemanticModel(semanticModelId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more metricsIds as a Metrics to a SemanticModel
	//----------------------------------------------------------------------------
func AddMetricsToSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.AddMetricsToSemanticModel(semanticModelId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more metricsIds as a Metrics from a SemanticModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMetricsFromSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.RemoveMetricsFromSemanticModel(semanticModelId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dimensionsIds as a Dimensions to a SemanticModel
	//----------------------------------------------------------------------------
func AddDimensionsToSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.AddDimensionsToSemanticModel(semanticModelId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dimensionsIds as a Dimensions from a SemanticModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDimensionsFromSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.RemoveDimensionsFromSemanticModel(semanticModelId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more measuresIds as a Measures to a SemanticModel
	//----------------------------------------------------------------------------
func AddMeasuresToSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	measuresIds,_ := vars["measuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.AddMeasuresToSemanticModel(semanticModelId, measuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more measuresIds as a Measures from a SemanticModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMeasuresFromSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	measuresIds,_ := vars["measuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.RemoveMeasuresFromSemanticModel(semanticModelId, measuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more glossaryTermsIds as a GlossaryTerms to a SemanticModel
	//----------------------------------------------------------------------------
func AddGlossaryTermsToSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.AddGlossaryTermsToSemanticModel(semanticModelId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more glossaryTermsIds as a GlossaryTerms from a SemanticModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromSemanticModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	semanticModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SemanticModel DAO
	//----------------------------------------------------------------------------
	requestResult := SemanticModelDAO.RemoveGlossaryTermsFromSemanticModel(semanticModelId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
