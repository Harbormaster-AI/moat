package controller

import (
    NonconformanceDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to NonconformanceDAO for database creation
//----------------------------------------------------------------------------
func CreateNonconformance(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Nonconformance model
	//----------------------------------------------------------------------------
	data := model.Nonconformance{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Nonconformance model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance data access object to create
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.CreateNonconformance( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to NonconformanceDAO to find the relevant Nonconformance
//----------------------------------------------------------------------------
func GetNonconformance(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Nonconformance data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.GetNonconformance(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to NonconformanceDAO for database read of all Nonconformances
//----------------------------------------------------------------------------
func GetAllNonconformance(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance data access object to get all
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.GetAllNonconformance()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to NonconformanceDAO for database save
//----------------------------------------------------------------------------
func UpdateNonconformance(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Nonconformance model
	//----------------------------------------------------------------------------
	var data = model.Nonconformance{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Nonconformance model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.UpdateNonconformance(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to NonconformanceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteNonconformance(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Nonconformance data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := NonconformanceDAO.DeleteNonconformance(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Item on a Nonconformance
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToNonconformance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.AssignItemToNonconformance(nonconformanceId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a Nonconformance
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromNonconformance( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.UnassignItemFromNonconformance(nonconformanceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkOrder on a Nonconformance
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkOrderToNonconformance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrderId,_ := strconv.ParseUint( vars["workOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.AssignWorkOrderToNonconformance(nonconformanceId, workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkOrder on a Nonconformance
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkOrderFromNonconformance( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.UnassignWorkOrderFromNonconformance(nonconformanceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InspectionLot on a Nonconformance
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInspectionLotToNonconformance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inspectionLotId,_ := strconv.ParseUint( vars["inspectionLotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.AssignInspectionLotToNonconformance(nonconformanceId, inspectionLotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InspectionLot on a Nonconformance
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInspectionLotFromNonconformance( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.UnassignInspectionLotFromNonconformance(nonconformanceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CorrectiveAction on a Nonconformance
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCorrectiveActionToNonconformance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionId,_ := strconv.ParseUint( vars["correctiveActionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.AssignCorrectiveActionToNonconformance(nonconformanceId, correctiveActionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CorrectiveAction on a Nonconformance
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCorrectiveActionFromNonconformance( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	nonconformanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Nonconformance DAO
	//----------------------------------------------------------------------------
	requestResult := NonconformanceDAO.UnassignCorrectiveActionFromNonconformance(nonconformanceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


