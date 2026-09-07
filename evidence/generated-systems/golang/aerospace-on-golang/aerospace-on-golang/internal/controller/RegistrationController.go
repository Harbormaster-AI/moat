package controller

import (
    RegistrationDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RegistrationDAO for database creation
//----------------------------------------------------------------------------
func CreateRegistration(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Registration model
	//----------------------------------------------------------------------------
	data := model.Registration{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Registration model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Registration data access object to create
	//----------------------------------------------------------------------------
	requestResult := RegistrationDAO.CreateRegistration( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RegistrationDAO to find the relevant Registration
//----------------------------------------------------------------------------
func GetRegistration(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Registration data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RegistrationDAO.GetRegistration(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RegistrationDAO for database read of all Registrations
//----------------------------------------------------------------------------
func GetAllRegistration(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Registration data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RegistrationDAO.GetAllRegistration()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RegistrationDAO for database save
//----------------------------------------------------------------------------
func UpdateRegistration(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Registration model
	//----------------------------------------------------------------------------
	var data = model.Registration{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Registration model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Registration data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RegistrationDAO.UpdateRegistration(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RegistrationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRegistration(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Registration data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RegistrationDAO.DeleteRegistration(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Aircraft on a Registration
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftToRegistration(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	registrationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftId,_ := strconv.ParseUint( vars["aircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Registration DAO
	//----------------------------------------------------------------------------
	requestResult := RegistrationDAO.AssignAircraftToRegistration(registrationId, aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Aircraft on a Registration
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftFromRegistration( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	registrationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Registration DAO
	//----------------------------------------------------------------------------
	requestResult := RegistrationDAO.UnassignAircraftFromRegistration(registrationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


