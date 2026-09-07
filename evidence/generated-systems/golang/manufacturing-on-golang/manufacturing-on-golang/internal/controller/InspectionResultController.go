package controller

import (
    InspectionResultDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InspectionResultDAO for database creation
//----------------------------------------------------------------------------
func CreateInspectionResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionResult model
	//----------------------------------------------------------------------------
	data := model.InspectionResult{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionResult model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult data access object to create
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.CreateInspectionResult( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InspectionResultDAO to find the relevant InspectionResult
//----------------------------------------------------------------------------
func GetInspectionResult(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionResult data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.GetInspectionResult(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InspectionResultDAO for database read of all InspectionResults
//----------------------------------------------------------------------------
func GetAllInspectionResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.GetAllInspectionResult()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InspectionResultDAO for database save
//----------------------------------------------------------------------------
func UpdateInspectionResult(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionResult model
	//----------------------------------------------------------------------------
	var data = model.InspectionResult{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionResult model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.UpdateInspectionResult(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InspectionResultDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInspectionResult(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionResult data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InspectionResultDAO.DeleteInspectionResult(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a InspectionLot on a InspectionResult
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInspectionLotToInspectionResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inspectionLotId,_ := strconv.ParseUint( vars["inspectionLotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.AssignInspectionLotToInspectionResult(inspectionResultId, inspectionLotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InspectionLot on a InspectionResult
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInspectionLotFromInspectionResult( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.UnassignInspectionLotFromInspectionResult(inspectionResultId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Characteristic on a InspectionResult
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCharacteristicToInspectionResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	characteristicId,_ := strconv.ParseUint( vars["characteristicId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.AssignCharacteristicToInspectionResult(inspectionResultId, characteristicId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Characteristic on a InspectionResult
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCharacteristicFromInspectionResult( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionResultId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionResult DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionResultDAO.UnassignCharacteristicFromInspectionResult(inspectionResultId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


