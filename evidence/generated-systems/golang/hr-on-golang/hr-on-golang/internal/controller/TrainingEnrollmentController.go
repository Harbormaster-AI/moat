package controller

import (
    TrainingEnrollmentDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TrainingEnrollmentDAO for database creation
//----------------------------------------------------------------------------
func CreateTrainingEnrollment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrainingEnrollment model
	//----------------------------------------------------------------------------
	data := model.TrainingEnrollment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrainingEnrollment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment data access object to create
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.CreateTrainingEnrollment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TrainingEnrollmentDAO to find the relevant TrainingEnrollment
//----------------------------------------------------------------------------
func GetTrainingEnrollment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrainingEnrollment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.GetTrainingEnrollment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TrainingEnrollmentDAO for database read of all TrainingEnrollments
//----------------------------------------------------------------------------
func GetAllTrainingEnrollment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.GetAllTrainingEnrollment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TrainingEnrollmentDAO for database save
//----------------------------------------------------------------------------
func UpdateTrainingEnrollment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrainingEnrollment model
	//----------------------------------------------------------------------------
	var data = model.TrainingEnrollment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrainingEnrollment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.UpdateTrainingEnrollment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TrainingEnrollmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTrainingEnrollment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrainingEnrollment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TrainingEnrollmentDAO.DeleteTrainingEnrollment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Course on a TrainingEnrollment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCourseToTrainingEnrollment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	courseId,_ := strconv.ParseUint( vars["courseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.AssignCourseToTrainingEnrollment(trainingEnrollmentId, courseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Course on a TrainingEnrollment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCourseFromTrainingEnrollment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.UnassignCourseFromTrainingEnrollment(trainingEnrollmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a TrainingEnrollment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToTrainingEnrollment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.AssignEmployeeToTrainingEnrollment(trainingEnrollmentId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a TrainingEnrollment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromTrainingEnrollment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.UnassignEmployeeFromTrainingEnrollment(trainingEnrollmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Instructor on a TrainingEnrollment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInstructorToTrainingEnrollment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	instructorId,_ := strconv.ParseUint( vars["instructorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.AssignInstructorToTrainingEnrollment(trainingEnrollmentId, instructorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Instructor on a TrainingEnrollment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInstructorFromTrainingEnrollment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingEnrollmentDAO.UnassignInstructorFromTrainingEnrollment(trainingEnrollmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


