package controller

import (
    OrganizationDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OrganizationDAO for database creation
//----------------------------------------------------------------------------
func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Organization model
	//----------------------------------------------------------------------------
	data := model.Organization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Organization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Organization data access object to create
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.CreateOrganization( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OrganizationDAO to find the relevant Organization
//----------------------------------------------------------------------------
func GetOrganization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Organization data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.GetOrganization(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OrganizationDAO for database read of all Organizations
//----------------------------------------------------------------------------
func GetAllOrganization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Organization data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.GetAllOrganization()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OrganizationDAO for database save
//----------------------------------------------------------------------------
func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Organization model
	//----------------------------------------------------------------------------
	var data = model.Organization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Organization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Organization data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.UpdateOrganization(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OrganizationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Organization data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OrganizationDAO.DeleteOrganization(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more usersIds as a Users to a Organization
	//----------------------------------------------------------------------------
func AddUsersToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddUsersToOrganization(organizationId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usersIds as a Users from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsersFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveUsersFromOrganization(organizationId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Organization
	//----------------------------------------------------------------------------
func AddAccountsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddAccountsToOrganization(organizationId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveAccountsFromOrganization(organizationId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more teamsIds as a Teams to a Organization
	//----------------------------------------------------------------------------
func AddTeamsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddTeamsToOrganization(organizationId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more teamsIds as a Teams from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTeamsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveTeamsFromOrganization(organizationId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more territoriesIds as a Territories to a Organization
	//----------------------------------------------------------------------------
func AddTerritoriesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	territoriesIds,_ := vars["territoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddTerritoriesToOrganization(organizationId, territoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more territoriesIds as a Territories from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTerritoriesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	territoriesIds,_ := vars["territoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveTerritoriesFromOrganization(organizationId, territoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more productsIds as a Products to a Organization
	//----------------------------------------------------------------------------
func AddProductsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productsIds,_ := vars["productsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddProductsToOrganization(organizationId, productsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productsIds as a Products from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productsIds,_ := vars["productsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveProductsFromOrganization(organizationId, productsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more priceBooksIds as a PriceBooks to a Organization
	//----------------------------------------------------------------------------
func AddPriceBooksToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBooksIds,_ := vars["priceBooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddPriceBooksToOrganization(organizationId, priceBooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more priceBooksIds as a PriceBooks from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePriceBooksFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBooksIds,_ := vars["priceBooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemovePriceBooksFromOrganization(organizationId, priceBooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a Organization
	//----------------------------------------------------------------------------
func AddCampaignsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddCampaignsToOrganization(organizationId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveCampaignsFromOrganization(organizationId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
