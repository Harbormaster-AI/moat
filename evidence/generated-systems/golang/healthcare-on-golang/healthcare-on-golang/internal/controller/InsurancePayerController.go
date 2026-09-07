package controller

import (
    InsurancePayerDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InsurancePayerDAO for database creation
//----------------------------------------------------------------------------
func CreateInsurancePayer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsurancePayer model
	//----------------------------------------------------------------------------
	data := model.InsurancePayer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsurancePayer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer data access object to create
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.CreateInsurancePayer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InsurancePayerDAO to find the relevant InsurancePayer
//----------------------------------------------------------------------------
func GetInsurancePayer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsurancePayer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.GetInsurancePayer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InsurancePayerDAO for database read of all InsurancePayers
//----------------------------------------------------------------------------
func GetAllInsurancePayer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.GetAllInsurancePayer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InsurancePayerDAO for database save
//----------------------------------------------------------------------------
func UpdateInsurancePayer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsurancePayer model
	//----------------------------------------------------------------------------
	var data = model.InsurancePayer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsurancePayer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.UpdateInsurancePayer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InsurancePayerDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInsurancePayer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsurancePayer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InsurancePayerDAO.DeleteInsurancePayer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more plansIds as a Plans to a InsurancePayer
	//----------------------------------------------------------------------------
func AddPlansToInsurancePayer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurancePayerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plansIds,_ := vars["plansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.AddPlansToInsurancePayer(insurancePayerId, plansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more plansIds as a Plans from a InsurancePayer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlansFromInsurancePayer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurancePayerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plansIds,_ := vars["plansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.RemovePlansFromInsurancePayer(insurancePayerId, plansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a InsurancePayer
	//----------------------------------------------------------------------------
func AddClaimsToInsurancePayer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurancePayerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.AddClaimsToInsurancePayer(insurancePayerId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a InsurancePayer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromInsurancePayer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurancePayerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsurancePayer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurancePayerDAO.RemoveClaimsFromInsurancePayer(insurancePayerId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
