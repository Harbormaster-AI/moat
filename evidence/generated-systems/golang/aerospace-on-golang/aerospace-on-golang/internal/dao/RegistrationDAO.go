package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RegistrationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRegistration - creates a new db entry
//----------------------------------------------------------------------------
func CreateRegistration(obj model.Registration)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Registration with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Registration", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRegistration", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRegistration - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRegistration(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Registration

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Registration with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Registration using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Registration using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRegistration", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRegistration - returns all
//----------------------------------------------------------------------------
func GetAllRegistration()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Registration

	//----------------------------------------------------------------------------
	// Request the ORM to find all Registration
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Registration" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Registration", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRegistration", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRegistration - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRegistration(obj model.Registration)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Registration using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Registration using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRegistration", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRegistration - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRegistration(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Registration with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRegistration(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Registration so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Registration)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Registration using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Registration using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRegistration", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Aircraft on a Registration
//----------------------------------------------------------------------------
func AssignAircraftToRegistration( registrationId uint64, aircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Registration with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRegistration(registrationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Registration so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Registration)

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
			// assign the Aircraft	to the Registration
			//----------------------------------------------------------------------------
			parentObj.Aircraft = &childObj

			//----------------------------------------------------------------------------
			// save the Registration
			//----------------------------------------------------------------------------
			return UpdateRegistration(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
			return utils.RequestResult{false, msg, "assignAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Aircraft on a Registration
//----------------------------------------------------------------------------
func UnassignAircraftFromRegistration(registrationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Registration with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRegistration(registrationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Registration so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Registration)

		//----------------------------------------------------------------------------
		// assign an empty Aircraft to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.Aircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.AircraftId = nil;

		//----------------------------------------------------------------------------
		// save the Registration
		//----------------------------------------------------------------------------
		return UpdateRegistration(parentObj)

	} else {
		return parentRequestResult
	}

}


