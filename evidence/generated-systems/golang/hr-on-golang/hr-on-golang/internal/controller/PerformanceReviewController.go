package controller

import (
    PerformanceReviewDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PerformanceReviewDAO for database creation
//----------------------------------------------------------------------------
func CreatePerformanceReview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PerformanceReview model
	//----------------------------------------------------------------------------
	data := model.PerformanceReview{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PerformanceReview model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview data access object to create
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.CreatePerformanceReview( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PerformanceReviewDAO to find the relevant PerformanceReview
//----------------------------------------------------------------------------
func GetPerformanceReview(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PerformanceReview data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.GetPerformanceReview(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PerformanceReviewDAO for database read of all PerformanceReviews
//----------------------------------------------------------------------------
func GetAllPerformanceReview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.GetAllPerformanceReview()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PerformanceReviewDAO for database save
//----------------------------------------------------------------------------
func UpdatePerformanceReview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PerformanceReview model
	//----------------------------------------------------------------------------
	var data = model.PerformanceReview{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PerformanceReview model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.UpdatePerformanceReview(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PerformanceReviewDAO for database deletion
//----------------------------------------------------------------------------
func DeletePerformanceReview(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PerformanceReview data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PerformanceReviewDAO.DeletePerformanceReview(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a PerformanceReview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToPerformanceReview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.AssignEmployeeToPerformanceReview(performanceReviewId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a PerformanceReview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromPerformanceReview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.UnassignEmployeeFromPerformanceReview(performanceReviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Reviewer on a PerformanceReview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignReviewerToPerformanceReview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reviewerId,_ := strconv.ParseUint( vars["reviewerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.AssignReviewerToPerformanceReview(performanceReviewId, reviewerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Reviewer on a PerformanceReview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignReviewerFromPerformanceReview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.UnassignReviewerFromPerformanceReview(performanceReviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Cycle on a PerformanceReview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCycleToPerformanceReview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cycleId,_ := strconv.ParseUint( vars["cycleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.AssignCycleToPerformanceReview(performanceReviewId, cycleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Cycle on a PerformanceReview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCycleFromPerformanceReview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.UnassignCycleFromPerformanceReview(performanceReviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more competencyRatingsIds as a CompetencyRatings to a PerformanceReview
	//----------------------------------------------------------------------------
func AddCompetencyRatingsToPerformanceReview(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competencyRatingsIds,_ := vars["competencyRatingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.AddCompetencyRatingsToPerformanceReview(performanceReviewId, competencyRatingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more competencyRatingsIds as a CompetencyRatings from a PerformanceReview
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCompetencyRatingsFromPerformanceReview(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competencyRatingsIds,_ := vars["competencyRatingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.RemoveCompetencyRatingsFromPerformanceReview(performanceReviewId, competencyRatingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more goalsIds as a Goals to a PerformanceReview
	//----------------------------------------------------------------------------
func AddGoalsToPerformanceReview(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goalsIds,_ := vars["goalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.AddGoalsToPerformanceReview(performanceReviewId, goalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more goalsIds as a Goals from a PerformanceReview
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGoalsFromPerformanceReview(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	performanceReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goalsIds,_ := vars["goalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PerformanceReview DAO
	//----------------------------------------------------------------------------
	requestResult := PerformanceReviewDAO.RemoveGoalsFromPerformanceReview(performanceReviewId, goalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
