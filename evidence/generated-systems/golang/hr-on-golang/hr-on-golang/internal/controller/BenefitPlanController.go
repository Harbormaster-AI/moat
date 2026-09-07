package controller

import (
    BenefitPlanDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BenefitPlanDAO for database creation
//----------------------------------------------------------------------------
func CreateBenefitPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BenefitPlan model
	//----------------------------------------------------------------------------
	data := model.BenefitPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BenefitPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.CreateBenefitPlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BenefitPlanDAO to find the relevant BenefitPlan
//----------------------------------------------------------------------------
func GetBenefitPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BenefitPlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.GetBenefitPlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BenefitPlanDAO for database read of all BenefitPlans
//----------------------------------------------------------------------------
func GetAllBenefitPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.GetAllBenefitPlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BenefitPlanDAO for database save
//----------------------------------------------------------------------------
func UpdateBenefitPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BenefitPlan model
	//----------------------------------------------------------------------------
	var data = model.BenefitPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BenefitPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.UpdateBenefitPlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BenefitPlanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBenefitPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BenefitPlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BenefitPlanDAO.DeleteBenefitPlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a BenefitPlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToBenefitPlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	benefitPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.AssignOrganizationToBenefitPlan(benefitPlanId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a BenefitPlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromBenefitPlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	benefitPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.UnassignOrganizationFromBenefitPlan(benefitPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more enrollmentsIds as a Enrollments to a BenefitPlan
	//----------------------------------------------------------------------------
func AddEnrollmentsToBenefitPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	benefitPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enrollmentsIds,_ := vars["enrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.AddEnrollmentsToBenefitPlan(benefitPlanId, enrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more enrollmentsIds as a Enrollments from a BenefitPlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEnrollmentsFromBenefitPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	benefitPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enrollmentsIds,_ := vars["enrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BenefitPlan DAO
	//----------------------------------------------------------------------------
	requestResult := BenefitPlanDAO.RemoveEnrollmentsFromBenefitPlan(benefitPlanId, enrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
