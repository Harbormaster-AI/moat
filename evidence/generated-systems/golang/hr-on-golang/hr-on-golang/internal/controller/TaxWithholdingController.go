package controller

import (
    TaxWithholdingDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TaxWithholdingDAO for database creation
//----------------------------------------------------------------------------
func CreateTaxWithholding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TaxWithholding model
	//----------------------------------------------------------------------------
	data := model.TaxWithholding{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TaxWithholding model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TaxWithholding data access object to create
	//----------------------------------------------------------------------------
	requestResult := TaxWithholdingDAO.CreateTaxWithholding( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TaxWithholdingDAO to find the relevant TaxWithholding
//----------------------------------------------------------------------------
func GetTaxWithholding(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TaxWithholding data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TaxWithholdingDAO.GetTaxWithholding(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TaxWithholdingDAO for database read of all TaxWithholdings
//----------------------------------------------------------------------------
func GetAllTaxWithholding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TaxWithholding data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TaxWithholdingDAO.GetAllTaxWithholding()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TaxWithholdingDAO for database save
//----------------------------------------------------------------------------
func UpdateTaxWithholding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TaxWithholding model
	//----------------------------------------------------------------------------
	var data = model.TaxWithholding{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TaxWithholding model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TaxWithholding data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TaxWithholdingDAO.UpdateTaxWithholding(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TaxWithholdingDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTaxWithholding(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TaxWithholding data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TaxWithholdingDAO.DeleteTaxWithholding(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a TaxWithholding
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToTaxWithholding(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	taxWithholdingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TaxWithholding DAO
	//----------------------------------------------------------------------------
	requestResult := TaxWithholdingDAO.AssignEmployeeToTaxWithholding(taxWithholdingId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a TaxWithholding
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromTaxWithholding( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	taxWithholdingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TaxWithholding DAO
	//----------------------------------------------------------------------------
	requestResult := TaxWithholdingDAO.UnassignEmployeeFromTaxWithholding(taxWithholdingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


