package controller

import (
    DeviceCriterionDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DeviceCriterionDAO for database creation
//----------------------------------------------------------------------------
func CreateDeviceCriterion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeviceCriterion model
	//----------------------------------------------------------------------------
	data := model.DeviceCriterion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DeviceCriterion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceCriterion data access object to create
	//----------------------------------------------------------------------------
	requestResult := DeviceCriterionDAO.CreateDeviceCriterion( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DeviceCriterionDAO to find the relevant DeviceCriterion
//----------------------------------------------------------------------------
func GetDeviceCriterion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DeviceCriterion data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DeviceCriterionDAO.GetDeviceCriterion(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DeviceCriterionDAO for database read of all DeviceCriterions
//----------------------------------------------------------------------------
func GetAllDeviceCriterion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DeviceCriterion data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DeviceCriterionDAO.GetAllDeviceCriterion()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DeviceCriterionDAO for database save
//----------------------------------------------------------------------------
func UpdateDeviceCriterion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeviceCriterion model
	//----------------------------------------------------------------------------
	var data = model.DeviceCriterion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DeviceCriterion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceCriterion data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DeviceCriterionDAO.UpdateDeviceCriterion(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DeviceCriterionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDeviceCriterion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DeviceCriterion data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DeviceCriterionDAO.DeleteDeviceCriterion(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a TargetingProfile on a DeviceCriterion
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTargetingProfileToDeviceCriterion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	deviceCriterionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	targetingProfileId,_ := strconv.ParseUint( vars["targetingProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceCriterion DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceCriterionDAO.AssignTargetingProfileToDeviceCriterion(deviceCriterionId, targetingProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TargetingProfile on a DeviceCriterion
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTargetingProfileFromDeviceCriterion( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	deviceCriterionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceCriterion DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceCriterionDAO.UnassignTargetingProfileFromDeviceCriterion(deviceCriterionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


