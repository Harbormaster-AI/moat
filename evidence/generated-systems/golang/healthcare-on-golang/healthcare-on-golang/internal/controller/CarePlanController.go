package controller

import (
    CarePlanDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CarePlanDAO for database creation
//----------------------------------------------------------------------------
func CreateCarePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CarePlan model
	//----------------------------------------------------------------------------
	data := model.CarePlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CarePlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.CreateCarePlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CarePlanDAO to find the relevant CarePlan
//----------------------------------------------------------------------------
func GetCarePlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CarePlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.GetCarePlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CarePlanDAO for database read of all CarePlans
//----------------------------------------------------------------------------
func GetAllCarePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CarePlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.GetAllCarePlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CarePlanDAO for database save
//----------------------------------------------------------------------------
func UpdateCarePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CarePlan model
	//----------------------------------------------------------------------------
	var data = model.CarePlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CarePlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.UpdateCarePlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CarePlanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCarePlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CarePlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CarePlanDAO.DeleteCarePlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a CarePlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToCarePlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.AssignPatientToCarePlan(carePlanId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a CarePlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromCarePlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.UnassignPatientFromCarePlan(carePlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CareTeam on a CarePlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCareTeamToCarePlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamId,_ := strconv.ParseUint( vars["careTeamId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.AssignCareTeamToCarePlan(carePlanId, careTeamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CareTeam on a CarePlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCareTeamFromCarePlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.UnassignCareTeamFromCarePlan(carePlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more encountersIds as a Encounters to a CarePlan
	//----------------------------------------------------------------------------
func AddEncountersToCarePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encountersIds,_ := vars["encountersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.AddEncountersToCarePlan(carePlanId, encountersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more encountersIds as a Encounters from a CarePlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEncountersFromCarePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encountersIds,_ := vars["encountersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.RemoveEncountersFromCarePlan(carePlanId, encountersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tasksIds as a Tasks to a CarePlan
	//----------------------------------------------------------------------------
func AddTasksToCarePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tasksIds,_ := vars["tasksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.AddTasksToCarePlan(carePlanId, tasksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tasksIds as a Tasks from a CarePlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTasksFromCarePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	carePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tasksIds,_ := vars["tasksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CarePlan DAO
	//----------------------------------------------------------------------------
	requestResult := CarePlanDAO.RemoveTasksFromCarePlan(carePlanId, tasksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
