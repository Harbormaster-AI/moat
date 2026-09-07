package controller

import (
    UsageLimitDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UsageLimitDAO for database creation
//----------------------------------------------------------------------------
func CreateUsageLimit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UsageLimit model
	//----------------------------------------------------------------------------
	data := model.UsageLimit{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UsageLimit model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UsageLimit data access object to create
	//----------------------------------------------------------------------------
	requestResult := UsageLimitDAO.CreateUsageLimit( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UsageLimitDAO to find the relevant UsageLimit
//----------------------------------------------------------------------------
func GetUsageLimit(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UsageLimit data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UsageLimitDAO.GetUsageLimit(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UsageLimitDAO for database read of all UsageLimits
//----------------------------------------------------------------------------
func GetAllUsageLimit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the UsageLimit data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UsageLimitDAO.GetAllUsageLimit()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UsageLimitDAO for database save
//----------------------------------------------------------------------------
func UpdateUsageLimit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UsageLimit model
	//----------------------------------------------------------------------------
	var data = model.UsageLimit{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UsageLimit model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UsageLimit data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UsageLimitDAO.UpdateUsageLimit(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UsageLimitDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUsageLimit(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UsageLimit data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UsageLimitDAO.DeleteUsageLimit(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PricingPlan on a UsageLimit
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPricingPlanToUsageLimit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageLimitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pricingPlanId,_ := strconv.ParseUint( vars["pricingPlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageLimit DAO
	//----------------------------------------------------------------------------
	requestResult := UsageLimitDAO.AssignPricingPlanToUsageLimit(usageLimitId, pricingPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PricingPlan on a UsageLimit
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPricingPlanFromUsageLimit( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageLimitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageLimit DAO
	//----------------------------------------------------------------------------
	requestResult := UsageLimitDAO.UnassignPricingPlanFromUsageLimit(usageLimitId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


