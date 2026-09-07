package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MaintenanceWorkOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMaintenanceWorkOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateMaintenanceWorkOrder(obj model.MaintenanceWorkOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MaintenanceWorkOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MaintenanceWorkOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMaintenanceWorkOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMaintenanceWorkOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMaintenanceWorkOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MaintenanceWorkOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MaintenanceWorkOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MaintenanceWorkOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMaintenanceWorkOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMaintenanceWorkOrder - returns all
//----------------------------------------------------------------------------
func GetAllMaintenanceWorkOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MaintenanceWorkOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all MaintenanceWorkOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MaintenanceWorkOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MaintenanceWorkOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMaintenanceWorkOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMaintenanceWorkOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMaintenanceWorkOrder(obj model.MaintenanceWorkOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MaintenanceWorkOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MaintenanceWorkOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMaintenanceWorkOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMaintenanceWorkOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMaintenanceWorkOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMaintenanceWorkOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MaintenanceWorkOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MaintenanceWorkOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MaintenanceWorkOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMaintenanceWorkOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Aircraft on a MaintenanceWorkOrder
//----------------------------------------------------------------------------
func AssignAircraftToMaintenanceWorkOrder( maintenanceWorkOrderId uint64, aircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceWorkOrder(maintenanceWorkOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceWorkOrder)

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
			// assign the Aircraft	to the MaintenanceWorkOrder
			//----------------------------------------------------------------------------
			parentObj.Aircraft = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceWorkOrder
			//----------------------------------------------------------------------------
			return UpdateMaintenanceWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
			return utils.RequestResult{false, msg, "assignAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Aircraft on a MaintenanceWorkOrder
//----------------------------------------------------------------------------
func UnassignAircraftFromMaintenanceWorkOrder(maintenanceWorkOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceWorkOrder(maintenanceWorkOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceWorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty Aircraft to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.Aircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.AircraftId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceWorkOrder
		//----------------------------------------------------------------------------
		return UpdateMaintenanceWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AirworthinessDirective on a MaintenanceWorkOrder
//----------------------------------------------------------------------------
func AssignAirworthinessDirectiveToMaintenanceWorkOrder( maintenanceWorkOrderId uint64, airworthinessDirectiveId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceWorkOrder(maintenanceWorkOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceWorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AirworthinessDirective

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AirworthinessDirective with a
		// matching airworthinessDirectiveId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, airworthinessDirectiveId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AirworthinessDirective	to the MaintenanceWorkOrder
			//----------------------------------------------------------------------------
			parentObj.AirworthinessDirective = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceWorkOrder
			//----------------------------------------------------------------------------
			return UpdateMaintenanceWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AirworthinessDirective", airworthinessDirectiveId )
			return utils.RequestResult{false, msg, "assignAirworthinessDirective", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AirworthinessDirective on a MaintenanceWorkOrder
//----------------------------------------------------------------------------
func UnassignAirworthinessDirectiveFromMaintenanceWorkOrder(maintenanceWorkOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceWorkOrder(maintenanceWorkOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceWorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty AirworthinessDirective to the AirworthinessDirective
		//----------------------------------------------------------------------------
		parentObj.AirworthinessDirective = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AirworthinessDirective
		//----------------------------------------------------------------------------
		parentObj.AirworthinessDirectiveId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceWorkOrder
		//----------------------------------------------------------------------------
		return UpdateMaintenanceWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ServiceBulletin on a MaintenanceWorkOrder
//----------------------------------------------------------------------------
func AssignServiceBulletinToMaintenanceWorkOrder( maintenanceWorkOrderId uint64, serviceBulletinId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceWorkOrder(maintenanceWorkOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceWorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ServiceBulletin

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ServiceBulletin with a
		// matching serviceBulletinId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, serviceBulletinId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ServiceBulletin	to the MaintenanceWorkOrder
			//----------------------------------------------------------------------------
			parentObj.ServiceBulletin = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceWorkOrder
			//----------------------------------------------------------------------------
			return UpdateMaintenanceWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ServiceBulletin", serviceBulletinId )
			return utils.RequestResult{false, msg, "assignServiceBulletin", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ServiceBulletin on a MaintenanceWorkOrder
//----------------------------------------------------------------------------
func UnassignServiceBulletinFromMaintenanceWorkOrder(maintenanceWorkOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceWorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceWorkOrder(maintenanceWorkOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceWorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceWorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty ServiceBulletin to the ServiceBulletin
		//----------------------------------------------------------------------------
		parentObj.ServiceBulletin = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ServiceBulletin
		//----------------------------------------------------------------------------
		parentObj.ServiceBulletinId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceWorkOrder
		//----------------------------------------------------------------------------
		return UpdateMaintenanceWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


