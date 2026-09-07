package controller

import (
    TrainingCourseDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TrainingCourseDAO for database creation
//----------------------------------------------------------------------------
func CreateTrainingCourse(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrainingCourse model
	//----------------------------------------------------------------------------
	data := model.TrainingCourse{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrainingCourse model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse data access object to create
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.CreateTrainingCourse( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TrainingCourseDAO to find the relevant TrainingCourse
//----------------------------------------------------------------------------
func GetTrainingCourse(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrainingCourse data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.GetTrainingCourse(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TrainingCourseDAO for database read of all TrainingCourses
//----------------------------------------------------------------------------
func GetAllTrainingCourse(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.GetAllTrainingCourse()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TrainingCourseDAO for database save
//----------------------------------------------------------------------------
func UpdateTrainingCourse(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrainingCourse model
	//----------------------------------------------------------------------------
	var data = model.TrainingCourse{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrainingCourse model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.UpdateTrainingCourse(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TrainingCourseDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTrainingCourse(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrainingCourse data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TrainingCourseDAO.DeleteTrainingCourse(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more prerequisitesIds as a Prerequisites to a TrainingCourse
	//----------------------------------------------------------------------------
func AddPrerequisitesToTrainingCourse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingCourseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	prerequisitesIds,_ := vars["prerequisitesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.AddPrerequisitesToTrainingCourse(trainingCourseId, prerequisitesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more prerequisitesIds as a Prerequisites from a TrainingCourse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePrerequisitesFromTrainingCourse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingCourseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	prerequisitesIds,_ := vars["prerequisitesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.RemovePrerequisitesFromTrainingCourse(trainingCourseId, prerequisitesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more enrollmentsIds as a Enrollments to a TrainingCourse
	//----------------------------------------------------------------------------
func AddEnrollmentsToTrainingCourse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingCourseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enrollmentsIds,_ := vars["enrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.AddEnrollmentsToTrainingCourse(trainingCourseId, enrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more enrollmentsIds as a Enrollments from a TrainingCourse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEnrollmentsFromTrainingCourse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingCourseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enrollmentsIds,_ := vars["enrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.RemoveEnrollmentsFromTrainingCourse(trainingCourseId, enrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more jobProfilesIds as a JobProfiles to a TrainingCourse
	//----------------------------------------------------------------------------
func AddJobProfilesToTrainingCourse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingCourseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfilesIds,_ := vars["jobProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.AddJobProfilesToTrainingCourse(trainingCourseId, jobProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more jobProfilesIds as a JobProfiles from a TrainingCourse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveJobProfilesFromTrainingCourse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingCourseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfilesIds,_ := vars["jobProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingCourse DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingCourseDAO.RemoveJobProfilesFromTrainingCourse(trainingCourseId, jobProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
