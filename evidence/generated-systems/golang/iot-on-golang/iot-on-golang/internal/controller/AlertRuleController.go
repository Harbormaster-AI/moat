package controller

import (
    AlertRuleDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AlertRuleDAO for database creation
//----------------------------------------------------------------------------
func CreateAlertRule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AlertRule model
	//----------------------------------------------------------------------------
	data := model.AlertRule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AlertRule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule data access object to create
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.CreateAlertRule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AlertRuleDAO to find the relevant AlertRule
//----------------------------------------------------------------------------
func GetAlertRule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AlertRule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.GetAlertRule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AlertRuleDAO for database read of all AlertRules
//----------------------------------------------------------------------------
func GetAllAlertRule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AlertRule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.GetAllAlertRule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AlertRuleDAO for database save
//----------------------------------------------------------------------------
func UpdateAlertRule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AlertRule model
	//----------------------------------------------------------------------------
	var data = model.AlertRule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AlertRule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.UpdateAlertRule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AlertRuleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAlertRule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AlertRule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AlertRuleDAO.DeleteAlertRule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a AlertRule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToAlertRule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule DAO
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.AssignTenantToAlertRule(alertRuleId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a AlertRule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromAlertRule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	alertRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule DAO
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.UnassignTenantFromAlertRule(alertRuleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more streamsIds as a Streams to a AlertRule
	//----------------------------------------------------------------------------
func AddStreamsToAlertRule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	streamsIds,_ := vars["streamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule DAO
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.AddStreamsToAlertRule(alertRuleId, streamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more streamsIds as a Streams from a AlertRule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStreamsFromAlertRule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	streamsIds,_ := vars["streamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule DAO
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.RemoveStreamsFromAlertRule(alertRuleId, streamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a AlertRule
	//----------------------------------------------------------------------------
func AddAlertsToAlertRule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule DAO
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.AddAlertsToAlertRule(alertRuleId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a AlertRule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromAlertRule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	alertRuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AlertRule DAO
	//----------------------------------------------------------------------------
	requestResult := AlertRuleDAO.RemoveAlertsFromAlertRule(alertRuleId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
