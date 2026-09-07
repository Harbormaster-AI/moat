package controller

import (
    AircraftDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraft(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Aircraft model
	//----------------------------------------------------------------------------
	data := model.Aircraft{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Aircraft model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.CreateAircraft( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftDAO to find the relevant Aircraft
//----------------------------------------------------------------------------
func GetAircraft(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Aircraft data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.GetAircraft(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftDAO for database read of all Aircrafts
//----------------------------------------------------------------------------
func GetAllAircraft(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Aircraft data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.GetAllAircraft()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraft(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Aircraft model
	//----------------------------------------------------------------------------
	var data = model.Aircraft{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Aircraft model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UpdateAircraft(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraft(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Aircraft data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftDAO.DeleteAircraft(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Variant on a Aircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignVariantToAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantId,_ := strconv.ParseUint( vars["variantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AssignVariantToAircraft(aircraftId, variantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Variant on a Aircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignVariantFromAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UnassignVariantFromAircraft(aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Operator on a Aircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOperatorToAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatorId,_ := strconv.ParseUint( vars["operatorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AssignOperatorToAircraft(aircraftId, operatorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Operator on a Aircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOperatorFromAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UnassignOperatorFromAircraft(aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Registration on a Aircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRegistrationToAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	registrationId,_ := strconv.ParseUint( vars["registrationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AssignRegistrationToAircraft(aircraftId, registrationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Registration on a Aircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRegistrationFromAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UnassignRegistrationFromAircraft(aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warranty on a Aircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarrantyToAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warrantyId,_ := strconv.ParseUint( vars["warrantyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AssignWarrantyToAircraft(aircraftId, warrantyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warranty on a Aircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarrantyFromAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UnassignWarrantyFromAircraft(aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ConnectedAircraft on a Aircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConnectedAircraftToAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectedAircraftId,_ := strconv.ParseUint( vars["connectedAircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AssignConnectedAircraftToAircraft(aircraftId, connectedAircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConnectedAircraft on a Aircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConnectedAircraftFromAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UnassignConnectedAircraftFromAircraft(aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CabinLayout on a Aircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCabinLayoutToAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cabinLayoutId,_ := strconv.ParseUint( vars["cabinLayoutId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AssignCabinLayoutToAircraft(aircraftId, cabinLayoutId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CabinLayout on a Aircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCabinLayoutFromAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.UnassignCabinLayoutFromAircraft(aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more maintenanceRecordsIds as a MaintenanceRecords to a Aircraft
	//----------------------------------------------------------------------------
func AddMaintenanceRecordsToAircraft(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceRecordsIds,_ := vars["maintenanceRecordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.AddMaintenanceRecordsToAircraft(aircraftId, maintenanceRecordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more maintenanceRecordsIds as a MaintenanceRecords from a Aircraft
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMaintenanceRecordsFromAircraft(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceRecordsIds,_ := vars["maintenanceRecordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Aircraft DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftDAO.RemoveMaintenanceRecordsFromAircraft(aircraftId, maintenanceRecordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
