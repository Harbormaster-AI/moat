package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SoftwareLoadDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSoftwareLoad - creates a new db entry
//----------------------------------------------------------------------------
func CreateSoftwareLoad(obj model.SoftwareLoad)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SoftwareLoad with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SoftwareLoad", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSoftwareLoad", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSoftwareLoad - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSoftwareLoad(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SoftwareLoad

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SoftwareLoad with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SoftwareLoad using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SoftwareLoad using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSoftwareLoad", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSoftwareLoad - returns all
//----------------------------------------------------------------------------
func GetAllSoftwareLoad()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SoftwareLoad

	//----------------------------------------------------------------------------
	// Request the ORM to find all SoftwareLoad
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SoftwareLoad" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SoftwareLoad", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSoftwareLoad", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSoftwareLoad - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSoftwareLoad(obj model.SoftwareLoad)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SoftwareLoad using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SoftwareLoad using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSoftwareLoad", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSoftwareLoad - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSoftwareLoad(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SoftwareLoad with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSoftwareLoad(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareLoad so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SoftwareLoad)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SoftwareLoad using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SoftwareLoad using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSoftwareLoad", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ConnectedAircraft on a SoftwareLoad
//----------------------------------------------------------------------------
func AssignConnectedAircraftToSoftwareLoad( softwareLoadId uint64, connectedAircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SoftwareLoad with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSoftwareLoad(softwareLoadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareLoad so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SoftwareLoad)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ConnectedAircraft

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ConnectedAircraft with a
		// matching connectedAircraftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, connectedAircraftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ConnectedAircraft	to the SoftwareLoad
			//----------------------------------------------------------------------------
			parentObj.ConnectedAircraft = &childObj

			//----------------------------------------------------------------------------
			// save the SoftwareLoad
			//----------------------------------------------------------------------------
			return UpdateSoftwareLoad(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConnectedAircraft", connectedAircraftId )
			return utils.RequestResult{false, msg, "assignConnectedAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ConnectedAircraft on a SoftwareLoad
//----------------------------------------------------------------------------
func UnassignConnectedAircraftFromSoftwareLoad(softwareLoadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SoftwareLoad with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSoftwareLoad(softwareLoadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareLoad so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SoftwareLoad)

		//----------------------------------------------------------------------------
		// assign an empty ConnectedAircraft to the ConnectedAircraft
		//----------------------------------------------------------------------------
		parentObj.ConnectedAircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ConnectedAircraft
		//----------------------------------------------------------------------------
		parentObj.ConnectedAircraftId = nil;

		//----------------------------------------------------------------------------
		// save the SoftwareLoad
		//----------------------------------------------------------------------------
		return UpdateSoftwareLoad(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AvionicsSuite on a SoftwareLoad
//----------------------------------------------------------------------------
func AssignAvionicsSuiteToSoftwareLoad( softwareLoadId uint64, avionicsSuiteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SoftwareLoad with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSoftwareLoad(softwareLoadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareLoad so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SoftwareLoad)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AvionicsSuite

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AvionicsSuite with a
		// matching avionicsSuiteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, avionicsSuiteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AvionicsSuite	to the SoftwareLoad
			//----------------------------------------------------------------------------
			parentObj.AvionicsSuite = &childObj

			//----------------------------------------------------------------------------
			// save the SoftwareLoad
			//----------------------------------------------------------------------------
			return UpdateSoftwareLoad(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AvionicsSuite", avionicsSuiteId )
			return utils.RequestResult{false, msg, "assignAvionicsSuite", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AvionicsSuite on a SoftwareLoad
//----------------------------------------------------------------------------
func UnassignAvionicsSuiteFromSoftwareLoad(softwareLoadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SoftwareLoad with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSoftwareLoad(softwareLoadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SoftwareLoad so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SoftwareLoad)

		//----------------------------------------------------------------------------
		// assign an empty AvionicsSuite to the AvionicsSuite
		//----------------------------------------------------------------------------
		parentObj.AvionicsSuite = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AvionicsSuite
		//----------------------------------------------------------------------------
		parentObj.AvionicsSuiteId = nil;

		//----------------------------------------------------------------------------
		// save the SoftwareLoad
		//----------------------------------------------------------------------------
		return UpdateSoftwareLoad(parentObj)

	} else {
		return parentRequestResult
	}

}


