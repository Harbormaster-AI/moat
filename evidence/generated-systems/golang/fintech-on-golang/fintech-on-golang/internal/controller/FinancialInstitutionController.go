package controller

import (
    FinancialInstitutionDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FinancialInstitutionDAO for database creation
//----------------------------------------------------------------------------
func CreateFinancialInstitution(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FinancialInstitution model
	//----------------------------------------------------------------------------
	data := model.FinancialInstitution{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FinancialInstitution model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution data access object to create
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.CreateFinancialInstitution( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FinancialInstitutionDAO to find the relevant FinancialInstitution
//----------------------------------------------------------------------------
func GetFinancialInstitution(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FinancialInstitution data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.GetFinancialInstitution(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FinancialInstitutionDAO for database read of all FinancialInstitutions
//----------------------------------------------------------------------------
func GetAllFinancialInstitution(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.GetAllFinancialInstitution()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FinancialInstitutionDAO for database save
//----------------------------------------------------------------------------
func UpdateFinancialInstitution(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FinancialInstitution model
	//----------------------------------------------------------------------------
	var data = model.FinancialInstitution{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FinancialInstitution model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.UpdateFinancialInstitution(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FinancialInstitutionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFinancialInstitution(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FinancialInstitution data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FinancialInstitutionDAO.DeleteFinancialInstitution(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more branchesIds as a Branches to a FinancialInstitution
	//----------------------------------------------------------------------------
func AddBranchesToFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	branchesIds,_ := vars["branchesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.AddBranchesToFinancialInstitution(financialInstitutionId, branchesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more branchesIds as a Branches from a FinancialInstitution
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBranchesFromFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	branchesIds,_ := vars["branchesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.RemoveBranchesFromFinancialInstitution(financialInstitutionId, branchesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more customersIds as a Customers to a FinancialInstitution
	//----------------------------------------------------------------------------
func AddCustomersToFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.AddCustomersToFinancialInstitution(financialInstitutionId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more customersIds as a Customers from a FinancialInstitution
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCustomersFromFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customersIds,_ := vars["customersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.RemoveCustomersFromFinancialInstitution(financialInstitutionId, customersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more productOfferingsIds as a ProductOfferings to a FinancialInstitution
	//----------------------------------------------------------------------------
func AddProductOfferingsToFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productOfferingsIds,_ := vars["productOfferingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.AddProductOfferingsToFinancialInstitution(financialInstitutionId, productOfferingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productOfferingsIds as a ProductOfferings from a FinancialInstitution
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductOfferingsFromFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productOfferingsIds,_ := vars["productOfferingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.RemoveProductOfferingsFromFinancialInstitution(financialInstitutionId, productOfferingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more paymentProcessorsIds as a PaymentProcessors to a FinancialInstitution
	//----------------------------------------------------------------------------
func AddPaymentProcessorsToFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentProcessorsIds,_ := vars["paymentProcessorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.AddPaymentProcessorsToFinancialInstitution(financialInstitutionId, paymentProcessorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentProcessorsIds as a PaymentProcessors from a FinancialInstitution
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentProcessorsFromFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentProcessorsIds,_ := vars["paymentProcessorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.RemovePaymentProcessorsFromFinancialInstitution(financialInstitutionId, paymentProcessorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more compliancePoliciesIds as a CompliancePolicies to a FinancialInstitution
	//----------------------------------------------------------------------------
func AddCompliancePoliciesToFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compliancePoliciesIds,_ := vars["compliancePoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.AddCompliancePoliciesToFinancialInstitution(financialInstitutionId, compliancePoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more compliancePoliciesIds as a CompliancePolicies from a FinancialInstitution
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCompliancePoliciesFromFinancialInstitution(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	financialInstitutionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compliancePoliciesIds,_ := vars["compliancePoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FinancialInstitution DAO
	//----------------------------------------------------------------------------
	requestResult := FinancialInstitutionDAO.RemoveCompliancePoliciesFromFinancialInstitution(financialInstitutionId, compliancePoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
