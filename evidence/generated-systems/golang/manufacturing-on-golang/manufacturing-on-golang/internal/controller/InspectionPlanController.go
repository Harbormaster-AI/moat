package controller

import (
    InspectionPlanDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InspectionPlanDAO for database creation
//----------------------------------------------------------------------------
func CreateInspectionPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionPlan model
	//----------------------------------------------------------------------------
	data := model.InspectionPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.CreateInspectionPlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InspectionPlanDAO to find the relevant InspectionPlan
//----------------------------------------------------------------------------
func GetInspectionPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionPlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.GetInspectionPlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InspectionPlanDAO for database read of all InspectionPlans
//----------------------------------------------------------------------------
func GetAllInspectionPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.GetAllInspectionPlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InspectionPlanDAO for database save
//----------------------------------------------------------------------------
func UpdateInspectionPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionPlan model
	//----------------------------------------------------------------------------
	var data = model.InspectionPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.UpdateInspectionPlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InspectionPlanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInspectionPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionPlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InspectionPlanDAO.DeleteInspectionPlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Item on a InspectionPlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToInspectionPlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.AssignItemToInspectionPlan(inspectionPlanId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a InspectionPlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromInspectionPlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.UnassignItemFromInspectionPlan(inspectionPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more characteristicsIds as a Characteristics to a InspectionPlan
	//----------------------------------------------------------------------------
func AddCharacteristicsToInspectionPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inspectionPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	characteristicsIds,_ := vars["characteristicsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.AddCharacteristicsToInspectionPlan(inspectionPlanId, characteristicsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more characteristicsIds as a Characteristics from a InspectionPlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCharacteristicsFromInspectionPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inspectionPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	characteristicsIds,_ := vars["characteristicsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InspectionPlan DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionPlanDAO.RemoveCharacteristicsFromInspectionPlan(inspectionPlanId, characteristicsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
