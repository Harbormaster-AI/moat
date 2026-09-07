package controller

import (
    SoftwareLoadDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SoftwareLoadDAO for database creation
//----------------------------------------------------------------------------
func CreateSoftwareLoad(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SoftwareLoad model
	//----------------------------------------------------------------------------
	data := model.SoftwareLoad{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SoftwareLoad model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad data access object to create
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.CreateSoftwareLoad( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SoftwareLoadDAO to find the relevant SoftwareLoad
//----------------------------------------------------------------------------
func GetSoftwareLoad(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SoftwareLoad data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.GetSoftwareLoad(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SoftwareLoadDAO for database read of all SoftwareLoads
//----------------------------------------------------------------------------
func GetAllSoftwareLoad(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.GetAllSoftwareLoad()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SoftwareLoadDAO for database save
//----------------------------------------------------------------------------
func UpdateSoftwareLoad(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SoftwareLoad model
	//----------------------------------------------------------------------------
	var data = model.SoftwareLoad{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SoftwareLoad model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.UpdateSoftwareLoad(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SoftwareLoadDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSoftwareLoad(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SoftwareLoad data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SoftwareLoadDAO.DeleteSoftwareLoad(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ConnectedAircraft on a SoftwareLoad
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConnectedAircraftToSoftwareLoad(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	softwareLoadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectedAircraftId,_ := strconv.ParseUint( vars["connectedAircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad DAO
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.AssignConnectedAircraftToSoftwareLoad(softwareLoadId, connectedAircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConnectedAircraft on a SoftwareLoad
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConnectedAircraftFromSoftwareLoad( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	softwareLoadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad DAO
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.UnassignConnectedAircraftFromSoftwareLoad(softwareLoadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AvionicsSuite on a SoftwareLoad
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAvionicsSuiteToSoftwareLoad(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	softwareLoadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	avionicsSuiteId,_ := strconv.ParseUint( vars["avionicsSuiteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad DAO
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.AssignAvionicsSuiteToSoftwareLoad(softwareLoadId, avionicsSuiteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AvionicsSuite on a SoftwareLoad
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAvionicsSuiteFromSoftwareLoad( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	softwareLoadId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareLoad DAO
	//----------------------------------------------------------------------------
	requestResult := SoftwareLoadDAO.UnassignAvionicsSuiteFromSoftwareLoad(softwareLoadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


