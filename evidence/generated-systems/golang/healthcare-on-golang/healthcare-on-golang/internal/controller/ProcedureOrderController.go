package controller

import (
    ProcedureOrderDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProcedureOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateProcedureOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProcedureOrder model
	//----------------------------------------------------------------------------
	data := model.ProcedureOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProcedureOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.CreateProcedureOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProcedureOrderDAO to find the relevant ProcedureOrder
//----------------------------------------------------------------------------
func GetProcedureOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProcedureOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.GetProcedureOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProcedureOrderDAO for database read of all ProcedureOrders
//----------------------------------------------------------------------------
func GetAllProcedureOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.GetAllProcedureOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProcedureOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateProcedureOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProcedureOrder model
	//----------------------------------------------------------------------------
	var data = model.ProcedureOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProcedureOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.UpdateProcedureOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProcedureOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProcedureOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProcedureOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProcedureOrderDAO.DeleteProcedureOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Order on a ProcedureOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderToProcedureOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderId,_ := strconv.ParseUint( vars["orderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.AssignOrderToProcedureOrder(procedureOrderId, orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Order on a ProcedureOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderFromProcedureOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.UnassignOrderFromProcedureOrder(procedureOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Facility on a ProcedureOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToProcedureOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.AssignFacilityToProcedureOrder(procedureOrderId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a ProcedureOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromProcedureOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.UnassignFacilityFromProcedureOrder(procedureOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Procedure on a ProcedureOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProcedureToProcedureOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	procedureId,_ := strconv.ParseUint( vars["procedureId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.AssignProcedureToProcedureOrder(procedureOrderId, procedureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Procedure on a ProcedureOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProcedureFromProcedureOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	procedureOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProcedureOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProcedureOrderDAO.UnassignProcedureFromProcedureOrder(procedureOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


