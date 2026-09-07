package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MaintenanceAppointmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMaintenanceAppointment - creates a new db entry
//----------------------------------------------------------------------------
func CreateMaintenanceAppointment(obj model.MaintenanceAppointment)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a MaintenanceAppointment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MaintenanceAppointment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMaintenanceAppointment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMaintenanceAppointment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMaintenanceAppointment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MaintenanceAppointment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MaintenanceAppointment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MaintenanceAppointment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MaintenanceAppointment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMaintenanceAppointment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMaintenanceAppointment - returns all
//----------------------------------------------------------------------------
func GetAllMaintenanceAppointment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MaintenanceAppointment

	//----------------------------------------------------------------------------
	// Request the ORM to find all MaintenanceAppointment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MaintenanceAppointment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MaintenanceAppointment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMaintenanceAppointment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMaintenanceAppointment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMaintenanceAppointment(obj model.MaintenanceAppointment)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a MaintenanceAppointment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MaintenanceAppointment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMaintenanceAppointment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMaintenanceAppointment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMaintenanceAppointment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMaintenanceAppointment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MaintenanceAppointment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MaintenanceAppointment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMaintenanceAppointment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Aircraft on a MaintenanceAppointment
//----------------------------------------------------------------------------
func AssignAircraftToMaintenanceAppointment( maintenanceAppointmentId uint64, aircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceAppointment(maintenanceAppointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Aircraft

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Aircraft with a
		// matching aircraftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, aircraftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Aircraft	to the MaintenanceAppointment
			//----------------------------------------------------------------------------
			parentObj.Aircraft = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceAppointment
			//----------------------------------------------------------------------------
			return UpdateMaintenanceAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
			return utils.RequestResult{false, msg, "assignAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Aircraft on a MaintenanceAppointment
//----------------------------------------------------------------------------
func UnassignAircraftFromMaintenanceAppointment(maintenanceAppointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceAppointment(maintenanceAppointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// assign an empty Aircraft to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.Aircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.AircraftId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceAppointment
		//----------------------------------------------------------------------------
		return UpdateMaintenanceAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a MroFacility on a MaintenanceAppointment
//----------------------------------------------------------------------------
func AssignMroFacilityToMaintenanceAppointment( maintenanceAppointmentId uint64, mroFacilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceAppointment(maintenanceAppointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MROFacility

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MROFacility with a
		// matching mroFacilityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, mroFacilityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the MroFacility	to the MaintenanceAppointment
			//----------------------------------------------------------------------------
			parentObj.MroFacility = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceAppointment
			//----------------------------------------------------------------------------
			return UpdateMaintenanceAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MroFacility", mroFacilityId )
			return utils.RequestResult{false, msg, "assignMroFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a MroFacility on a MaintenanceAppointment
//----------------------------------------------------------------------------
func UnassignMroFacilityFromMaintenanceAppointment(maintenanceAppointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceAppointment(maintenanceAppointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// assign an empty MROFacility to the MroFacility
		//----------------------------------------------------------------------------
		parentObj.MroFacility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the MroFacility
		//----------------------------------------------------------------------------
		parentObj.MroFacilityId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceAppointment
		//----------------------------------------------------------------------------
		return UpdateMaintenanceAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkOrder on a MaintenanceAppointment
//----------------------------------------------------------------------------
func AssignWorkOrderToMaintenanceAppointment( maintenanceAppointmentId uint64, workOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceAppointment(maintenanceAppointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MaintenanceWorkOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder with a
		// matching workOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkOrder	to the MaintenanceAppointment
			//----------------------------------------------------------------------------
			parentObj.WorkOrder = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceAppointment
			//----------------------------------------------------------------------------
			return UpdateMaintenanceAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrder", workOrderId )
			return utils.RequestResult{false, msg, "assignWorkOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkOrder on a MaintenanceAppointment
//----------------------------------------------------------------------------
func UnassignWorkOrderFromMaintenanceAppointment(maintenanceAppointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceAppointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceAppointment(maintenanceAppointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceAppointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceAppointment)

		//----------------------------------------------------------------------------
		// assign an empty MaintenanceWorkOrder to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrderId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceAppointment
		//----------------------------------------------------------------------------
		return UpdateMaintenanceAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}


