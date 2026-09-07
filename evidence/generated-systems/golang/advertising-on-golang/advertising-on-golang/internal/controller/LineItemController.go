package controller

import (
    LineItemDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LineItemDAO for database creation
//----------------------------------------------------------------------------
func CreateLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LineItem model
	//----------------------------------------------------------------------------
	data := model.LineItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LineItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.CreateLineItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LineItemDAO to find the relevant LineItem
//----------------------------------------------------------------------------
func GetLineItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LineItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.GetLineItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LineItemDAO for database read of all LineItems
//----------------------------------------------------------------------------
func GetAllLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LineItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.GetAllLineItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LineItemDAO for database save
//----------------------------------------------------------------------------
func UpdateLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LineItem model
	//----------------------------------------------------------------------------
	var data = model.LineItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LineItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.UpdateLineItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LineItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLineItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LineItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LineItemDAO.DeleteLineItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a LineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.AssignCampaignToLineItem(lineItemId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a LineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.UnassignCampaignFromLineItem(lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TargetingProfile on a LineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTargetingProfileToLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	targetingProfileId,_ := strconv.ParseUint( vars["targetingProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.AssignTargetingProfileToLineItem(lineItemId, targetingProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TargetingProfile on a LineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTargetingProfileFromLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.UnassignTargetingProfileFromLineItem(lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Deal on a LineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDealToLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dealId,_ := strconv.ParseUint( vars["dealId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.AssignDealToLineItem(lineItemId, dealId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Deal on a LineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDealFromLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.UnassignDealFromLineItem(lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more placementsIds as a Placements to a LineItem
	//----------------------------------------------------------------------------
func AddPlacementsToLineItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementsIds,_ := vars["placementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.AddPlacementsToLineItem(lineItemId, placementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more placementsIds as a Placements from a LineItem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlacementsFromLineItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementsIds,_ := vars["placementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.RemovePlacementsFromLineItem(lineItemId, placementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more creativesIds as a Creatives to a LineItem
	//----------------------------------------------------------------------------
func AddCreativesToLineItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativesIds,_ := vars["creativesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.AddCreativesToLineItem(lineItemId, creativesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more creativesIds as a Creatives from a LineItem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCreativesFromLineItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativesIds,_ := vars["creativesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.RemoveCreativesFromLineItem(lineItemId, creativesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more performanceMetricsIds as a PerformanceMetrics to a LineItem
	//----------------------------------------------------------------------------
func AddPerformanceMetricsToLineItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performanceMetricsIds,_ := vars["performanceMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.AddPerformanceMetricsToLineItem(lineItemId, performanceMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more performanceMetricsIds as a PerformanceMetrics from a LineItem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePerformanceMetricsFromLineItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performanceMetricsIds,_ := vars["performanceMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineItem DAO
	//----------------------------------------------------------------------------
	requestResult := LineItemDAO.RemovePerformanceMetricsFromLineItem(lineItemId, performanceMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
