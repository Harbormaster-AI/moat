package controller

import (
    PolicyDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PolicyDAO for database creation
//----------------------------------------------------------------------------
func CreatePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Policy model
	//----------------------------------------------------------------------------
	data := model.Policy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Policy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Policy data access object to create
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.CreatePolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PolicyDAO to find the relevant Policy
//----------------------------------------------------------------------------
func GetPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Policy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.GetPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PolicyDAO for database read of all Policys
//----------------------------------------------------------------------------
func GetAllPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Policy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.GetAllPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PolicyDAO for database save
//----------------------------------------------------------------------------
func UpdatePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Policy model
	//----------------------------------------------------------------------------
	var data = model.Policy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Policy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Policy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UpdatePolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeletePolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Policy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PolicyDAO.DeletePolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Insurer on a Policy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInsurerToPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insurerId,_ := strconv.ParseUint( vars["insurerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AssignInsurerToPolicy(policyId, insurerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Insurer on a Policy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInsurerFromPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UnassignInsurerFromPolicy(policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Policy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AssignCustomerToPolicy(policyId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Policy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UnassignCustomerFromPolicy(policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a Policy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AssignProductToPolicy(policyId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a Policy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UnassignProductFromPolicy(policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Agent on a Policy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAgentToPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agentId,_ := strconv.ParseUint( vars["agentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AssignAgentToPolicy(policyId, agentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Agent on a Policy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAgentFromPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UnassignAgentFromPolicy(policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a BillingAccount on a Policy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBillingAccountToPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	billingAccountId,_ := strconv.ParseUint( vars["billingAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AssignBillingAccountToPolicy(policyId, billingAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BillingAccount on a Policy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBillingAccountFromPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.UnassignBillingAccountFromPolicy(policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more coveragesIds as a Coverages to a Policy
	//----------------------------------------------------------------------------
func AddCoveragesToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddCoveragesToPolicy(policyId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more coveragesIds as a Coverages from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCoveragesFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveCoveragesFromPolicy(policyId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more insuredObjectsIds as a InsuredObjects to a Policy
	//----------------------------------------------------------------------------
func AddInsuredObjectsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectsIds,_ := vars["insuredObjectsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddInsuredObjectsToPolicy(policyId, insuredObjectsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more insuredObjectsIds as a InsuredObjects from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInsuredObjectsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insuredObjectsIds,_ := vars["insuredObjectsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveInsuredObjectsFromPolicy(policyId, insuredObjectsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more endorsementsIds as a Endorsements to a Policy
	//----------------------------------------------------------------------------
func AddEndorsementsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	endorsementsIds,_ := vars["endorsementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddEndorsementsToPolicy(policyId, endorsementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more endorsementsIds as a Endorsements from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEndorsementsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	endorsementsIds,_ := vars["endorsementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveEndorsementsFromPolicy(policyId, endorsementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more beneficiariesIds as a Beneficiaries to a Policy
	//----------------------------------------------------------------------------
func AddBeneficiariesToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	beneficiariesIds,_ := vars["beneficiariesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddBeneficiariesToPolicy(policyId, beneficiariesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more beneficiariesIds as a Beneficiaries from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBeneficiariesFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	beneficiariesIds,_ := vars["beneficiariesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveBeneficiariesFromPolicy(policyId, beneficiariesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a Policy
	//----------------------------------------------------------------------------
func AddClaimsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddClaimsToPolicy(policyId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveClaimsFromPolicy(policyId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reinsuranceAgreementsIds as a ReinsuranceAgreements to a Policy
	//----------------------------------------------------------------------------
func AddReinsuranceAgreementsToPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reinsuranceAgreementsIds,_ := vars["reinsuranceAgreementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.AddReinsuranceAgreementsToPolicy(policyId, reinsuranceAgreementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reinsuranceAgreementsIds as a ReinsuranceAgreements from a Policy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReinsuranceAgreementsFromPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	policyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reinsuranceAgreementsIds,_ := vars["reinsuranceAgreementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Policy DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyDAO.RemoveReinsuranceAgreementsFromPolicy(policyId, reinsuranceAgreementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
