package controller

import (
    ExposureDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExposureDAO for database creation
//----------------------------------------------------------------------------
func CreateExposure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Exposure model
	//----------------------------------------------------------------------------
	data := model.Exposure{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Exposure model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.CreateExposure( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExposureDAO to find the relevant Exposure
//----------------------------------------------------------------------------
func GetExposure(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Exposure data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.GetExposure(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExposureDAO for database read of all Exposures
//----------------------------------------------------------------------------
func GetAllExposure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Exposure data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.GetAllExposure()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExposureDAO for database save
//----------------------------------------------------------------------------
func UpdateExposure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Exposure model
	//----------------------------------------------------------------------------
	var data = model.Exposure{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Exposure model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.UpdateExposure(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExposureDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExposure(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Exposure data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExposureDAO.DeleteExposure(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Claim on a Exposure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClaimToExposure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimId,_ := strconv.ParseUint( vars["claimId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.AssignClaimToExposure(exposureId, claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Claim on a Exposure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClaimFromExposure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.UnassignClaimFromExposure(exposureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PolicyCoverage on a Exposure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyCoverageToExposure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyCoverageId,_ := strconv.ParseUint( vars["policyCoverageId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.AssignPolicyCoverageToExposure(exposureId, policyCoverageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PolicyCoverage on a Exposure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyCoverageFromExposure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.UnassignPolicyCoverageFromExposure(exposureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InsuredObject on a Exposure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInsuredObjectToExposure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectId,_ := strconv.ParseUint( vars["insuredObjectId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.AssignInsuredObjectToExposure(exposureId, insuredObjectId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InsuredObject on a Exposure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInsuredObjectFromExposure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.UnassignInsuredObjectFromExposure(exposureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more reservesIds as a Reserves to a Exposure
	//----------------------------------------------------------------------------
func AddReservesToExposure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservesIds,_ := vars["reservesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.AddReservesToExposure(exposureId, reservesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reservesIds as a Reserves from a Exposure
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReservesFromExposure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservesIds,_ := vars["reservesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.RemoveReservesFromExposure(exposureId, reservesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentsIds as a Payments to a Exposure
	//----------------------------------------------------------------------------
func AddPaymentsToExposure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.AddPaymentsToExposure(exposureId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentsIds as a Payments from a Exposure
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentsFromExposure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exposureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Exposure DAO
	//----------------------------------------------------------------------------
	requestResult := ExposureDAO.RemovePaymentsFromExposure(exposureId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
