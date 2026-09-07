package controller

import (
    PricingPlanDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PricingPlanDAO for database creation
//----------------------------------------------------------------------------
func CreatePricingPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PricingPlan model
	//----------------------------------------------------------------------------
	data := model.PricingPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PricingPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.CreatePricingPlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PricingPlanDAO to find the relevant PricingPlan
//----------------------------------------------------------------------------
func GetPricingPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PricingPlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.GetPricingPlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PricingPlanDAO for database read of all PricingPlans
//----------------------------------------------------------------------------
func GetAllPricingPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.GetAllPricingPlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PricingPlanDAO for database save
//----------------------------------------------------------------------------
func UpdatePricingPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PricingPlan model
	//----------------------------------------------------------------------------
	var data = model.PricingPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PricingPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.UpdatePricingPlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PricingPlanDAO for database deletion
//----------------------------------------------------------------------------
func DeletePricingPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PricingPlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PricingPlanDAO.DeletePricingPlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ProductOffering on a PricingPlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductOfferingToPricingPlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	pricingPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productOfferingId,_ := strconv.ParseUint( vars["productOfferingId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan DAO
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.AssignProductOfferingToPricingPlan(pricingPlanId, productOfferingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ProductOffering on a PricingPlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductOfferingFromPricingPlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	pricingPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan DAO
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.UnassignProductOfferingFromPricingPlan(pricingPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more feeSchedulesIds as a FeeSchedules to a PricingPlan
	//----------------------------------------------------------------------------
func AddFeeSchedulesToPricingPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pricingPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feeSchedulesIds,_ := vars["feeSchedulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan DAO
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.AddFeeSchedulesToPricingPlan(pricingPlanId, feeSchedulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more feeSchedulesIds as a FeeSchedules from a PricingPlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeeSchedulesFromPricingPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pricingPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	feeSchedulesIds,_ := vars["feeSchedulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan DAO
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.RemoveFeeSchedulesFromPricingPlan(pricingPlanId, feeSchedulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more limitsIds as a Limits to a PricingPlan
	//----------------------------------------------------------------------------
func AddLimitsToPricingPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pricingPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	limitsIds,_ := vars["limitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan DAO
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.AddLimitsToPricingPlan(pricingPlanId, limitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more limitsIds as a Limits from a PricingPlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLimitsFromPricingPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pricingPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	limitsIds,_ := vars["limitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PricingPlan DAO
	//----------------------------------------------------------------------------
	requestResult := PricingPlanDAO.RemoveLimitsFromPricingPlan(pricingPlanId, limitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
