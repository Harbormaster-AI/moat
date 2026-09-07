package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SoftwareUpdateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSoftwareUpdate - creates a new db entry
//----------------------------------------------------------------------------
func CreateSoftwareUpdate(obj model.SoftwareUpdate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SoftwareUpdate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SoftwareUpdate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSoftwareUpdate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSoftwareUpdate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSoftwareUpdate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SoftwareUpdate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SoftwareUpdate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SoftwareUpdate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SoftwareUpdate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSoftwareUpdate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSoftwareUpdate - returns all
//----------------------------------------------------------------------------
func GetAllSoftwareUpdate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SoftwareUpdate

	//----------------------------------------------------------------------------
	// Request the ORM to find all SoftwareUpdate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SoftwareUpdate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SoftwareUpdate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSoftwareUpdate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSoftwareUpdate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSoftwareUpdate(obj model.SoftwareUpdate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SoftwareUpdate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SoftwareUpdate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSoftwareUpdate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSoftwareUpdate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSoftwareUpdate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SoftwareUpdate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSoftwareUpdate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareUpdate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SoftwareUpdate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SoftwareUpdate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SoftwareUpdate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSoftwareUpdate", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Device on a SoftwareUpdate
//----------------------------------------------------------------------------
func AssignDeviceToSoftwareUpdate( softwareUpdateId uint64, deviceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SoftwareUpdate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSoftwareUpdate(softwareUpdateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareUpdate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SoftwareUpdate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MedicalDevice

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MedicalDevice with a
		// matching deviceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, deviceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Device	to the SoftwareUpdate
			//----------------------------------------------------------------------------
			parentObj.Device = &childObj

			//----------------------------------------------------------------------------
			// save the SoftwareUpdate
			//----------------------------------------------------------------------------
			return UpdateSoftwareUpdate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Device", deviceId )
			return utils.RequestResult{false, msg, "assignDevice", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Device on a SoftwareUpdate
//----------------------------------------------------------------------------
func UnassignDeviceFromSoftwareUpdate(softwareUpdateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SoftwareUpdate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSoftwareUpdate(softwareUpdateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareUpdate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SoftwareUpdate)

		//----------------------------------------------------------------------------
		// assign an empty MedicalDevice to the Device
		//----------------------------------------------------------------------------
		parentObj.Device = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Device
		//----------------------------------------------------------------------------
		parentObj.DeviceId = nil;

		//----------------------------------------------------------------------------
		// save the SoftwareUpdate
		//----------------------------------------------------------------------------
		return UpdateSoftwareUpdate(parentObj)

	} else {
		return parentRequestResult
	}

}


