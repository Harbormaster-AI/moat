package controller

import (
    InsurancePlanDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InsurancePlanDAO for database creation
//----------------------------------------------------------------------------
func CreateInsurancePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsurancePlan model
	//----------------------------------------------------------------------------
	data := model.InsurancePlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsurancePlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.CreateInsurancePlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InsurancePlanDAO to find the relevant InsurancePlan
//----------------------------------------------------------------------------
func GetInsurancePlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsurancePlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.GetInsurancePlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InsurancePlanDAO for database read of all InsurancePlans
//----------------------------------------------------------------------------
func GetAllInsurancePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.GetAllInsurancePlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InsurancePlanDAO for database save
//----------------------------------------------------------------------------
func UpdateInsurancePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsurancePlan model
	//----------------------------------------------------------------------------
	var data = model.InsurancePlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsurancePlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.UpdateInsurancePlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InsurancePlanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInsurancePlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsurancePlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InsurancePlanDAO.DeleteInsurancePlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Payer on a InsurancePlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPayerToInsurancePlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insurancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payerId,_ := strconv.ParseUint( vars["payerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.AssignPayerToInsurancePlan(insurancePlanId, payerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Payer on a InsurancePlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPayerFromInsurancePlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insurancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.UnassignPayerFromInsurancePlan(insurancePlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more coveragesIds as a Coverages to a InsurancePlan
	//----------------------------------------------------------------------------
func AddCoveragesToInsurancePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.AddCoveragesToInsurancePlan(insurancePlanId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more coveragesIds as a Coverages from a InsurancePlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCoveragesFromInsurancePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePlanDAO.RemoveCoveragesFromInsurancePlan(insurancePlanId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
