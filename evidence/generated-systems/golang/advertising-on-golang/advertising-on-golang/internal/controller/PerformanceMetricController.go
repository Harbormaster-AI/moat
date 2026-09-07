package controller

import (
    PerformanceMetricDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PerformanceMetricDAO for database creation
//----------------------------------------------------------------------------
func CreatePerformanceMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PerformanceMetric model
	//----------------------------------------------------------------------------
	data := model.PerformanceMetric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PerformanceMetric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric data access object to create
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.CreatePerformanceMetric( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PerformanceMetricDAO to find the relevant PerformanceMetric
//----------------------------------------------------------------------------
func GetPerformanceMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PerformanceMetric data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.GetPerformanceMetric(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PerformanceMetricDAO for database read of all PerformanceMetrics
//----------------------------------------------------------------------------
func GetAllPerformanceMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.GetAllPerformanceMetric()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PerformanceMetricDAO for database save
//----------------------------------------------------------------------------
func UpdatePerformanceMetric(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PerformanceMetric model
	//----------------------------------------------------------------------------
	var data = model.PerformanceMetric{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PerformanceMetric model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.UpdatePerformanceMetric(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PerformanceMetricDAO for database deletion
//----------------------------------------------------------------------------
func DeletePerformanceMetric(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PerformanceMetric data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PerformanceMetricDAO.DeletePerformanceMetric(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a AdAccount on a PerformanceMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdAccountToPerformanceMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountId,_ := strconv.ParseUint( vars["adAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.AssignAdAccountToPerformanceMetric(performanceMetricId, adAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AdAccount on a PerformanceMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdAccountFromPerformanceMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.UnassignAdAccountFromPerformanceMetric(performanceMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a PerformanceMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToPerformanceMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.AssignCampaignToPerformanceMetric(performanceMetricId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a PerformanceMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromPerformanceMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.UnassignCampaignFromPerformanceMetric(performanceMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LineItem on a PerformanceMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineItemToPerformanceMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemId,_ := strconv.ParseUint( vars["lineItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.AssignLineItemToPerformanceMetric(performanceMetricId, lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineItem on a PerformanceMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineItemFromPerformanceMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.UnassignLineItemFromPerformanceMetric(performanceMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Placement on a PerformanceMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlacementToPerformanceMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementId,_ := strconv.ParseUint( vars["placementId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.AssignPlacementToPerformanceMetric(performanceMetricId, placementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Placement on a PerformanceMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlacementFromPerformanceMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.UnassignPlacementFromPerformanceMetric(performanceMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CreativeAsset on a PerformanceMetric
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCreativeAssetToPerformanceMetric(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeAssetId,_ := strconv.ParseUint( vars["creativeAssetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.AssignCreativeAssetToPerformanceMetric(performanceMetricId, creativeAssetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CreativeAsset on a PerformanceMetric
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCreativeAssetFromPerformanceMetric( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceMetricId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceMetric DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceMetricDAO.UnassignCreativeAssetFromPerformanceMetric(performanceMetricId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


