package controller

import (
    SubscriberDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SubscriberDAO for database creation
//----------------------------------------------------------------------------
func CreateSubscriber(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Subscriber model
	//----------------------------------------------------------------------------
	data := model.Subscriber{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Subscriber model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Subscriber data access object to create
	//----------------------------------------------------------------------------
	requestResult := SubscriberDAO.CreateSubscriber( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SubscriberDAO to find the relevant Subscriber
//----------------------------------------------------------------------------
func GetSubscriber(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Subscriber data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SubscriberDAO.GetSubscriber(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SubscriberDAO for database read of all Subscribers
//----------------------------------------------------------------------------
func GetAllSubscriber(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Subscriber data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SubscriberDAO.GetAllSubscriber()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SubscriberDAO for database save
//----------------------------------------------------------------------------
func UpdateSubscriber(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Subscriber model
	//----------------------------------------------------------------------------
	var data = model.Subscriber{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Subscriber model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Subscriber data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SubscriberDAO.UpdateSubscriber(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SubscriberDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSubscriber(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Subscriber data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SubscriberDAO.DeleteSubscriber(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a Subscriber
	//----------------------------------------------------------------------------
func AddAlertsToSubscriber(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	subscriberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Subscriber DAO
	//----------------------------------------------------------------------------
	requestResult := SubscriberDAO.AddAlertsToSubscriber(subscriberId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a Subscriber
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromSubscriber(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	subscriberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Subscriber DAO
	//----------------------------------------------------------------------------
	requestResult := SubscriberDAO.RemoveAlertsFromSubscriber(subscriberId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
