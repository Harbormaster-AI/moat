package controller

import (
    InspectionCharacteristicDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InspectionCharacteristicDAO for database creation
//----------------------------------------------------------------------------
func CreateInspectionCharacteristic(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionCharacteristic model
	//----------------------------------------------------------------------------
	data := model.InspectionCharacteristic{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionCharacteristic model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionCharacteristic data access object to create
	//----------------------------------------------------------------------------
	requestResult := InspectionCharacteristicDAO.CreateInspectionCharacteristic( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InspectionCharacteristicDAO to find the relevant InspectionCharacteristic
//----------------------------------------------------------------------------
func GetInspectionCharacteristic(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionCharacteristic data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionCharacteristicDAO.GetInspectionCharacteristic(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InspectionCharacteristicDAO for database read of all InspectionCharacteristics
//----------------------------------------------------------------------------
func GetAllInspectionCharacteristic(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InspectionCharacteristic data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InspectionCharacteristicDAO.GetAllInspectionCharacteristic()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InspectionCharacteristicDAO for database save
//----------------------------------------------------------------------------
func UpdateInspectionCharacteristic(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionCharacteristic model
	//----------------------------------------------------------------------------
	var data = model.InspectionCharacteristic{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionCharacteristic model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionCharacteristic data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionCharacteristicDAO.UpdateInspectionCharacteristic(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InspectionCharacteristicDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInspectionCharacteristic(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionCharacteristic data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InspectionCharacteristicDAO.DeleteInspectionCharacteristic(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a InspectionPlan on a InspectionCharacteristic
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInspectionPlanToInspectionCharacteristic(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionCharacteristicId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inspectionPlanId,_ := strconv.ParseUint( vars["inspectionPlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionCharacteristic DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionCharacteristicDAO.AssignInspectionPlanToInspectionCharacteristic(inspectionCharacteristicId, inspectionPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InspectionPlan on a InspectionCharacteristic
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInspectionPlanFromInspectionCharacteristic( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionCharacteristicId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionCharacteristic DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionCharacteristicDAO.UnassignInspectionPlanFromInspectionCharacteristic(inspectionCharacteristicId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


