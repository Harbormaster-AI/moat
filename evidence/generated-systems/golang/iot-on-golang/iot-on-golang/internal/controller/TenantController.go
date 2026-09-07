package controller

import (
    TenantDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TenantDAO for database creation
//----------------------------------------------------------------------------
func CreateTenant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Tenant model
	//----------------------------------------------------------------------------
	data := model.Tenant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Tenant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Tenant data access object to create
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.CreateTenant( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TenantDAO to find the relevant Tenant
//----------------------------------------------------------------------------
func GetTenant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Tenant data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.GetTenant(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TenantDAO for database read of all Tenants
//----------------------------------------------------------------------------
func GetAllTenant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Tenant data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.GetAllTenant()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TenantDAO for database save
//----------------------------------------------------------------------------
func UpdateTenant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Tenant model
	//----------------------------------------------------------------------------
	var data = model.Tenant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Tenant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Tenant data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.UpdateTenant(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TenantDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTenant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Tenant data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TenantDAO.DeleteTenant(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more sitesIds as a Sites to a Tenant
	//----------------------------------------------------------------------------
func AddSitesToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sitesIds,_ := vars["sitesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddSitesToTenant(tenantId, sitesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more sitesIds as a Sites from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSitesFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sitesIds,_ := vars["sitesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveSitesFromTenant(tenantId, sitesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more usersIds as a Users to a Tenant
	//----------------------------------------------------------------------------
func AddUsersToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddUsersToTenant(tenantId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usersIds as a Users from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsersFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveUsersFromTenant(tenantId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more devicesIds as a Devices to a Tenant
	//----------------------------------------------------------------------------
func AddDevicesToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	devicesIds,_ := vars["devicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddDevicesToTenant(tenantId, devicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more devicesIds as a Devices from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDevicesFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	devicesIds,_ := vars["devicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveDevicesFromTenant(tenantId, devicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataRetentionPoliciesIds as a DataRetentionPolicies to a Tenant
	//----------------------------------------------------------------------------
func AddDataRetentionPoliciesToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataRetentionPoliciesIds,_ := vars["dataRetentionPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddDataRetentionPoliciesToTenant(tenantId, dataRetentionPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataRetentionPoliciesIds as a DataRetentionPolicies from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataRetentionPoliciesFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataRetentionPoliciesIds,_ := vars["dataRetentionPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveDataRetentionPoliciesFromTenant(tenantId, dataRetentionPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more connectivityPlansIds as a ConnectivityPlans to a Tenant
	//----------------------------------------------------------------------------
func AddConnectivityPlansToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectivityPlansIds,_ := vars["connectivityPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddConnectivityPlansToTenant(tenantId, connectivityPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more connectivityPlansIds as a ConnectivityPlans from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConnectivityPlansFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectivityPlansIds,_ := vars["connectivityPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveConnectivityPlansFromTenant(tenantId, connectivityPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more simCardsIds as a SimCards to a Tenant
	//----------------------------------------------------------------------------
func AddSimCardsToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	simCardsIds,_ := vars["simCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddSimCardsToTenant(tenantId, simCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more simCardsIds as a SimCards from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSimCardsFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	simCardsIds,_ := vars["simCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveSimCardsFromTenant(tenantId, simCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more messagingEndpointsIds as a MessagingEndpoints to a Tenant
	//----------------------------------------------------------------------------
func AddMessagingEndpointsToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	messagingEndpointsIds,_ := vars["messagingEndpointsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddMessagingEndpointsToTenant(tenantId, messagingEndpointsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more messagingEndpointsIds as a MessagingEndpoints from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMessagingEndpointsFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	messagingEndpointsIds,_ := vars["messagingEndpointsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveMessagingEndpointsFromTenant(tenantId, messagingEndpointsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more accessPoliciesIds as a AccessPolicies to a Tenant
	//----------------------------------------------------------------------------
func AddAccessPoliciesToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accessPoliciesIds,_ := vars["accessPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddAccessPoliciesToTenant(tenantId, accessPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accessPoliciesIds as a AccessPolicies from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccessPoliciesFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accessPoliciesIds,_ := vars["accessPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveAccessPoliciesFromTenant(tenantId, accessPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more deviceGroupsIds as a DeviceGroups to a Tenant
	//----------------------------------------------------------------------------
func AddDeviceGroupsToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceGroupsIds,_ := vars["deviceGroupsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddDeviceGroupsToTenant(tenantId, deviceGroupsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more deviceGroupsIds as a DeviceGroups from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDeviceGroupsFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceGroupsIds,_ := vars["deviceGroupsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveDeviceGroupsFromTenant(tenantId, deviceGroupsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertRulesIds as a AlertRules to a Tenant
	//----------------------------------------------------------------------------
func AddAlertRulesToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertRulesIds,_ := vars["alertRulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddAlertRulesToTenant(tenantId, alertRulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertRulesIds as a AlertRules from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertRulesFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertRulesIds,_ := vars["alertRulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveAlertRulesFromTenant(tenantId, alertRulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more maintenanceTicketsIds as a MaintenanceTickets to a Tenant
	//----------------------------------------------------------------------------
func AddMaintenanceTicketsToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceTicketsIds,_ := vars["maintenanceTicketsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddMaintenanceTicketsToTenant(tenantId, maintenanceTicketsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more maintenanceTicketsIds as a MaintenanceTickets from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMaintenanceTicketsFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceTicketsIds,_ := vars["maintenanceTicketsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveMaintenanceTicketsFromTenant(tenantId, maintenanceTicketsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more usageRecordsIds as a UsageRecords to a Tenant
	//----------------------------------------------------------------------------
func AddUsageRecordsToTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usageRecordsIds,_ := vars["usageRecordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.AddUsageRecordsToTenant(tenantId, usageRecordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usageRecordsIds as a UsageRecords from a Tenant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsageRecordsFromTenant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usageRecordsIds,_ := vars["usageRecordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Tenant DAO
	//----------------------------------------------------------------------------
	requestResult := TenantDAO.RemoveUsageRecordsFromTenant(tenantId, usageRecordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
