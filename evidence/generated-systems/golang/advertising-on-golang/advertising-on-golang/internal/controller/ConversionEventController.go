package controller

import (
    ConversionEventDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ConversionEventDAO for database creation
//----------------------------------------------------------------------------
func CreateConversionEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ConversionEvent model
	//----------------------------------------------------------------------------
	data := model.ConversionEvent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ConversionEvent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent data access object to create
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.CreateConversionEvent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ConversionEventDAO to find the relevant ConversionEvent
//----------------------------------------------------------------------------
func GetConversionEvent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ConversionEvent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.GetConversionEvent(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ConversionEventDAO for database read of all ConversionEvents
//----------------------------------------------------------------------------
func GetAllConversionEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.GetAllConversionEvent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ConversionEventDAO for database save
//----------------------------------------------------------------------------
func UpdateConversionEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ConversionEvent model
	//----------------------------------------------------------------------------
	var data = model.ConversionEvent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ConversionEvent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.UpdateConversionEvent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ConversionEventDAO for database deletion
//----------------------------------------------------------------------------
func DeleteConversionEvent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ConversionEvent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ConversionEventDAO.DeleteConversionEvent(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a ConversionEvent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToConversionEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conversionEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent DAO
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.AssignCampaignToConversionEvent(conversionEventId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a ConversionEvent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromConversionEvent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conversionEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent DAO
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.UnassignCampaignFromConversionEvent(conversionEventId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LineItem on a ConversionEvent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineItemToConversionEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conversionEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemId,_ := strconv.ParseUint( vars["lineItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent DAO
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.AssignLineItemToConversionEvent(conversionEventId, lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineItem on a ConversionEvent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineItemFromConversionEvent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conversionEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent DAO
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.UnassignLineItemFromConversionEvent(conversionEventId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TrackingPixel on a ConversionEvent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTrackingPixelToConversionEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conversionEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trackingPixelId,_ := strconv.ParseUint( vars["trackingPixelId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent DAO
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.AssignTrackingPixelToConversionEvent(conversionEventId, trackingPixelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TrackingPixel on a ConversionEvent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTrackingPixelFromConversionEvent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	conversionEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConversionEvent DAO
	//----------------------------------------------------------------------------
	requestResult := ConversionEventDAO.UnassignTrackingPixelFromConversionEvent(conversionEventId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


