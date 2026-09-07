package controller

import (
    ServiceBulletinDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ServiceBulletinDAO for database creation
//----------------------------------------------------------------------------
func CreateServiceBulletin(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ServiceBulletin model
	//----------------------------------------------------------------------------
	data := model.ServiceBulletin{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ServiceBulletin model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin data access object to create
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.CreateServiceBulletin( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ServiceBulletinDAO to find the relevant ServiceBulletin
//----------------------------------------------------------------------------
func GetServiceBulletin(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ServiceBulletin data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.GetServiceBulletin(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ServiceBulletinDAO for database read of all ServiceBulletins
//----------------------------------------------------------------------------
func GetAllServiceBulletin(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.GetAllServiceBulletin()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ServiceBulletinDAO for database save
//----------------------------------------------------------------------------
func UpdateServiceBulletin(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ServiceBulletin model
	//----------------------------------------------------------------------------
	var data = model.ServiceBulletin{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ServiceBulletin model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.UpdateServiceBulletin(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ServiceBulletinDAO for database deletion
//----------------------------------------------------------------------------
func DeleteServiceBulletin(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ServiceBulletin data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ServiceBulletinDAO.DeleteServiceBulletin(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more workOrdersIds as a WorkOrders to a ServiceBulletin
	//----------------------------------------------------------------------------
func AddWorkOrdersToServiceBulletin(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	serviceBulletinId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin DAO
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.AddWorkOrdersToServiceBulletin(serviceBulletinId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workOrdersIds as a WorkOrders from a ServiceBulletin
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkOrdersFromServiceBulletin(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	serviceBulletinId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin DAO
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.RemoveWorkOrdersFromServiceBulletin(serviceBulletinId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a ServiceBulletin
	//----------------------------------------------------------------------------
func AddVariantsToServiceBulletin(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	serviceBulletinId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin DAO
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.AddVariantsToServiceBulletin(serviceBulletinId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a ServiceBulletin
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromServiceBulletin(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	serviceBulletinId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ServiceBulletin DAO
	//----------------------------------------------------------------------------
	requestResult := ServiceBulletinDAO.RemoveVariantsFromServiceBulletin(serviceBulletinId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
