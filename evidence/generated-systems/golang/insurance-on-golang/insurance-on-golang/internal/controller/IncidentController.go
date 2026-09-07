package controller

import (
    IncidentDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to IncidentDAO for database creation
//----------------------------------------------------------------------------
func CreateIncident(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Incident model
	//----------------------------------------------------------------------------
	data := model.Incident{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Incident model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Incident data access object to create
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.CreateIncident( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to IncidentDAO to find the relevant Incident
//----------------------------------------------------------------------------
func GetIncident(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Incident data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.GetIncident(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to IncidentDAO for database read of all Incidents
//----------------------------------------------------------------------------
func GetAllIncident(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Incident data access object to get all
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.GetAllIncident()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to IncidentDAO for database save
//----------------------------------------------------------------------------
func UpdateIncident(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Incident model
	//----------------------------------------------------------------------------
	var data = model.Incident{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Incident model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Incident data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.UpdateIncident(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to IncidentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteIncident(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Incident data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := IncidentDAO.DeleteIncident(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Claim on a Incident
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClaimToIncident(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	incidentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimId,_ := strconv.ParseUint( vars["claimId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Incident DAO
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.AssignClaimToIncident(incidentId, claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Claim on a Incident
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClaimFromIncident( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	incidentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Incident DAO
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.UnassignClaimFromIncident(incidentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more insuredObjectsIds as a InsuredObjects to a Incident
	//----------------------------------------------------------------------------
func AddInsuredObjectsToIncident(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	incidentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectsIds,_ := vars["insuredObjectsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Incident DAO
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.AddInsuredObjectsToIncident(incidentId, insuredObjectsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more insuredObjectsIds as a InsuredObjects from a Incident
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInsuredObjectsFromIncident(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	incidentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectsIds,_ := vars["insuredObjectsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Incident DAO
	//----------------------------------------------------------------------------
	requestResult := IncidentDAO.RemoveInsuredObjectsFromIncident(incidentId, insuredObjectsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
