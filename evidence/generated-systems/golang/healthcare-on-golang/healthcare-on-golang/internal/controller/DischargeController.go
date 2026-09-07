package controller

import (
    DischargeDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DischargeDAO for database creation
//----------------------------------------------------------------------------
func CreateDischarge(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Discharge model
	//----------------------------------------------------------------------------
	data := model.Discharge{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Discharge model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Discharge data access object to create
	//----------------------------------------------------------------------------
	requestResult := DischargeDAO.CreateDischarge( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DischargeDAO to find the relevant Discharge
//----------------------------------------------------------------------------
func GetDischarge(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Discharge data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DischargeDAO.GetDischarge(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DischargeDAO for database read of all Discharges
//----------------------------------------------------------------------------
func GetAllDischarge(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Discharge data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DischargeDAO.GetAllDischarge()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DischargeDAO for database save
//----------------------------------------------------------------------------
func UpdateDischarge(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Discharge model
	//----------------------------------------------------------------------------
	var data = model.Discharge{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Discharge model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Discharge data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DischargeDAO.UpdateDischarge(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DischargeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDischarge(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Discharge data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DischargeDAO.DeleteDischarge(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a Discharge
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToDischarge(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dischargeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Discharge DAO
	//----------------------------------------------------------------------------
	requestResult := DischargeDAO.AssignEncounterToDischarge(dischargeId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a Discharge
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromDischarge( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dischargeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Discharge DAO
	//----------------------------------------------------------------------------
	requestResult := DischargeDAO.UnassignEncounterFromDischarge(dischargeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


