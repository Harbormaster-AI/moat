package controller

import (
    AdAccountDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AdAccountDAO for database creation
//----------------------------------------------------------------------------
func CreateAdAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AdAccount model
	//----------------------------------------------------------------------------
	data := model.AdAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AdAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount data access object to create
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.CreateAdAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AdAccountDAO to find the relevant AdAccount
//----------------------------------------------------------------------------
func GetAdAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AdAccount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.GetAdAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AdAccountDAO for database read of all AdAccounts
//----------------------------------------------------------------------------
func GetAllAdAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AdAccount data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.GetAllAdAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AdAccountDAO for database save
//----------------------------------------------------------------------------
func UpdateAdAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AdAccount model
	//----------------------------------------------------------------------------
	var data = model.AdAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AdAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.UpdateAdAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AdAccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAdAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AdAccount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AdAccountDAO.DeleteAdAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Advertiser on a AdAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdvertiserToAdAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	advertiserId,_ := strconv.ParseUint( vars["advertiserId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.AssignAdvertiserToAdAccount(adAccountId, advertiserId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Advertiser on a AdAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdvertiserFromAdAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.UnassignAdvertiserFromAdAccount(adAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a BillingProfile on a AdAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBillingProfileToAdAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	billingProfileId,_ := strconv.ParseUint( vars["billingProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.AssignBillingProfileToAdAccount(adAccountId, billingProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BillingProfile on a AdAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBillingProfileFromAdAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.UnassignBillingProfileFromAdAccount(adAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Dsp on a AdAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDspToAdAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dspId,_ := strconv.ParseUint( vars["dspId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.AssignDspToAdAccount(adAccountId, dspId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dsp on a AdAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDspFromAdAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.UnassignDspFromAdAccount(adAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more usersIds as a Users to a AdAccount
	//----------------------------------------------------------------------------
func AddUsersToAdAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.AddUsersToAdAccount(adAccountId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usersIds as a Users from a AdAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsersFromAdAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.RemoveUsersFromAdAccount(adAccountId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a AdAccount
	//----------------------------------------------------------------------------
func AddCampaignsToAdAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.AddCampaignsToAdAccount(adAccountId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a AdAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromAdAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.RemoveCampaignsFromAdAccount(adAccountId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more performanceMetricsIds as a PerformanceMetrics to a AdAccount
	//----------------------------------------------------------------------------
func AddPerformanceMetricsToAdAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performanceMetricsIds,_ := vars["performanceMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.AddPerformanceMetricsToAdAccount(adAccountId, performanceMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more performanceMetricsIds as a PerformanceMetrics from a AdAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePerformanceMetricsFromAdAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performanceMetricsIds,_ := vars["performanceMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdAccount DAO
	//----------------------------------------------------------------------------
	requestResult := AdAccountDAO.RemovePerformanceMetricsFromAdAccount(adAccountId, performanceMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
