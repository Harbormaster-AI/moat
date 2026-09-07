package controller

import (
    OnboardingTaskDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OnboardingTaskDAO for database creation
//----------------------------------------------------------------------------
func CreateOnboardingTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OnboardingTask model
	//----------------------------------------------------------------------------
	data := model.OnboardingTask{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OnboardingTask model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask data access object to create
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.CreateOnboardingTask( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OnboardingTaskDAO to find the relevant OnboardingTask
//----------------------------------------------------------------------------
func GetOnboardingTask(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OnboardingTask data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.GetOnboardingTask(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OnboardingTaskDAO for database read of all OnboardingTasks
//----------------------------------------------------------------------------
func GetAllOnboardingTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.GetAllOnboardingTask()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OnboardingTaskDAO for database save
//----------------------------------------------------------------------------
func UpdateOnboardingTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OnboardingTask model
	//----------------------------------------------------------------------------
	var data = model.OnboardingTask{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OnboardingTask model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.UpdateOnboardingTask(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OnboardingTaskDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOnboardingTask(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OnboardingTask data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OnboardingTaskDAO.DeleteOnboardingTask(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a OnboardingTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToOnboardingTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.AssignEmployeeToOnboardingTask(onboardingTaskId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a OnboardingTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromOnboardingTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.UnassignEmployeeFromOnboardingTask(onboardingTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AssignedTo on a OnboardingTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAssignedToToOnboardingTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignedToId,_ := strconv.ParseUint( vars["assignedToId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.AssignAssignedToToOnboardingTask(onboardingTaskId, assignedToId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AssignedTo on a OnboardingTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAssignedToFromOnboardingTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.UnassignAssignedToFromOnboardingTask(onboardingTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RelatedOffer on a OnboardingTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRelatedOfferToOnboardingTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedOfferId,_ := strconv.ParseUint( vars["relatedOfferId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.AssignRelatedOfferToOnboardingTask(onboardingTaskId, relatedOfferId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RelatedOffer on a OnboardingTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRelatedOfferFromOnboardingTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.UnassignRelatedOfferFromOnboardingTask(onboardingTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more dependenciesIds as a Dependencies to a OnboardingTask
	//----------------------------------------------------------------------------
func AddDependenciesToOnboardingTask(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dependenciesIds,_ := vars["dependenciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.AddDependenciesToOnboardingTask(onboardingTaskId, dependenciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dependenciesIds as a Dependencies from a OnboardingTask
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDependenciesFromOnboardingTask(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	onboardingTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dependenciesIds,_ := vars["dependenciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the OnboardingTask DAO
	//----------------------------------------------------------------------------
	requestResult := OnboardingTaskDAO.RemoveDependenciesFromOnboardingTask(onboardingTaskId, dependenciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
