package controller

import (
    InsurerDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InsurerDAO for database creation
//----------------------------------------------------------------------------
func CreateInsurer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Insurer model
	//----------------------------------------------------------------------------
	data := model.Insurer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Insurer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Insurer data access object to create
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.CreateInsurer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InsurerDAO to find the relevant Insurer
//----------------------------------------------------------------------------
func GetInsurer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Insurer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.GetInsurer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InsurerDAO for database read of all Insurers
//----------------------------------------------------------------------------
func GetAllInsurer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Insurer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.GetAllInsurer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InsurerDAO for database save
//----------------------------------------------------------------------------
func UpdateInsurer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Insurer model
	//----------------------------------------------------------------------------
	var data = model.Insurer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Insurer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Insurer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.UpdateInsurer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InsurerDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInsurer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Insurer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InsurerDAO.DeleteInsurer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more productsIds as a Products to a Insurer
	//----------------------------------------------------------------------------
func AddProductsToInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productsIds,_ := vars["productsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.AddProductsToInsurer(insurerId, productsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productsIds as a Products from a Insurer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductsFromInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productsIds,_ := vars["productsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.RemoveProductsFromInsurer(insurerId, productsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more distributionPartnersIds as a DistributionPartners to a Insurer
	//----------------------------------------------------------------------------
func AddDistributionPartnersToInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	distributionPartnersIds,_ := vars["distributionPartnersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.AddDistributionPartnersToInsurer(insurerId, distributionPartnersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more distributionPartnersIds as a DistributionPartners from a Insurer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDistributionPartnersFromInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	distributionPartnersIds,_ := vars["distributionPartnersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.RemoveDistributionPartnersFromInsurer(insurerId, distributionPartnersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a Insurer
	//----------------------------------------------------------------------------
func AddPoliciesToInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.AddPoliciesToInsurer(insurerId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a Insurer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.RemovePoliciesFromInsurer(insurerId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a Insurer
	//----------------------------------------------------------------------------
func AddClaimsToInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.AddClaimsToInsurer(insurerId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a Insurer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.RemoveClaimsFromInsurer(insurerId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reinsuranceAgreementsIds as a ReinsuranceAgreements to a Insurer
	//----------------------------------------------------------------------------
func AddReinsuranceAgreementsToInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reinsuranceAgreementsIds,_ := vars["reinsuranceAgreementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.AddReinsuranceAgreementsToInsurer(insurerId, reinsuranceAgreementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reinsuranceAgreementsIds as a ReinsuranceAgreements from a Insurer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReinsuranceAgreementsFromInsurer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insurerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reinsuranceAgreementsIds,_ := vars["reinsuranceAgreementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Insurer DAO
	//----------------------------------------------------------------------------
	requestResult := InsurerDAO.RemoveReinsuranceAgreementsFromInsurer(insurerId, reinsuranceAgreementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
