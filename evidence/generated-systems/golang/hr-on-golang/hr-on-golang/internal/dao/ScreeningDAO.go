package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ScreeningDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateScreening - creates a new db entry
//----------------------------------------------------------------------------
func CreateScreening(obj model.Screening)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Screening with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Screening", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateScreening", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetScreening - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetScreening(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Screening

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Screening with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Screening using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Screening using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetScreening", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllScreening - returns all
//----------------------------------------------------------------------------
func GetAllScreening()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Screening

	//----------------------------------------------------------------------------
	// Request the ORM to find all Screening
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Screening" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Screening", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllScreening", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateScreening - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateScreening(obj model.Screening)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Screening using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Screening using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateScreening", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteScreening - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteScreening(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetScreening(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Screening)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Screening using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Screening using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteScreening", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Application on a Screening
//----------------------------------------------------------------------------
func AssignApplicationToScreening( screeningId uint64, applicationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreening(screeningId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Screening)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.JobApplication

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a JobApplication with a
		// matching applicationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, applicationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Application	to the Screening
			//----------------------------------------------------------------------------
			parentObj.Application = &childObj

			//----------------------------------------------------------------------------
			// save the Screening
			//----------------------------------------------------------------------------
			return UpdateScreening(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Application", applicationId )
			return utils.RequestResult{false, msg, "assignApplication", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Application on a Screening
//----------------------------------------------------------------------------
func UnassignApplicationFromScreening(screeningId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreening(screeningId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Screening)

		//----------------------------------------------------------------------------
		// assign an empty JobApplication to the Application
		//----------------------------------------------------------------------------
		parentObj.Application = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Application
		//----------------------------------------------------------------------------
		parentObj.ApplicationId = nil;

		//----------------------------------------------------------------------------
		// save the Screening
		//----------------------------------------------------------------------------
		return UpdateScreening(parentObj)

	} else {
		return parentRequestResult
	}

}


