package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PerformanceReviewDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePerformanceReview - creates a new db entry
//----------------------------------------------------------------------------
func CreatePerformanceReview(obj model.PerformanceReview)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PerformanceReview with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PerformanceReview", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePerformanceReview", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPerformanceReview - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPerformanceReview(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PerformanceReview

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PerformanceReview with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PerformanceReview using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PerformanceReview using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPerformanceReview", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPerformanceReview - returns all
//----------------------------------------------------------------------------
func GetAllPerformanceReview()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PerformanceReview

	//----------------------------------------------------------------------------
	// Request the ORM to find all PerformanceReview
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PerformanceReview" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PerformanceReview", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPerformanceReview", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePerformanceReview - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePerformanceReview(obj model.PerformanceReview)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PerformanceReview using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PerformanceReview using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePerformanceReview", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePerformanceReview - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePerformanceReview(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPerformanceReview(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PerformanceReview)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PerformanceReview using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PerformanceReview using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePerformanceReview", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a PerformanceReview
//----------------------------------------------------------------------------
func AssignEmployeeToPerformanceReview( performanceReviewId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

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
			// assign the Employee	to the PerformanceReview
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceReview
			//----------------------------------------------------------------------------
			return UpdatePerformanceReview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a PerformanceReview
//----------------------------------------------------------------------------
func UnassignEmployeeFromPerformanceReview(performanceReviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceReview
		//----------------------------------------------------------------------------
		return UpdatePerformanceReview(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Reviewer on a PerformanceReview
//----------------------------------------------------------------------------
func AssignReviewerToPerformanceReview( performanceReviewId uint64, reviewerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching reviewerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, reviewerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Reviewer	to the PerformanceReview
			//----------------------------------------------------------------------------
			parentObj.Reviewer = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceReview
			//----------------------------------------------------------------------------
			return UpdatePerformanceReview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reviewer", reviewerId )
			return utils.RequestResult{false, msg, "assignReviewer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Reviewer on a PerformanceReview
//----------------------------------------------------------------------------
func UnassignReviewerFromPerformanceReview(performanceReviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Reviewer
		//----------------------------------------------------------------------------
		parentObj.Reviewer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Reviewer
		//----------------------------------------------------------------------------
		parentObj.ReviewerId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceReview
		//----------------------------------------------------------------------------
		return UpdatePerformanceReview(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Cycle on a PerformanceReview
//----------------------------------------------------------------------------
func AssignCycleToPerformanceReview( performanceReviewId uint64, cycleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PerformanceCycle

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PerformanceCycle with a
		// matching cycleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, cycleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Cycle	to the PerformanceReview
			//----------------------------------------------------------------------------
			parentObj.Cycle = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceReview
			//----------------------------------------------------------------------------
			return UpdatePerformanceReview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cycle", cycleId )
			return utils.RequestResult{false, msg, "assignCycle", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Cycle on a PerformanceReview
//----------------------------------------------------------------------------
func UnassignCycleFromPerformanceReview(performanceReviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		//----------------------------------------------------------------------------
		// assign an empty PerformanceCycle to the Cycle
		//----------------------------------------------------------------------------
		parentObj.Cycle = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Cycle
		//----------------------------------------------------------------------------
		parentObj.CycleId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceReview
		//----------------------------------------------------------------------------
		return UpdatePerformanceReview(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more competencyRatingsIds as a CompetencyRatings to a PerformanceReview
//----------------------------------------------------------------------------
func AddCompetencyRatingsToPerformanceReview ( performanceReviewId uint64, competencyRatingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		// slice the ids on comma with no spaces
		ids := strings.Split( competencyRatingsIds, ",")

		for _, competencyRatingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompetencyRating

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompetencyRating
			// with a matching competencyRatingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , competencyRatingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CompetencyRatings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompetencyRatings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompetencyRatings", competencyRatingsId )
				return utils.RequestResult{false, msg, "unassignCompetencyRatings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PerformanceReview from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceReview(performanceReviewId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more competencyRatingsIds as a CompetencyRatings from a PerformanceReview
//----------------------------------------------------------------------------
func RemoveCompetencyRatingsFromPerformanceReview( performanceReviewId uint64, competencyRatingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

		// slice the ids on comma with no spaces
		ids := strings.Split( competencyRatingsIds, ",")

		for _, competencyRatingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompetencyRating

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompetencyRating
			// with a matching competencyRatingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , competencyRatingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CompetencyRatingObj from the CompetencyRatings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompetencyRatings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompetencyRatings", competencyRatingsId )
				return utils.RequestResult{false, msg, "removeCompetencyRatings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PerformanceReview from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceReview(performanceReviewId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more goalsIds as a Goals to a PerformanceReview
//----------------------------------------------------------------------------
func AddGoalsToPerformanceReview ( performanceReviewId uint64, goalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

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
		// retrieve the modified PerformanceReview from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceReview(performanceReviewId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more goalsIds as a Goals from a PerformanceReview
//----------------------------------------------------------------------------
func RemoveGoalsFromPerformanceReview( performanceReviewId uint64, goalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PerformanceReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceReview(performanceReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceReview)

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
		// retrieve the modified PerformanceReview from the gorm
		//----------------------------------------------------------------------------
		return GetPerformanceReview(performanceReviewId)

	} else {
		return parentRequestResult
	}
}

