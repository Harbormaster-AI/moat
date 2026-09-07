package controller

import (
    AudienceSegmentDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AudienceSegmentDAO for database creation
//----------------------------------------------------------------------------
func CreateAudienceSegment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AudienceSegment model
	//----------------------------------------------------------------------------
	data := model.AudienceSegment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AudienceSegment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment data access object to create
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.CreateAudienceSegment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AudienceSegmentDAO to find the relevant AudienceSegment
//----------------------------------------------------------------------------
func GetAudienceSegment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AudienceSegment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.GetAudienceSegment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AudienceSegmentDAO for database read of all AudienceSegments
//----------------------------------------------------------------------------
func GetAllAudienceSegment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.GetAllAudienceSegment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AudienceSegmentDAO for database save
//----------------------------------------------------------------------------
func UpdateAudienceSegment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AudienceSegment model
	//----------------------------------------------------------------------------
	var data = model.AudienceSegment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AudienceSegment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.UpdateAudienceSegment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AudienceSegmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAudienceSegment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AudienceSegment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AudienceSegmentDAO.DeleteAudienceSegment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Provider on a AudienceSegment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProviderToAudienceSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	audienceSegmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	providerId,_ := strconv.ParseUint( vars["providerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment DAO
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.AssignProviderToAudienceSegment(audienceSegmentId, providerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Provider on a AudienceSegment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProviderFromAudienceSegment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	audienceSegmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment DAO
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.UnassignProviderFromAudienceSegment(audienceSegmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a AudienceSegment
	//----------------------------------------------------------------------------
func AddCampaignsToAudienceSegment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	audienceSegmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment DAO
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.AddCampaignsToAudienceSegment(audienceSegmentId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a AudienceSegment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromAudienceSegment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	audienceSegmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AudienceSegment DAO
	//----------------------------------------------------------------------------
	requestResult := AudienceSegmentDAO.RemoveCampaignsFromAudienceSegment(audienceSegmentId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
