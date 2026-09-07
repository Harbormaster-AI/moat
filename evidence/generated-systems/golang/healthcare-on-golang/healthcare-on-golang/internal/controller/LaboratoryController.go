package controller

import (
    LaboratoryDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LaboratoryDAO for database creation
//----------------------------------------------------------------------------
func CreateLaboratory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Laboratory model
	//----------------------------------------------------------------------------
	data := model.Laboratory{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Laboratory model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory data access object to create
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.CreateLaboratory( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LaboratoryDAO to find the relevant Laboratory
//----------------------------------------------------------------------------
func GetLaboratory(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Laboratory data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.GetLaboratory(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LaboratoryDAO for database read of all Laboratorys
//----------------------------------------------------------------------------
func GetAllLaboratory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Laboratory data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.GetAllLaboratory()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LaboratoryDAO for database save
//----------------------------------------------------------------------------
func UpdateLaboratory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Laboratory model
	//----------------------------------------------------------------------------
	var data = model.Laboratory{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Laboratory model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.UpdateLaboratory(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LaboratoryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLaboratory(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Laboratory data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LaboratoryDAO.DeleteLaboratory(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Facility on a Laboratory
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToLaboratory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	laboratoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.AssignFacilityToLaboratory(laboratoryId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a Laboratory
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromLaboratory( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	laboratoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.UnassignFacilityFromLaboratory(laboratoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more laboratoryOrdersIds as a LaboratoryOrders to a Laboratory
	//----------------------------------------------------------------------------
func AddLaboratoryOrdersToLaboratory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	laboratoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryOrdersIds,_ := vars["laboratoryOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.AddLaboratoryOrdersToLaboratory(laboratoryId, laboratoryOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more laboratoryOrdersIds as a LaboratoryOrders from a Laboratory
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLaboratoryOrdersFromLaboratory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	laboratoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryOrdersIds,_ := vars["laboratoryOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.RemoveLaboratoryOrdersFromLaboratory(laboratoryId, laboratoryOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more labResultsIds as a LabResults to a Laboratory
	//----------------------------------------------------------------------------
func AddLabResultsToLaboratory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	laboratoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	labResultsIds,_ := vars["labResultsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.AddLabResultsToLaboratory(laboratoryId, labResultsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more labResultsIds as a LabResults from a Laboratory
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLabResultsFromLaboratory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	laboratoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	labResultsIds,_ := vars["labResultsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Laboratory DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryDAO.RemoveLabResultsFromLaboratory(laboratoryId, labResultsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
