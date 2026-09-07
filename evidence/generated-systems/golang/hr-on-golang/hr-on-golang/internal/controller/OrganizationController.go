package controller

import (
    OrganizationDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
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
	// adds one or more departmentsIds as a Departments to a Organization
	//----------------------------------------------------------------------------
func AddDepartmentsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddDepartmentsToOrganization(organizationId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more departmentsIds as a Departments from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDepartmentsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveDepartmentsFromOrganization(organizationId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more locationsIds as a Locations to a Organization
	//----------------------------------------------------------------------------
func AddLocationsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationsIds,_ := vars["locationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddLocationsToOrganization(organizationId, locationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more locationsIds as a Locations from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLocationsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationsIds,_ := vars["locationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveLocationsFromOrganization(organizationId, locationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more jobFamiliesIds as a JobFamilies to a Organization
	//----------------------------------------------------------------------------
func AddJobFamiliesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobFamiliesIds,_ := vars["jobFamiliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddJobFamiliesToOrganization(organizationId, jobFamiliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more jobFamiliesIds as a JobFamilies from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveJobFamiliesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobFamiliesIds,_ := vars["jobFamiliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveJobFamiliesFromOrganization(organizationId, jobFamiliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more benefitPlansIds as a BenefitPlans to a Organization
	//----------------------------------------------------------------------------
func AddBenefitPlansToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	benefitPlansIds,_ := vars["benefitPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddBenefitPlansToOrganization(organizationId, benefitPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more benefitPlansIds as a BenefitPlans from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBenefitPlansFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	benefitPlansIds,_ := vars["benefitPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveBenefitPlansFromOrganization(organizationId, benefitPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more costCentersIds as a CostCenters to a Organization
	//----------------------------------------------------------------------------
func AddCostCentersToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	costCentersIds,_ := vars["costCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddCostCentersToOrganization(organizationId, costCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more costCentersIds as a CostCenters from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCostCentersFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	costCentersIds,_ := vars["costCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveCostCentersFromOrganization(organizationId, costCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more payrollCalendarsIds as a PayrollCalendars to a Organization
	//----------------------------------------------------------------------------
func AddPayrollCalendarsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollCalendarsIds,_ := vars["payrollCalendarsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddPayrollCalendarsToOrganization(organizationId, payrollCalendarsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more payrollCalendarsIds as a PayrollCalendars from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePayrollCalendarsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollCalendarsIds,_ := vars["payrollCalendarsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemovePayrollCalendarsFromOrganization(organizationId, payrollCalendarsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
