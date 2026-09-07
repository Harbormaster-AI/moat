package controller

import (
    BusinessGlossaryTermDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BusinessGlossaryTermDAO for database creation
//----------------------------------------------------------------------------
func CreateBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BusinessGlossaryTerm model
	//----------------------------------------------------------------------------
	data := model.BusinessGlossaryTerm{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BusinessGlossaryTerm model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm data access object to create
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.CreateBusinessGlossaryTerm( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BusinessGlossaryTermDAO to find the relevant BusinessGlossaryTerm
//----------------------------------------------------------------------------
func GetBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BusinessGlossaryTerm data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.GetBusinessGlossaryTerm(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BusinessGlossaryTermDAO for database read of all BusinessGlossaryTerms
//----------------------------------------------------------------------------
func GetAllBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.GetAllBusinessGlossaryTerm()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BusinessGlossaryTermDAO for database save
//----------------------------------------------------------------------------
func UpdateBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BusinessGlossaryTerm model
	//----------------------------------------------------------------------------
	var data = model.BusinessGlossaryTerm{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BusinessGlossaryTerm model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.UpdateBusinessGlossaryTerm(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BusinessGlossaryTermDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BusinessGlossaryTerm data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BusinessGlossaryTermDAO.DeleteBusinessGlossaryTerm(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more relatedTermsIds as a RelatedTerms to a BusinessGlossaryTerm
	//----------------------------------------------------------------------------
func AddRelatedTermsToBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedTermsIds,_ := vars["relatedTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.AddRelatedTermsToBusinessGlossaryTerm(businessGlossaryTermId, relatedTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more relatedTermsIds as a RelatedTerms from a BusinessGlossaryTerm
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRelatedTermsFromBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedTermsIds,_ := vars["relatedTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.RemoveRelatedTermsFromBusinessGlossaryTerm(businessGlossaryTermId, relatedTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more metricsIds as a Metrics to a BusinessGlossaryTerm
	//----------------------------------------------------------------------------
func AddMetricsToBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.AddMetricsToBusinessGlossaryTerm(businessGlossaryTermId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more metricsIds as a Metrics from a BusinessGlossaryTerm
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMetricsFromBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.RemoveMetricsFromBusinessGlossaryTerm(businessGlossaryTermId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a BusinessGlossaryTerm
	//----------------------------------------------------------------------------
func AddDatasetsToBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.AddDatasetsToBusinessGlossaryTerm(businessGlossaryTermId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a BusinessGlossaryTerm
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.RemoveDatasetsFromBusinessGlossaryTerm(businessGlossaryTermId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dimensionsIds as a Dimensions to a BusinessGlossaryTerm
	//----------------------------------------------------------------------------
func AddDimensionsToBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.AddDimensionsToBusinessGlossaryTerm(businessGlossaryTermId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dimensionsIds as a Dimensions from a BusinessGlossaryTerm
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDimensionsFromBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.RemoveDimensionsFromBusinessGlossaryTerm(businessGlossaryTermId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more measuresIds as a Measures to a BusinessGlossaryTerm
	//----------------------------------------------------------------------------
func AddMeasuresToBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	measuresIds,_ := vars["measuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.AddMeasuresToBusinessGlossaryTerm(businessGlossaryTermId, measuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more measuresIds as a Measures from a BusinessGlossaryTerm
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMeasuresFromBusinessGlossaryTerm(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessGlossaryTermId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	measuresIds,_ := vars["measuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessGlossaryTerm DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessGlossaryTermDAO.RemoveMeasuresFromBusinessGlossaryTerm(businessGlossaryTermId, measuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
