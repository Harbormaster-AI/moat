package controller

import (
    BenefitEnrollmentDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BenefitEnrollmentDAO for database creation
//----------------------------------------------------------------------------
func CreateBenefitEnrollment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BenefitEnrollment model
	//----------------------------------------------------------------------------
	data := model.BenefitEnrollment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BenefitEnrollment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment data access object to create
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.CreateBenefitEnrollment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BenefitEnrollmentDAO to find the relevant BenefitEnrollment
//----------------------------------------------------------------------------
func GetBenefitEnrollment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BenefitEnrollment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.GetBenefitEnrollment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BenefitEnrollmentDAO for database read of all BenefitEnrollments
//----------------------------------------------------------------------------
func GetAllBenefitEnrollment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.GetAllBenefitEnrollment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BenefitEnrollmentDAO for database save
//----------------------------------------------------------------------------
func UpdateBenefitEnrollment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BenefitEnrollment model
	//----------------------------------------------------------------------------
	var data = model.BenefitEnrollment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BenefitEnrollment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.UpdateBenefitEnrollment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BenefitEnrollmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBenefitEnrollment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BenefitEnrollment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BenefitEnrollmentDAO.DeleteBenefitEnrollment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a BenefitPlan on a BenefitEnrollment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBenefitPlanToBenefitEnrollment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	benefitEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	benefitPlanId,_ := strconv.ParseUint( vars["benefitPlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.AssignBenefitPlanToBenefitEnrollment(benefitEnrollmentId, benefitPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BenefitPlan on a BenefitEnrollment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBenefitPlanFromBenefitEnrollment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	benefitEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.UnassignBenefitPlanFromBenefitEnrollment(benefitEnrollmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a BenefitEnrollment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToBenefitEnrollment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	benefitEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.AssignEmployeeToBenefitEnrollment(benefitEnrollmentId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a BenefitEnrollment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromBenefitEnrollment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	benefitEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.UnassignEmployeeFromBenefitEnrollment(benefitEnrollmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more dependentsIds as a Dependents to a BenefitEnrollment
	//----------------------------------------------------------------------------
func AddDependentsToBenefitEnrollment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	benefitEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dependentsIds,_ := vars["dependentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.AddDependentsToBenefitEnrollment(benefitEnrollmentId, dependentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dependentsIds as a Dependents from a BenefitEnrollment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDependentsFromBenefitEnrollment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	benefitEnrollmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dependentsIds,_ := vars["dependentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BenefitEnrollment DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitEnrollmentDAO.RemoveDependentsFromBenefitEnrollment(benefitEnrollmentId, dependentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
