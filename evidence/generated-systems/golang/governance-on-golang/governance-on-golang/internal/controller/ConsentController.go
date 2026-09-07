package controller

import (
    ConsentDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ConsentDAO for database creation
//----------------------------------------------------------------------------
func CreateConsent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Consent model
	//----------------------------------------------------------------------------
	data := model.Consent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Consent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object to create
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.CreateConsent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ConsentDAO to find the relevant Consent
//----------------------------------------------------------------------------
func GetConsent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Consent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.GetConsent(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ConsentDAO for database read of all Consents
//----------------------------------------------------------------------------
func GetAllConsent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.GetAllConsent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ConsentDAO for database save
//----------------------------------------------------------------------------
func UpdateConsent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Consent model
	//----------------------------------------------------------------------------
	var data = model.Consent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Consent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Consent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UpdateConsent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ConsentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteConsent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Consent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ConsentDAO.DeleteConsent(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PrivacyNotice on a Consent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPrivacyNoticeToConsent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	privacyNoticeId,_ := strconv.ParseUint( vars["privacyNoticeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AssignPrivacyNoticeToConsent(consentId, privacyNoticeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PrivacyNotice on a Consent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPrivacyNoticeFromConsent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.UnassignPrivacyNoticeFromConsent(consentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a Consent
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToConsent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.AddProcessingActivitiesToConsent(consentId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a Consent
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromConsent(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	consentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Consent DAO
	//----------------------------------------------------------------------------
	requestResult := ConsentDAO.RemoveProcessingActivitiesFromConsent(consentId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
