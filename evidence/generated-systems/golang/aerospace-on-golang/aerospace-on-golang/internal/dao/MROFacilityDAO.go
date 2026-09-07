package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MROFacilityDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMROFacility - creates a new db entry
//----------------------------------------------------------------------------
func CreateMROFacility(obj model.MROFacility)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MROFacility with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MROFacility", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMROFacility", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMROFacility - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMROFacility(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MROFacility

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MROFacility with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MROFacility using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MROFacility using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMROFacility", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMROFacility - returns all
//----------------------------------------------------------------------------
func GetAllMROFacility()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MROFacility

	//----------------------------------------------------------------------------
	// Request the ORM to find all MROFacility
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MROFacility" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MROFacility", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMROFacility", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMROFacility - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMROFacility(obj model.MROFacility)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MROFacility using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MROFacility using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMROFacility", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMROFacility - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMROFacility(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MROFacility with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMROFacility(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MROFacility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MROFacility)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MROFacility using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MROFacility using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMROFacility", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more appointmentsIds as a Appointments to a MROFacility
//----------------------------------------------------------------------------
func AddAppointmentsToMROFacility ( mROFacilityId uint64, appointmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MROFacility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMROFacility(mROFacilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MROFacility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MROFacility)

		// slice the ids on comma with no spaces
		ids := strings.Split( appointmentsIds, ",")

		for _, appointmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceAppointment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceAppointment
			// with a matching appointmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , appointmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Appointments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Appointments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointments", appointmentsId )
				return utils.RequestResult{false, msg, "unassignAppointments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MROFacility from the gorm
		//----------------------------------------------------------------------------
		return GetMROFacility(mROFacilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more appointmentsIds as a Appointments from a MROFacility
//----------------------------------------------------------------------------
func RemoveAppointmentsFromMROFacility( mROFacilityId uint64, appointmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MROFacility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMROFacility(mROFacilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MROFacility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MROFacility)

		// slice the ids on comma with no spaces
		ids := strings.Split( appointmentsIds, ",")

		for _, appointmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceAppointment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceAppointment
			// with a matching appointmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , appointmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MaintenanceAppointmentObj from the Appointments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Appointments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointments", appointmentsId )
				return utils.RequestResult{false, msg, "removeAppointments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MROFacility from the gorm
		//----------------------------------------------------------------------------
		return GetMROFacility(mROFacilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more workOrdersIds as a WorkOrders to a MROFacility
//----------------------------------------------------------------------------
func AddWorkOrdersToMROFacility ( mROFacilityId uint64, workOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MROFacility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMROFacility(mROFacilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MROFacility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MROFacility)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceWorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the WorkOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "unassignWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MROFacility from the gorm
		//----------------------------------------------------------------------------
		return GetMROFacility(mROFacilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workOrdersIds as a WorkOrders from a MROFacility
//----------------------------------------------------------------------------
func RemoveWorkOrdersFromMROFacility( mROFacilityId uint64, workOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MROFacility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMROFacility(mROFacilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MROFacility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MROFacility)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceWorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MaintenanceWorkOrderObj from the WorkOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "removeWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MROFacility from the gorm
		//----------------------------------------------------------------------------
		return GetMROFacility(mROFacilityId)

	} else {
		return parentRequestResult
	}
}

