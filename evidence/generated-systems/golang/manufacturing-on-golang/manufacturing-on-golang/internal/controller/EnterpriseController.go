package controller

import (
    EnterpriseDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EnterpriseDAO for database creation
//----------------------------------------------------------------------------
func CreateEnterprise(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Enterprise model
	//----------------------------------------------------------------------------
	data := model.Enterprise{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Enterprise model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise data access object to create
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.CreateEnterprise( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EnterpriseDAO to find the relevant Enterprise
//----------------------------------------------------------------------------
func GetEnterprise(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Enterprise data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.GetEnterprise(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EnterpriseDAO for database read of all Enterprises
//----------------------------------------------------------------------------
func GetAllEnterprise(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Enterprise data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.GetAllEnterprise()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EnterpriseDAO for database save
//----------------------------------------------------------------------------
func UpdateEnterprise(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Enterprise model
	//----------------------------------------------------------------------------
	var data = model.Enterprise{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Enterprise model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.UpdateEnterprise(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EnterpriseDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEnterprise(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Enterprise data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EnterpriseDAO.DeleteEnterprise(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more businessUnitsIds as a BusinessUnits to a Enterprise
	//----------------------------------------------------------------------------
func AddBusinessUnitsToEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitsIds,_ := vars["businessUnitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.AddBusinessUnitsToEnterprise(enterpriseId, businessUnitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more businessUnitsIds as a BusinessUnits from a Enterprise
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBusinessUnitsFromEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitsIds,_ := vars["businessUnitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.RemoveBusinessUnitsFromEnterprise(enterpriseId, businessUnitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more plantsIds as a Plants to a Enterprise
	//----------------------------------------------------------------------------
func AddPlantsToEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantsIds,_ := vars["plantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.AddPlantsToEnterprise(enterpriseId, plantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more plantsIds as a Plants from a Enterprise
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlantsFromEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantsIds,_ := vars["plantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.RemovePlantsFromEnterprise(enterpriseId, plantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more suppliersIds as a Suppliers to a Enterprise
	//----------------------------------------------------------------------------
func AddSuppliersToEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.AddSuppliersToEnterprise(enterpriseId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more suppliersIds as a Suppliers from a Enterprise
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSuppliersFromEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.RemoveSuppliersFromEnterprise(enterpriseId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more customersIds as a Customers to a Enterprise
	//----------------------------------------------------------------------------
func AddCustomersToEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.AddCustomersToEnterprise(enterpriseId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more customersIds as a Customers from a Enterprise
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCustomersFromEnterprise(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	enterpriseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Enterprise DAO
	//----------------------------------------------------------------------------
	requestResult := EnterpriseDAO.RemoveCustomersFromEnterprise(enterpriseId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
