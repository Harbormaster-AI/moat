package controller

import (
    ProductionScheduleDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProductionScheduleDAO for database creation
//----------------------------------------------------------------------------
func CreateProductionSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionSchedule model
	//----------------------------------------------------------------------------
	data := model.ProductionSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.CreateProductionSchedule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProductionScheduleDAO to find the relevant ProductionSchedule
//----------------------------------------------------------------------------
func GetProductionSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionSchedule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.GetProductionSchedule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProductionScheduleDAO for database read of all ProductionSchedules
//----------------------------------------------------------------------------
func GetAllProductionSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.GetAllProductionSchedule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProductionScheduleDAO for database save
//----------------------------------------------------------------------------
func UpdateProductionSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionSchedule model
	//----------------------------------------------------------------------------
	var data = model.ProductionSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.UpdateProductionSchedule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProductionScheduleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProductionSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionSchedule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProductionScheduleDAO.DeleteProductionSchedule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Plant on a ProductionSchedule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToProductionSchedule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.AssignPlantToProductionSchedule(productionScheduleId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a ProductionSchedule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromProductionSchedule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.UnassignPlantFromProductionSchedule(productionScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more workOrdersIds as a WorkOrders to a ProductionSchedule
	//----------------------------------------------------------------------------
func AddWorkOrdersToProductionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.AddWorkOrdersToProductionSchedule(productionScheduleId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workOrdersIds as a WorkOrders from a ProductionSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkOrdersFromProductionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ProductionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionScheduleDAO.RemoveWorkOrdersFromProductionSchedule(productionScheduleId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
