package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DeviceCriterionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDeviceCriterion - creates a new db entry
//----------------------------------------------------------------------------
func CreateDeviceCriterion(obj model.DeviceCriterion)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DeviceCriterion with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DeviceCriterion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDeviceCriterion", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDeviceCriterion - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDeviceCriterion(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DeviceCriterion

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DeviceCriterion with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DeviceCriterion using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DeviceCriterion using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDeviceCriterion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDeviceCriterion - returns all
//----------------------------------------------------------------------------
func GetAllDeviceCriterion()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DeviceCriterion

	//----------------------------------------------------------------------------
	// Request the ORM to find all DeviceCriterion
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DeviceCriterion" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DeviceCriterion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDeviceCriterion", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDeviceCriterion - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDeviceCriterion(obj model.DeviceCriterion)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DeviceCriterion using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DeviceCriterion using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDeviceCriterion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDeviceCriterion - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDeviceCriterion(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DeviceCriterion with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDeviceCriterion(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DeviceCriterion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DeviceCriterion)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DeviceCriterion using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DeviceCriterion using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDeviceCriterion", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a TargetingProfile on a DeviceCriterion
//----------------------------------------------------------------------------
func AssignTargetingProfileToDeviceCriterion( deviceCriterionId uint64, targetingProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DeviceCriterion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeviceCriterion(deviceCriterionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DeviceCriterion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DeviceCriterion)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TargetingProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TargetingProfile with a
		// matching targetingProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, targetingProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TargetingProfile	to the DeviceCriterion
			//----------------------------------------------------------------------------
			parentObj.TargetingProfile = &childObj

			//----------------------------------------------------------------------------
			// save the DeviceCriterion
			//----------------------------------------------------------------------------
			return UpdateDeviceCriterion(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TargetingProfile", targetingProfileId )
			return utils.RequestResult{false, msg, "assignTargetingProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TargetingProfile on a DeviceCriterion
//----------------------------------------------------------------------------
func UnassignTargetingProfileFromDeviceCriterion(deviceCriterionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DeviceCriterion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeviceCriterion(deviceCriterionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DeviceCriterion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DeviceCriterion)

		//----------------------------------------------------------------------------
		// assign an empty TargetingProfile to the TargetingProfile
		//----------------------------------------------------------------------------
		parentObj.TargetingProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TargetingProfile
		//----------------------------------------------------------------------------
		parentObj.TargetingProfileId = nil;

		//----------------------------------------------------------------------------
		// save the DeviceCriterion
		//----------------------------------------------------------------------------
		return UpdateDeviceCriterion(parentObj)

	} else {
		return parentRequestResult
	}

}


