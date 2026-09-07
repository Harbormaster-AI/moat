package controller

import (
    LabResultDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LabResultDAO for database creation
//----------------------------------------------------------------------------
func CreateLabResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LabResult model
	//----------------------------------------------------------------------------
	data := model.LabResult{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LabResult model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LabResult data access object to create
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.CreateLabResult( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LabResultDAO to find the relevant LabResult
//----------------------------------------------------------------------------
func GetLabResult(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LabResult data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.GetLabResult(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LabResultDAO for database read of all LabResults
//----------------------------------------------------------------------------
func GetAllLabResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LabResult data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.GetAllLabResult()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LabResultDAO for database save
//----------------------------------------------------------------------------
func UpdateLabResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LabResult model
	//----------------------------------------------------------------------------
	var data = model.LabResult{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LabResult model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LabResult data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.UpdateLabResult(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LabResultDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLabResult(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LabResult data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LabResultDAO.DeleteLabResult(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a LaboratoryOrder on a LabResult
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLaboratoryOrderToLabResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	labResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryOrderId,_ := strconv.ParseUint( vars["laboratoryOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LabResult DAO
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.AssignLaboratoryOrderToLabResult(labResultId, laboratoryOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LaboratoryOrder on a LabResult
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLaboratoryOrderFromLabResult( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	labResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LabResult DAO
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.UnassignLaboratoryOrderFromLabResult(labResultId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Laboratory on a LabResult
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLaboratoryToLabResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	labResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryId,_ := strconv.ParseUint( vars["laboratoryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LabResult DAO
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.AssignLaboratoryToLabResult(labResultId, laboratoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Laboratory on a LabResult
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLaboratoryFromLabResult( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	labResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LabResult DAO
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.UnassignLaboratoryFromLabResult(labResultId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more observationsIds as a Observations to a LabResult
	//----------------------------------------------------------------------------
func AddObservationsToLabResult(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	labResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LabResult DAO
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.AddObservationsToLabResult(labResultId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more observationsIds as a Observations from a LabResult
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObservationsFromLabResult(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	labResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LabResult DAO
	//----------------------------------------------------------------------------
	requestResult := LabResultDAO.RemoveObservationsFromLabResult(labResultId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
