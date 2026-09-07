package controller

import (
    InsuranceProductDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InsuranceProductDAO for database creation
//----------------------------------------------------------------------------
func CreateInsuranceProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsuranceProduct model
	//----------------------------------------------------------------------------
	data := model.InsuranceProduct{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsuranceProduct model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct data access object to create
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.CreateInsuranceProduct( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InsuranceProductDAO to find the relevant InsuranceProduct
//----------------------------------------------------------------------------
func GetInsuranceProduct(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsuranceProduct data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.GetInsuranceProduct(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InsuranceProductDAO for database read of all InsuranceProducts
//----------------------------------------------------------------------------
func GetAllInsuranceProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.GetAllInsuranceProduct()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InsuranceProductDAO for database save
//----------------------------------------------------------------------------
func UpdateInsuranceProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsuranceProduct model
	//----------------------------------------------------------------------------
	var data = model.InsuranceProduct{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsuranceProduct model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.UpdateInsuranceProduct(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InsuranceProductDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInsuranceProduct(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsuranceProduct data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InsuranceProductDAO.DeleteInsuranceProduct(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Insurer on a InsuranceProduct
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInsurerToInsuranceProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insuranceProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insurerId,_ := strconv.ParseUint( vars["insurerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct DAO
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.AssignInsurerToInsuranceProduct(insuranceProductId, insurerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Insurer on a InsuranceProduct
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInsurerFromInsuranceProduct( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insuranceProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct DAO
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.UnassignInsurerFromInsuranceProduct(insuranceProductId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more coverageDefinitionsIds as a CoverageDefinitions to a InsuranceProduct
	//----------------------------------------------------------------------------
func AddCoverageDefinitionsToInsuranceProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insuranceProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coverageDefinitionsIds,_ := vars["coverageDefinitionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct DAO
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.AddCoverageDefinitionsToInsuranceProduct(insuranceProductId, coverageDefinitionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more coverageDefinitionsIds as a CoverageDefinitions from a InsuranceProduct
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCoverageDefinitionsFromInsuranceProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insuranceProductId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coverageDefinitionsIds,_ := vars["coverageDefinitionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsuranceProduct DAO
	//----------------------------------------------------------------------------
	requestResult := InsuranceProductDAO.RemoveCoverageDefinitionsFromInsuranceProduct(insuranceProductId, coverageDefinitionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
