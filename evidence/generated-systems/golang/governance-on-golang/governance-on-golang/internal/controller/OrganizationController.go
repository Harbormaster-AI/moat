package controller

import (
    OrganizationDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
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
	// adds one or more governanceBodiesIds as a GovernanceBodies to a Organization
	//----------------------------------------------------------------------------
func AddGovernanceBodiesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	governanceBodiesIds,_ := vars["governanceBodiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddGovernanceBodiesToOrganization(organizationId, governanceBodiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more governanceBodiesIds as a GovernanceBodies from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGovernanceBodiesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	governanceBodiesIds,_ := vars["governanceBodiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveGovernanceBodiesFromOrganization(organizationId, governanceBodiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a Organization
	//----------------------------------------------------------------------------
func AddPoliciesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddPoliciesToOrganization(organizationId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemovePoliciesFromOrganization(organizationId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more risksIds as a Risks to a Organization
	//----------------------------------------------------------------------------
func AddRisksToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	risksIds,_ := vars["risksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddRisksToOrganization(organizationId, risksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more risksIds as a Risks from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRisksFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	risksIds,_ := vars["risksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveRisksFromOrganization(organizationId, risksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more thirdPartiesIds as a ThirdParties to a Organization
	//----------------------------------------------------------------------------
func AddThirdPartiesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartiesIds,_ := vars["thirdPartiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddThirdPartiesToOrganization(organizationId, thirdPartiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more thirdPartiesIds as a ThirdParties from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveThirdPartiesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartiesIds,_ := vars["thirdPartiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveThirdPartiesFromOrganization(organizationId, thirdPartiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsRepositoriesIds as a RecordsRepositories to a Organization
	//----------------------------------------------------------------------------
func AddRecordsRepositoriesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsRepositoriesIds,_ := vars["recordsRepositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddRecordsRepositoriesToOrganization(organizationId, recordsRepositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsRepositoriesIds as a RecordsRepositories from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsRepositoriesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsRepositoriesIds,_ := vars["recordsRepositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveRecordsRepositoriesFromOrganization(organizationId, recordsRepositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataProcessingActivitiesIds as a DataProcessingActivities to a Organization
	//----------------------------------------------------------------------------
func AddDataProcessingActivitiesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataProcessingActivitiesIds,_ := vars["dataProcessingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddDataProcessingActivitiesToOrganization(organizationId, dataProcessingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataProcessingActivitiesIds as a DataProcessingActivities from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataProcessingActivitiesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataProcessingActivitiesIds,_ := vars["dataProcessingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveDataProcessingActivitiesFromOrganization(organizationId, dataProcessingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more complianceProgramsIds as a CompliancePrograms to a Organization
	//----------------------------------------------------------------------------
func AddComplianceProgramsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	complianceProgramsIds,_ := vars["complianceProgramsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddComplianceProgramsToOrganization(organizationId, complianceProgramsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more complianceProgramsIds as a CompliancePrograms from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveComplianceProgramsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	complianceProgramsIds,_ := vars["complianceProgramsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveComplianceProgramsFromOrganization(organizationId, complianceProgramsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more auditProgramsIds as a AuditPrograms to a Organization
	//----------------------------------------------------------------------------
func AddAuditProgramsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	auditProgramsIds,_ := vars["auditProgramsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddAuditProgramsToOrganization(organizationId, auditProgramsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more auditProgramsIds as a AuditPrograms from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAuditProgramsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	auditProgramsIds,_ := vars["auditProgramsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveAuditProgramsFromOrganization(organizationId, auditProgramsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more businessUnitsIds as a BusinessUnits to a Organization
	//----------------------------------------------------------------------------
func AddBusinessUnitsToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitsIds,_ := vars["businessUnitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddBusinessUnitsToOrganization(organizationId, businessUnitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more businessUnitsIds as a BusinessUnits from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBusinessUnitsFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitsIds,_ := vars["businessUnitsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveBusinessUnitsFromOrganization(organizationId, businessUnitsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more mattersIds as a Matters to a Organization
	//----------------------------------------------------------------------------
func AddMattersToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mattersIds,_ := vars["mattersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddMattersToOrganization(organizationId, mattersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more mattersIds as a Matters from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMattersFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mattersIds,_ := vars["mattersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveMattersFromOrganization(organizationId, mattersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataBreachesIds as a DataBreaches to a Organization
	//----------------------------------------------------------------------------
func AddDataBreachesToOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.AddDataBreachesToOrganization(organizationId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataBreachesIds as a DataBreaches from a Organization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataBreachesFromOrganization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	organizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Organization DAO
	//----------------------------------------------------------------------------
	requestResult := OrganizationDAO.RemoveDataBreachesFromOrganization(organizationId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
