package controller

import (
    CustomerDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
	// adds one or more enterprisesIds as a Enterprises to a Customer
	//----------------------------------------------------------------------------
func AddEnterprisesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enterprisesIds,_ := vars["enterprisesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddEnterprisesToCustomer(customerId, enterprisesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more enterprisesIds as a Enterprises from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEnterprisesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enterprisesIds,_ := vars["enterprisesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveEnterprisesFromCustomer(customerId, enterprisesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more salesOrdersIds as a SalesOrders to a Customer
	//----------------------------------------------------------------------------
func AddSalesOrdersToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesOrdersIds,_ := vars["salesOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddSalesOrdersToCustomer(customerId, salesOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more salesOrdersIds as a SalesOrders from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSalesOrdersFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesOrdersIds,_ := vars["salesOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveSalesOrdersFromCustomer(customerId, salesOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
