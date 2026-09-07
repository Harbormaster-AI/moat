package controller

import (
    MROFacilityDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MROFacilityDAO for database creation
//----------------------------------------------------------------------------
func CreateMROFacility(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MROFacility model
	//----------------------------------------------------------------------------
	data := model.MROFacility{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MROFacility model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MROFacility data access object to create
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.CreateMROFacility( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MROFacilityDAO to find the relevant MROFacility
//----------------------------------------------------------------------------
func GetMROFacility(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MROFacility data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.GetMROFacility(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MROFacilityDAO for database read of all MROFacilitys
//----------------------------------------------------------------------------
func GetAllMROFacility(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MROFacility data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.GetAllMROFacility()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MROFacilityDAO for database save
//----------------------------------------------------------------------------
func UpdateMROFacility(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MROFacility model
	//----------------------------------------------------------------------------
	var data = model.MROFacility{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MROFacility model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MROFacility data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.UpdateMROFacility(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MROFacilityDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMROFacility(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MROFacility data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MROFacilityDAO.DeleteMROFacility(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more appointmentsIds as a Appointments to a MROFacility
	//----------------------------------------------------------------------------
func AddAppointmentsToMROFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	mROFacilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentsIds,_ := vars["appointmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MROFacility DAO
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.AddAppointmentsToMROFacility(mROFacilityId, appointmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more appointmentsIds as a Appointments from a MROFacility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAppointmentsFromMROFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	mROFacilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentsIds,_ := vars["appointmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MROFacility DAO
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.RemoveAppointmentsFromMROFacility(mROFacilityId, appointmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more workOrdersIds as a WorkOrders to a MROFacility
	//----------------------------------------------------------------------------
func AddWorkOrdersToMROFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	mROFacilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MROFacility DAO
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.AddWorkOrdersToMROFacility(mROFacilityId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workOrdersIds as a WorkOrders from a MROFacility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkOrdersFromMROFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	mROFacilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MROFacility DAO
	//----------------------------------------------------------------------------
	requestResult := MROFacilityDAO.RemoveWorkOrdersFromMROFacility(mROFacilityId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
