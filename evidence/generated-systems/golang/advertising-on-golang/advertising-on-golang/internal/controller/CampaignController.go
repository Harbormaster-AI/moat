package controller

import (
    CampaignDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CampaignDAO for database creation
//----------------------------------------------------------------------------
func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Campaign model
	//----------------------------------------------------------------------------
	data := model.Campaign{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Campaign model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign data access object to create
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.CreateCampaign( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CampaignDAO to find the relevant Campaign
//----------------------------------------------------------------------------
func GetCampaign(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Campaign data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.GetCampaign(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CampaignDAO for database read of all Campaigns
//----------------------------------------------------------------------------
func GetAllCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Campaign data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.GetAllCampaign()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CampaignDAO for database save
//----------------------------------------------------------------------------
func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Campaign model
	//----------------------------------------------------------------------------
	var data = model.Campaign{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Campaign model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.UpdateCampaign(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CampaignDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Campaign data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CampaignDAO.DeleteCampaign(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a AdAccount on a Campaign
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdAccountToCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountId,_ := strconv.ParseUint( vars["adAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AssignAdAccountToCampaign(campaignId, adAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AdAccount on a Campaign
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdAccountFromCampaign( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.UnassignAdAccountFromCampaign(campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InsertionOrder on a Campaign
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInsertionOrderToCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insertionOrderId,_ := strconv.ParseUint( vars["insertionOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AssignInsertionOrderToCampaign(campaignId, insertionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InsertionOrder on a Campaign
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInsertionOrderFromCampaign( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.UnassignInsertionOrderFromCampaign(campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more lineItemsIds as a LineItems to a Campaign
	//----------------------------------------------------------------------------
func AddLineItemsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemsIds,_ := vars["lineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddLineItemsToCampaign(campaignId, lineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more lineItemsIds as a LineItems from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLineItemsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemsIds,_ := vars["lineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveLineItemsFromCampaign(campaignId, lineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more kpisIds as a Kpis to a Campaign
	//----------------------------------------------------------------------------
func AddKpisToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kpisIds,_ := vars["kpisIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddKpisToCampaign(campaignId, kpisIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more kpisIds as a Kpis from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveKpisFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kpisIds,_ := vars["kpisIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveKpisFromCampaign(campaignId, kpisIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more trackingPixelsIds as a TrackingPixels to a Campaign
	//----------------------------------------------------------------------------
func AddTrackingPixelsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trackingPixelsIds,_ := vars["trackingPixelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddTrackingPixelsToCampaign(campaignId, trackingPixelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more trackingPixelsIds as a TrackingPixels from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTrackingPixelsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trackingPixelsIds,_ := vars["trackingPixelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveTrackingPixelsFromCampaign(campaignId, trackingPixelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more audiencesIds as a Audiences to a Campaign
	//----------------------------------------------------------------------------
func AddAudiencesToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	audiencesIds,_ := vars["audiencesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddAudiencesToCampaign(campaignId, audiencesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more audiencesIds as a Audiences from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAudiencesFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	audiencesIds,_ := vars["audiencesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveAudiencesFromCampaign(campaignId, audiencesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a Campaign
	//----------------------------------------------------------------------------
func AddReportsToCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.AddReportsToCampaign(campaignId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a Campaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	campaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Campaign DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignDAO.RemoveReportsFromCampaign(campaignId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
