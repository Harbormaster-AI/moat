package controller

import (
    PlannedOrderDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PlannedOrderDAO for database creation
//----------------------------------------------------------------------------
func CreatePlannedOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PlannedOrder model
	//----------------------------------------------------------------------------
	data := model.PlannedOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PlannedOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.CreatePlannedOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PlannedOrderDAO to find the relevant PlannedOrder
//----------------------------------------------------------------------------
func GetPlannedOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PlannedOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.GetPlannedOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PlannedOrderDAO for database read of all PlannedOrders
//----------------------------------------------------------------------------
func GetAllPlannedOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.GetAllPlannedOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PlannedOrderDAO for database save
//----------------------------------------------------------------------------
func UpdatePlannedOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PlannedOrder model
	//----------------------------------------------------------------------------
	var data = model.PlannedOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PlannedOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.UpdatePlannedOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PlannedOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeletePlannedOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PlannedOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PlannedOrderDAO.DeletePlannedOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a MrpRun on a PlannedOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMrpRunToPlannedOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plannedOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mrpRunId,_ := strconv.ParseUint( vars["mrpRunId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.AssignMrpRunToPlannedOrder(plannedOrderId, mrpRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a MrpRun on a PlannedOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMrpRunFromPlannedOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plannedOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.UnassignMrpRunFromPlannedOrder(plannedOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Item on a PlannedOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToPlannedOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plannedOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.AssignItemToPlannedOrder(plannedOrderId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a PlannedOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromPlannedOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plannedOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.UnassignItemFromPlannedOrder(plannedOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plant on a PlannedOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToPlannedOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plannedOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.AssignPlantToPlannedOrder(plannedOrderId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a PlannedOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromPlannedOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plannedOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PlannedOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PlannedOrderDAO.UnassignPlantFromPlannedOrder(plannedOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


