package controller

import (
    CompensationPackageDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CompensationPackageDAO for database creation
//----------------------------------------------------------------------------
func CreateCompensationPackage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CompensationPackage model
	//----------------------------------------------------------------------------
	data := model.CompensationPackage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CompensationPackage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage data access object to create
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.CreateCompensationPackage( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CompensationPackageDAO to find the relevant CompensationPackage
//----------------------------------------------------------------------------
func GetCompensationPackage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CompensationPackage data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.GetCompensationPackage(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CompensationPackageDAO for database read of all CompensationPackages
//----------------------------------------------------------------------------
func GetAllCompensationPackage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.GetAllCompensationPackage()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CompensationPackageDAO for database save
//----------------------------------------------------------------------------
func UpdateCompensationPackage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CompensationPackage model
	//----------------------------------------------------------------------------
	var data = model.CompensationPackage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CompensationPackage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.UpdateCompensationPackage(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CompensationPackageDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCompensationPackage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CompensationPackage data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CompensationPackageDAO.DeleteCompensationPackage(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Contract on a CompensationPackage
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContractToCompensationPackage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractId,_ := strconv.ParseUint( vars["contractId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.AssignContractToCompensationPackage(compensationPackageId, contractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contract on a CompensationPackage
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContractFromCompensationPackage( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.UnassignContractFromCompensationPackage(compensationPackageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more salaryComponentsIds as a SalaryComponents to a CompensationPackage
	//----------------------------------------------------------------------------
func AddSalaryComponentsToCompensationPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salaryComponentsIds,_ := vars["salaryComponentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.AddSalaryComponentsToCompensationPackage(compensationPackageId, salaryComponentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more salaryComponentsIds as a SalaryComponents from a CompensationPackage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSalaryComponentsFromCompensationPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salaryComponentsIds,_ := vars["salaryComponentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.RemoveSalaryComponentsFromCompensationPackage(compensationPackageId, salaryComponentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more bonusPlansIds as a BonusPlans to a CompensationPackage
	//----------------------------------------------------------------------------
func AddBonusPlansToCompensationPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bonusPlansIds,_ := vars["bonusPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.AddBonusPlansToCompensationPackage(compensationPackageId, bonusPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more bonusPlansIds as a BonusPlans from a CompensationPackage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBonusPlansFromCompensationPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bonusPlansIds,_ := vars["bonusPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.RemoveBonusPlansFromCompensationPackage(compensationPackageId, bonusPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more equityGrantsIds as a EquityGrants to a CompensationPackage
	//----------------------------------------------------------------------------
func AddEquityGrantsToCompensationPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	equityGrantsIds,_ := vars["equityGrantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.AddEquityGrantsToCompensationPackage(compensationPackageId, equityGrantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more equityGrantsIds as a EquityGrants from a CompensationPackage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEquityGrantsFromCompensationPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	compensationPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	equityGrantsIds,_ := vars["equityGrantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CompensationPackage DAO
	//----------------------------------------------------------------------------
	requestResult := CompensationPackageDAO.RemoveEquityGrantsFromCompensationPackage(compensationPackageId, equityGrantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
