package controller

import (
    PrivacyNoticeDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PrivacyNoticeDAO for database creation
//----------------------------------------------------------------------------
func CreatePrivacyNotice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PrivacyNotice model
	//----------------------------------------------------------------------------
	data := model.PrivacyNotice{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PrivacyNotice model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice data access object to create
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.CreatePrivacyNotice( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PrivacyNoticeDAO to find the relevant PrivacyNotice
//----------------------------------------------------------------------------
func GetPrivacyNotice(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PrivacyNotice data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.GetPrivacyNotice(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PrivacyNoticeDAO for database read of all PrivacyNotices
//----------------------------------------------------------------------------
func GetAllPrivacyNotice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.GetAllPrivacyNotice()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PrivacyNoticeDAO for database save
//----------------------------------------------------------------------------
func UpdatePrivacyNotice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PrivacyNotice model
	//----------------------------------------------------------------------------
	var data = model.PrivacyNotice{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PrivacyNotice model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.UpdatePrivacyNotice(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PrivacyNoticeDAO for database deletion
//----------------------------------------------------------------------------
func DeletePrivacyNotice(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PrivacyNotice data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PrivacyNoticeDAO.DeletePrivacyNotice(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a PrivacyNotice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToPrivacyNotice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	privacyNoticeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice DAO
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.AssignOrganizationToPrivacyNotice(privacyNoticeId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a PrivacyNotice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromPrivacyNotice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	privacyNoticeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice DAO
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.UnassignOrganizationFromPrivacyNotice(privacyNoticeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a PrivacyNotice
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToPrivacyNotice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	privacyNoticeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice DAO
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.AddProcessingActivitiesToPrivacyNotice(privacyNoticeId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a PrivacyNotice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromPrivacyNotice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	privacyNoticeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice DAO
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.RemoveProcessingActivitiesFromPrivacyNotice(privacyNoticeId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a PrivacyNotice
	//----------------------------------------------------------------------------
func AddConsentsToPrivacyNotice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	privacyNoticeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice DAO
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.AddConsentsToPrivacyNotice(privacyNoticeId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a PrivacyNotice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromPrivacyNotice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	privacyNoticeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PrivacyNotice DAO
	//----------------------------------------------------------------------------
	requestResult := PrivacyNoticeDAO.RemoveConsentsFromPrivacyNotice(privacyNoticeId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
