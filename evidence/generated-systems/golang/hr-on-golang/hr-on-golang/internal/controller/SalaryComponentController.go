package controller

import (
    SalaryComponentDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SalaryComponentDAO for database creation
//----------------------------------------------------------------------------
func CreateSalaryComponent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalaryComponent model
	//----------------------------------------------------------------------------
	data := model.SalaryComponent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalaryComponent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalaryComponent data access object to create
	//----------------------------------------------------------------------------
	requestResult := SalaryComponentDAO.CreateSalaryComponent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SalaryComponentDAO to find the relevant SalaryComponent
//----------------------------------------------------------------------------
func GetSalaryComponent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalaryComponent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalaryComponentDAO.GetSalaryComponent(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SalaryComponentDAO for database read of all SalaryComponents
//----------------------------------------------------------------------------
func GetAllSalaryComponent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SalaryComponent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SalaryComponentDAO.GetAllSalaryComponent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SalaryComponentDAO for database save
//----------------------------------------------------------------------------
func UpdateSalaryComponent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalaryComponent model
	//----------------------------------------------------------------------------
	var data = model.SalaryComponent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalaryComponent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalaryComponent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalaryComponentDAO.UpdateSalaryComponent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SalaryComponentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSalaryComponent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalaryComponent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SalaryComponentDAO.DeleteSalaryComponent(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CompensationPackage on a SalaryComponent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCompensationPackageToSalaryComponent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salaryComponentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compensationPackageId,_ := strconv.ParseUint( vars["compensationPackageId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalaryComponent DAO
	//----------------------------------------------------------------------------
	requestResult := SalaryComponentDAO.AssignCompensationPackageToSalaryComponent(salaryComponentId, compensationPackageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CompensationPackage on a SalaryComponent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCompensationPackageFromSalaryComponent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salaryComponentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalaryComponent DAO
	//----------------------------------------------------------------------------
	requestResult := SalaryComponentDAO.UnassignCompensationPackageFromSalaryComponent(salaryComponentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


