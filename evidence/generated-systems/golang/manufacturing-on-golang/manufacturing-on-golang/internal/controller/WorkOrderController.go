package controller

import (
    WorkOrderDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WorkOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateWorkOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkOrder model
	//----------------------------------------------------------------------------
	data := model.WorkOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.CreateWorkOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WorkOrderDAO to find the relevant WorkOrder
//----------------------------------------------------------------------------
func GetWorkOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.GetWorkOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WorkOrderDAO for database read of all WorkOrders
//----------------------------------------------------------------------------
func GetAllWorkOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.GetAllWorkOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WorkOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateWorkOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkOrder model
	//----------------------------------------------------------------------------
	var data = model.WorkOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UpdateWorkOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WorkOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWorkOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WorkOrderDAO.DeleteWorkOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Item on a WorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.AssignItemToWorkOrder(workOrderId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a WorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UnassignItemFromWorkOrder(workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plant on a WorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.AssignPlantToWorkOrder(workOrderId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a WorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UnassignPlantFromWorkOrder(workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Routing on a WorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRoutingToWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	routingId,_ := strconv.ParseUint( vars["routingId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.AssignRoutingToWorkOrder(workOrderId, routingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Routing on a WorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRoutingFromWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UnassignRoutingFromWorkOrder(workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Bom on a WorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBomToWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bomId,_ := strconv.ParseUint( vars["bomId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.AssignBomToWorkOrder(workOrderId, bomId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bom on a WorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBomFromWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UnassignBomFromWorkOrder(workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ProductionSchedule on a WorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductionScheduleToWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionScheduleId,_ := strconv.ParseUint( vars["productionScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.AssignProductionScheduleToWorkOrder(workOrderId, productionScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ProductionSchedule on a WorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductionScheduleFromWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UnassignProductionScheduleFromWorkOrder(workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SalesOrder on a WorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSalesOrderToWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesOrderId,_ := strconv.ParseUint( vars["salesOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.AssignSalesOrderToWorkOrder(workOrderId, salesOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SalesOrder on a WorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSalesOrderFromWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := WorkOrderDAO.UnassignSalesOrderFromWorkOrder(workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


