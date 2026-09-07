package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OnboardingTaskDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOnboardingTask - creates a new db entry
//----------------------------------------------------------------------------
func CreateOnboardingTask(obj model.OnboardingTask)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a OnboardingTask with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a OnboardingTask", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOnboardingTask", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOnboardingTask - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOnboardingTask(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.OnboardingTask

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a OnboardingTask with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a OnboardingTask using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a OnboardingTask using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOnboardingTask", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOnboardingTask - returns all
//----------------------------------------------------------------------------
func GetAllOnboardingTask()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.OnboardingTask

	//----------------------------------------------------------------------------
	// Request the ORM to find all OnboardingTask
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all OnboardingTask" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all OnboardingTask", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOnboardingTask", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOnboardingTask - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOnboardingTask(obj model.OnboardingTask)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a OnboardingTask using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a OnboardingTask using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOnboardingTask", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOnboardingTask - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOnboardingTask(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOnboardingTask(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a OnboardingTask using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a OnboardingTask using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOnboardingTask", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a OnboardingTask
//----------------------------------------------------------------------------
func AssignEmployeeToOnboardingTask( onboardingTaskId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching employeeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, employeeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Employee	to the OnboardingTask
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the OnboardingTask
			//----------------------------------------------------------------------------
			return UpdateOnboardingTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a OnboardingTask
//----------------------------------------------------------------------------
func UnassignEmployeeFromOnboardingTask(onboardingTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the OnboardingTask
		//----------------------------------------------------------------------------
		return UpdateOnboardingTask(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AssignedTo on a OnboardingTask
//----------------------------------------------------------------------------
func AssignAssignedToToOnboardingTask( onboardingTaskId uint64, assignedToId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching assignedToId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, assignedToId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AssignedTo	to the OnboardingTask
			//----------------------------------------------------------------------------
			parentObj.AssignedTo = &childObj

			//----------------------------------------------------------------------------
			// save the OnboardingTask
			//----------------------------------------------------------------------------
			return UpdateOnboardingTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AssignedTo", assignedToId )
			return utils.RequestResult{false, msg, "assignAssignedTo", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AssignedTo on a OnboardingTask
//----------------------------------------------------------------------------
func UnassignAssignedToFromOnboardingTask(onboardingTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the AssignedTo
		//----------------------------------------------------------------------------
		parentObj.AssignedTo = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AssignedTo
		//----------------------------------------------------------------------------
		parentObj.AssignedToId = nil;

		//----------------------------------------------------------------------------
		// save the OnboardingTask
		//----------------------------------------------------------------------------
		return UpdateOnboardingTask(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a RelatedOffer on a OnboardingTask
//----------------------------------------------------------------------------
func AssignRelatedOfferToOnboardingTask( onboardingTaskId uint64, relatedOfferId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Offer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Offer with a
		// matching relatedOfferId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, relatedOfferId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RelatedOffer	to the OnboardingTask
			//----------------------------------------------------------------------------
			parentObj.RelatedOffer = &childObj

			//----------------------------------------------------------------------------
			// save the OnboardingTask
			//----------------------------------------------------------------------------
			return UpdateOnboardingTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedOffer", relatedOfferId )
			return utils.RequestResult{false, msg, "assignRelatedOffer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RelatedOffer on a OnboardingTask
//----------------------------------------------------------------------------
func UnassignRelatedOfferFromOnboardingTask(onboardingTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		//----------------------------------------------------------------------------
		// assign an empty Offer to the RelatedOffer
		//----------------------------------------------------------------------------
		parentObj.RelatedOffer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RelatedOffer
		//----------------------------------------------------------------------------
		parentObj.RelatedOfferId = nil;

		//----------------------------------------------------------------------------
		// save the OnboardingTask
		//----------------------------------------------------------------------------
		return UpdateOnboardingTask(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more dependenciesIds as a Dependencies to a OnboardingTask
//----------------------------------------------------------------------------
func AddDependenciesToOnboardingTask ( onboardingTaskId uint64, dependenciesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		// slice the ids on comma with no spaces
		ids := strings.Split( dependenciesIds, ",")

		for _, dependenciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OnboardingTask

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OnboardingTask
			// with a matching dependenciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dependenciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dependencies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dependencies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dependencies", dependenciesId )
				return utils.RequestResult{false, msg, "unassignDependencies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified OnboardingTask from the gorm
		//----------------------------------------------------------------------------
		return GetOnboardingTask(onboardingTaskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dependenciesIds as a Dependencies from a OnboardingTask
//----------------------------------------------------------------------------
func RemoveDependenciesFromOnboardingTask( onboardingTaskId uint64, dependenciesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the OnboardingTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOnboardingTask(onboardingTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OnboardingTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OnboardingTask)

		// slice the ids on comma with no spaces
		ids := strings.Split( dependenciesIds, ",")

		for _, dependenciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OnboardingTask

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OnboardingTask
			// with a matching dependenciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dependenciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OnboardingTaskObj from the Dependencies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dependencies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dependencies", dependenciesId )
				return utils.RequestResult{false, msg, "removeDependencies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified OnboardingTask from the gorm
		//----------------------------------------------------------------------------
		return GetOnboardingTask(onboardingTaskId)

	} else {
		return parentRequestResult
	}
}

