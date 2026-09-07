package controller

import (
    TrackingPixelDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TrackingPixelDAO for database creation
//----------------------------------------------------------------------------
func CreateTrackingPixel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrackingPixel model
	//----------------------------------------------------------------------------
	data := model.TrackingPixel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrackingPixel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel data access object to create
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.CreateTrackingPixel( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TrackingPixelDAO to find the relevant TrackingPixel
//----------------------------------------------------------------------------
func GetTrackingPixel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrackingPixel data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.GetTrackingPixel(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TrackingPixelDAO for database read of all TrackingPixels
//----------------------------------------------------------------------------
func GetAllTrackingPixel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.GetAllTrackingPixel()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TrackingPixelDAO for database save
//----------------------------------------------------------------------------
func UpdateTrackingPixel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrackingPixel model
	//----------------------------------------------------------------------------
	var data = model.TrackingPixel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrackingPixel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.UpdateTrackingPixel(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TrackingPixelDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTrackingPixel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrackingPixel data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TrackingPixelDAO.DeleteTrackingPixel(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a TrackingPixel
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToTrackingPixel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trackingPixelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel DAO
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.AssignCampaignToTrackingPixel(trackingPixelId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a TrackingPixel
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromTrackingPixel( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trackingPixelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel DAO
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.UnassignCampaignFromTrackingPixel(trackingPixelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Advertiser on a TrackingPixel
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdvertiserToTrackingPixel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trackingPixelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	advertiserId,_ := strconv.ParseUint( vars["advertiserId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel DAO
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.AssignAdvertiserToTrackingPixel(trackingPixelId, advertiserId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Advertiser on a TrackingPixel
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdvertiserFromTrackingPixel( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trackingPixelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel DAO
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.UnassignAdvertiserFromTrackingPixel(trackingPixelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more conversionEventsIds as a ConversionEvents to a TrackingPixel
	//----------------------------------------------------------------------------
func AddConversionEventsToTrackingPixel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trackingPixelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	conversionEventsIds,_ := vars["conversionEventsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel DAO
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.AddConversionEventsToTrackingPixel(trackingPixelId, conversionEventsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more conversionEventsIds as a ConversionEvents from a TrackingPixel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConversionEventsFromTrackingPixel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trackingPixelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	conversionEventsIds,_ := vars["conversionEventsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrackingPixel DAO
	//----------------------------------------------------------------------------
	requestResult := TrackingPixelDAO.RemoveConversionEventsFromTrackingPixel(trackingPixelId, conversionEventsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
