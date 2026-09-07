package controller

import (
    AdvertiserDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AdvertiserDAO for database creation
//----------------------------------------------------------------------------
func CreateAdvertiser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Advertiser model
	//----------------------------------------------------------------------------
	data := model.Advertiser{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Advertiser model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser data access object to create
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.CreateAdvertiser( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AdvertiserDAO to find the relevant Advertiser
//----------------------------------------------------------------------------
func GetAdvertiser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Advertiser data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.GetAdvertiser(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AdvertiserDAO for database read of all Advertisers
//----------------------------------------------------------------------------
func GetAllAdvertiser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Advertiser data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.GetAllAdvertiser()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AdvertiserDAO for database save
//----------------------------------------------------------------------------
func UpdateAdvertiser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Advertiser model
	//----------------------------------------------------------------------------
	var data = model.Advertiser{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Advertiser model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.UpdateAdvertiser(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AdvertiserDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAdvertiser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Advertiser data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AdvertiserDAO.DeleteAdvertiser(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Agency on a Advertiser
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAgencyToAdvertiser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agencyId,_ := strconv.ParseUint( vars["agencyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.AssignAgencyToAdvertiser(advertiserId, agencyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Agency on a Advertiser
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAgencyFromAdvertiser( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.UnassignAgencyFromAdvertiser(advertiserId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more adAccountsIds as a AdAccounts to a Advertiser
	//----------------------------------------------------------------------------
func AddAdAccountsToAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.AddAdAccountsToAdvertiser(advertiserId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more adAccountsIds as a AdAccounts from a Advertiser
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdAccountsFromAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.RemoveAdAccountsFromAdvertiser(advertiserId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more billingProfilesIds as a BillingProfiles to a Advertiser
	//----------------------------------------------------------------------------
func AddBillingProfilesToAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	billingProfilesIds,_ := vars["billingProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.AddBillingProfilesToAdvertiser(advertiserId, billingProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more billingProfilesIds as a BillingProfiles from a Advertiser
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBillingProfilesFromAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	billingProfilesIds,_ := vars["billingProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.RemoveBillingProfilesFromAdvertiser(advertiserId, billingProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a Advertiser
	//----------------------------------------------------------------------------
func AddCampaignsToAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.AddCampaignsToAdvertiser(advertiserId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a Advertiser
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.RemoveCampaignsFromAdvertiser(advertiserId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more trackingPixelsIds as a TrackingPixels to a Advertiser
	//----------------------------------------------------------------------------
func AddTrackingPixelsToAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trackingPixelsIds,_ := vars["trackingPixelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.AddTrackingPixelsToAdvertiser(advertiserId, trackingPixelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more trackingPixelsIds as a TrackingPixels from a Advertiser
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTrackingPixelsFromAdvertiser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	advertiserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trackingPixelsIds,_ := vars["trackingPixelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Advertiser DAO
	//----------------------------------------------------------------------------
	requestResult := AdvertiserDAO.RemoveTrackingPixelsFromAdvertiser(advertiserId, trackingPixelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
