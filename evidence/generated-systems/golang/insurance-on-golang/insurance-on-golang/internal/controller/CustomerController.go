package controller

import (
    CustomerDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CustomerDAO for database creation
//----------------------------------------------------------------------------
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Customer model
	//----------------------------------------------------------------------------
	data := model.Customer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Customer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object to create
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.CreateCustomer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CustomerDAO to find the relevant Customer
//----------------------------------------------------------------------------
func GetCustomer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Customer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.GetCustomer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CustomerDAO for database read of all Customers
//----------------------------------------------------------------------------
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.GetAllCustomer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CustomerDAO for database save
//----------------------------------------------------------------------------
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Customer model
	//----------------------------------------------------------------------------
	var data = model.Customer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Customer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.UpdateCustomer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CustomerDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Customer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CustomerDAO.DeleteCustomer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more applicationsIds as a Applications to a Customer
	//----------------------------------------------------------------------------
func AddApplicationsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationsIds,_ := vars["applicationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddApplicationsToCustomer(customerId, applicationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more applicationsIds as a Applications from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApplicationsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationsIds,_ := vars["applicationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveApplicationsFromCustomer(customerId, applicationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a Customer
	//----------------------------------------------------------------------------
func AddPoliciesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddPoliciesToCustomer(customerId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemovePoliciesFromCustomer(customerId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a Customer
	//----------------------------------------------------------------------------
func AddClaimsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddClaimsToCustomer(customerId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveClaimsFromCustomer(customerId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more agentsIds as a Agents to a Customer
	//----------------------------------------------------------------------------
func AddAgentsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agentsIds,_ := vars["agentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddAgentsToCustomer(customerId, agentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more agentsIds as a Agents from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAgentsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agentsIds,_ := vars["agentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveAgentsFromCustomer(customerId, agentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more beneficiariesIds as a Beneficiaries to a Customer
	//----------------------------------------------------------------------------
func AddBeneficiariesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	beneficiariesIds,_ := vars["beneficiariesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddBeneficiariesToCustomer(customerId, beneficiariesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more beneficiariesIds as a Beneficiaries from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBeneficiariesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	beneficiariesIds,_ := vars["beneficiariesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveBeneficiariesFromCustomer(customerId, beneficiariesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
