package controller

import (
    CompetencyRatingDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CompetencyRatingDAO for database creation
//----------------------------------------------------------------------------
func CreateCompetencyRating(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CompetencyRating model
	//----------------------------------------------------------------------------
	data := model.CompetencyRating{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CompetencyRating model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating data access object to create
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.CreateCompetencyRating( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CompetencyRatingDAO to find the relevant CompetencyRating
//----------------------------------------------------------------------------
func GetCompetencyRating(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CompetencyRating data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.GetCompetencyRating(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CompetencyRatingDAO for database read of all CompetencyRatings
//----------------------------------------------------------------------------
func GetAllCompetencyRating(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.GetAllCompetencyRating()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CompetencyRatingDAO for database save
//----------------------------------------------------------------------------
func UpdateCompetencyRating(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CompetencyRating model
	//----------------------------------------------------------------------------
	var data = model.CompetencyRating{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CompetencyRating model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.UpdateCompetencyRating(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CompetencyRatingDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCompetencyRating(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CompetencyRating data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CompetencyRatingDAO.DeleteCompetencyRating(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Review on a CompetencyRating
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignReviewToCompetencyRating(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	competencyRatingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reviewId,_ := strconv.ParseUint( vars["reviewId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.AssignReviewToCompetencyRating(competencyRatingId, reviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Review on a CompetencyRating
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignReviewFromCompetencyRating( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	competencyRatingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.UnassignReviewFromCompetencyRating(competencyRatingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Competency on a CompetencyRating
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCompetencyToCompetencyRating(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	competencyRatingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competencyId,_ := strconv.ParseUint( vars["competencyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.AssignCompetencyToCompetencyRating(competencyRatingId, competencyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Competency on a CompetencyRating
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCompetencyFromCompetencyRating( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	competencyRatingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CompetencyRating DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyRatingDAO.UnassignCompetencyFromCompetencyRating(competencyRatingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


