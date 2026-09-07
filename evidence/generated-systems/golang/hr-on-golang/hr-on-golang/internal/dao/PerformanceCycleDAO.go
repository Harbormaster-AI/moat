package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PerformanceCycleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePerformanceCycle - creates a new db entry
//----------------------------------------------------------------------------
func CreatePerformanceCycle(obj model.PerformanceCycle)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PerformanceCycle with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PerformanceCycle", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePerformanceCycle", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPerformanceCycle - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPerformanceCycle(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PerformanceCycle

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PerformanceCycle with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PerformanceCycle using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PerformanceCycle using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPerformanceCycle", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPerformanceCycle - returns all
//----------------------------------------------------------------------------
func GetAllPerformanceCycle()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PerformanceCycle

	//----------------------------------------------------------------------------
	// Request the ORM to find all PerformanceCycle
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PerformanceCycle" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PerformanceCycle", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPerformanceCycle", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePerformanceCycle - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePerformanceCycle(obj model.PerformanceCycle)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PerformanceCycle using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PerformanceCycle using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePerformanceCycle", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePerformanceCycle - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePerformanceCycle(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPerformanceCycle(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PerformanceCycle)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PerformanceCycle using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PerformanceCycle using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePerformanceCycle", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a PerformanceCycle
//----------------------------------------------------------------------------
func AssignOrganizationToPerformanceCycle( performanceCycleId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceCycle(performanceCycleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceCycle)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the PerformanceCycle
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceCycle
			//----------------------------------------------------------------------------
			return UpdatePerformanceCycle(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a PerformanceCycle
//----------------------------------------------------------------------------
func UnassignOrganizationFromPerformanceCycle(performanceCycleId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceCycle(performanceCycleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceCycle)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceCycle
		//----------------------------------------------------------------------------
		return UpdatePerformanceCycle(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more reviewsIds as a Reviews to a PerformanceCycle
//----------------------------------------------------------------------------
func AddReviewsToPerformanceCycle ( performanceCycleId uint64, reviewsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceCycle(performanceCycleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceCycle)

		// slice the ids on comma with no spaces
		ids := strings.Split( reviewsIds, ",")

		for _, reviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceReview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceReview
			// with a matching reviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reviews using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reviews").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reviews", reviewsId )
				return utils.RequestResult{false, msg, "unassignReviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PerformanceCycle from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceCycle(performanceCycleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reviewsIds as a Reviews from a PerformanceCycle
//----------------------------------------------------------------------------
func RemoveReviewsFromPerformanceCycle( performanceCycleId uint64, reviewsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceCycle(performanceCycleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceCycle)

		// slice the ids on comma with no spaces
		ids := strings.Split( reviewsIds, ",")

		for _, reviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceReview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceReview
			// with a matching reviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PerformanceReviewObj from the Reviews array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reviews").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reviews", reviewsId )
				return utils.RequestResult{false, msg, "removeReviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PerformanceCycle from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceCycle(performanceCycleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more goalsIds as a Goals to a PerformanceCycle
//----------------------------------------------------------------------------
func AddGoalsToPerformanceCycle ( performanceCycleId uint64, goalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceCycle(performanceCycleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceCycle)

		// slice the ids on comma with no spaces
		ids := strings.Split( goalsIds, ",")

		for _, goalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Goal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Goal
			// with a matching goalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , goalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Goals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Goals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Goals", goalsId )
				return utils.RequestResult{false, msg, "unassignGoals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PerformanceCycle from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceCycle(performanceCycleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more goalsIds as a Goals from a PerformanceCycle
//----------------------------------------------------------------------------
func RemoveGoalsFromPerformanceCycle( performanceCycleId uint64, goalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PerformanceCycle with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceCycle(performanceCycleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceCycle so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceCycle)

		// slice the ids on comma with no spaces
		ids := strings.Split( goalsIds, ",")

		for _, goalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Goal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Goal
			// with a matching goalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , goalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GoalObj from the Goals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Goals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Goals", goalsId )
				return utils.RequestResult{false, msg, "removeGoals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PerformanceCycle from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceCycle(performanceCycleId)

	} else {
		return parentRequestResult
	}
}

