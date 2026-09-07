package controller

import (
    PerformanceCycleDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PerformanceCycleDAO for database creation
//----------------------------------------------------------------------------
func CreatePerformanceCycle(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PerformanceCycle model
	//----------------------------------------------------------------------------
	data := model.PerformanceCycle{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PerformanceCycle model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle data access object to create
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.CreatePerformanceCycle( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PerformanceCycleDAO to find the relevant PerformanceCycle
//----------------------------------------------------------------------------
func GetPerformanceCycle(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PerformanceCycle data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.GetPerformanceCycle(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PerformanceCycleDAO for database read of all PerformanceCycles
//----------------------------------------------------------------------------
func GetAllPerformanceCycle(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.GetAllPerformanceCycle()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PerformanceCycleDAO for database save
//----------------------------------------------------------------------------
func UpdatePerformanceCycle(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PerformanceCycle model
	//----------------------------------------------------------------------------
	var data = model.PerformanceCycle{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PerformanceCycle model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.UpdatePerformanceCycle(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PerformanceCycleDAO for database deletion
//----------------------------------------------------------------------------
func DeletePerformanceCycle(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PerformanceCycle data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PerformanceCycleDAO.DeletePerformanceCycle(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a PerformanceCycle
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToPerformanceCycle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceCycleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.AssignOrganizationToPerformanceCycle(performanceCycleId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a PerformanceCycle
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromPerformanceCycle( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceCycleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.UnassignOrganizationFromPerformanceCycle(performanceCycleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more reviewsIds as a Reviews to a PerformanceCycle
	//----------------------------------------------------------------------------
func AddReviewsToPerformanceCycle(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceCycleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reviewsIds,_ := vars["reviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.AddReviewsToPerformanceCycle(performanceCycleId, reviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reviewsIds as a Reviews from a PerformanceCycle
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReviewsFromPerformanceCycle(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceCycleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reviewsIds,_ := vars["reviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.RemoveReviewsFromPerformanceCycle(performanceCycleId, reviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more goalsIds as a Goals to a PerformanceCycle
	//----------------------------------------------------------------------------
func AddGoalsToPerformanceCycle(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceCycleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goalsIds,_ := vars["goalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.AddGoalsToPerformanceCycle(performanceCycleId, goalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more goalsIds as a Goals from a PerformanceCycle
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGoalsFromPerformanceCycle(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceCycleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goalsIds,_ := vars["goalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceCycle DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceCycleDAO.RemoveGoalsFromPerformanceCycle(performanceCycleId, goalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
