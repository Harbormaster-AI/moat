package controller

import (
    BuildScheduleDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BuildScheduleDAO for database creation
//----------------------------------------------------------------------------
func CreateBuildSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BuildSchedule model
	//----------------------------------------------------------------------------
	data := model.BuildSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BuildSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BuildSchedule data access object to create
	//----------------------------------------------------------------------------
	requestResult := BuildScheduleDAO.CreateBuildSchedule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BuildScheduleDAO to find the relevant BuildSchedule
//----------------------------------------------------------------------------
func GetBuildSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BuildSchedule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BuildScheduleDAO.GetBuildSchedule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BuildScheduleDAO for database read of all BuildSchedules
//----------------------------------------------------------------------------
func GetAllBuildSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BuildSchedule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BuildScheduleDAO.GetAllBuildSchedule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BuildScheduleDAO for database save
//----------------------------------------------------------------------------
func UpdateBuildSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BuildSchedule model
	//----------------------------------------------------------------------------
	var data = model.BuildSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BuildSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BuildSchedule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BuildScheduleDAO.UpdateBuildSchedule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BuildScheduleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBuildSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BuildSchedule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BuildScheduleDAO.DeleteBuildSchedule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more productionOrdersIds as a ProductionOrders to a BuildSchedule
	//----------------------------------------------------------------------------
func AddProductionOrdersToBuildSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	buildScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionOrdersIds,_ := vars["productionOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BuildSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := BuildScheduleDAO.AddProductionOrdersToBuildSchedule(buildScheduleId, productionOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productionOrdersIds as a ProductionOrders from a BuildSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductionOrdersFromBuildSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	buildScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionOrdersIds,_ := vars["productionOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BuildSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := BuildScheduleDAO.RemoveProductionOrdersFromBuildSchedule(buildScheduleId, productionOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
