package controller

import (
    LaboratoryOrderDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LaboratoryOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateLaboratoryOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LaboratoryOrder model
	//----------------------------------------------------------------------------
	data := model.LaboratoryOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LaboratoryOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.CreateLaboratoryOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LaboratoryOrderDAO to find the relevant LaboratoryOrder
//----------------------------------------------------------------------------
func GetLaboratoryOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LaboratoryOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.GetLaboratoryOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LaboratoryOrderDAO for database read of all LaboratoryOrders
//----------------------------------------------------------------------------
func GetAllLaboratoryOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.GetAllLaboratoryOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LaboratoryOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateLaboratoryOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LaboratoryOrder model
	//----------------------------------------------------------------------------
	var data = model.LaboratoryOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LaboratoryOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.UpdateLaboratoryOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LaboratoryOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLaboratoryOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LaboratoryOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LaboratoryOrderDAO.DeleteLaboratoryOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Order on a LaboratoryOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderToLaboratoryOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	laboratoryOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderId,_ := strconv.ParseUint( vars["orderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.AssignOrderToLaboratoryOrder(laboratoryOrderId, orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Order on a LaboratoryOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderFromLaboratoryOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	laboratoryOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.UnassignOrderFromLaboratoryOrder(laboratoryOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Laboratory on a LaboratoryOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLaboratoryToLaboratoryOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	laboratoryOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryId,_ := strconv.ParseUint( vars["laboratoryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.AssignLaboratoryToLaboratoryOrder(laboratoryOrderId, laboratoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Laboratory on a LaboratoryOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLaboratoryFromLaboratoryOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	laboratoryOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.UnassignLaboratoryFromLaboratoryOrder(laboratoryOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more resultsIds as a Results to a LaboratoryOrder
	//----------------------------------------------------------------------------
func AddResultsToLaboratoryOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	laboratoryOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	resultsIds,_ := vars["resultsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.AddResultsToLaboratoryOrder(laboratoryOrderId, resultsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more resultsIds as a Results from a LaboratoryOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveResultsFromLaboratoryOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	laboratoryOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	resultsIds,_ := vars["resultsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LaboratoryOrder DAO
	//----------------------------------------------------------------------------
	requestResult := LaboratoryOrderDAO.RemoveResultsFromLaboratoryOrder(laboratoryOrderId, resultsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
