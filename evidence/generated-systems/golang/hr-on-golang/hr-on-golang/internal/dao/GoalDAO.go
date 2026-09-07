package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing GoalDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateGoal - creates a new db entry
//----------------------------------------------------------------------------
func CreateGoal(obj model.Goal)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Goal with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Goal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateGoal", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetGoal - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetGoal(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Goal

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Goal with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Goal using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Goal using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetGoal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllGoal - returns all
//----------------------------------------------------------------------------
func GetAllGoal()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Goal

	//----------------------------------------------------------------------------
	// Request the ORM to find all Goal
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Goal" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Goal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllGoal", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateGoal - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateGoal(obj model.Goal)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Goal using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Goal using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateGoal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteGoal - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteGoal(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetGoal(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Goal)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Goal using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Goal using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteGoal", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a Goal
//----------------------------------------------------------------------------
func AssignEmployeeToGoal( goalId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

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
			// assign the Employee	to the Goal
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the Goal
			//----------------------------------------------------------------------------
			return UpdateGoal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a Goal
//----------------------------------------------------------------------------
func UnassignEmployeeFromGoal(goalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the Goal
		//----------------------------------------------------------------------------
		return UpdateGoal(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Cycle on a Goal
//----------------------------------------------------------------------------
func AssignCycleToGoal( goalId uint64, cycleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

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
			// assign the Cycle	to the Goal
			//----------------------------------------------------------------------------
			parentObj.Cycle = &childObj

			//----------------------------------------------------------------------------
			// save the Goal
			//----------------------------------------------------------------------------
			return UpdateGoal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cycle", cycleId )
			return utils.RequestResult{false, msg, "assignCycle", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Cycle on a Goal
//----------------------------------------------------------------------------
func UnassignCycleFromGoal(goalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

		//----------------------------------------------------------------------------
		// assign an empty PerformanceCycle to the Cycle
		//----------------------------------------------------------------------------
		parentObj.Cycle = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Cycle
		//----------------------------------------------------------------------------
		parentObj.CycleId = nil;

		//----------------------------------------------------------------------------
		// save the Goal
		//----------------------------------------------------------------------------
		return UpdateGoal(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ParentGoal on a Goal
//----------------------------------------------------------------------------
func AssignParentGoalToGoal( goalId uint64, parentGoalId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Goal

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Goal with a
		// matching parentGoalId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, parentGoalId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ParentGoal	to the Goal
			//----------------------------------------------------------------------------
			parentObj.ParentGoal = &childObj

			//----------------------------------------------------------------------------
			// save the Goal
			//----------------------------------------------------------------------------
			return UpdateGoal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ParentGoal", parentGoalId )
			return utils.RequestResult{false, msg, "assignParentGoal", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ParentGoal on a Goal
//----------------------------------------------------------------------------
func UnassignParentGoalFromGoal(goalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

		//----------------------------------------------------------------------------
		// assign an empty Goal to the ParentGoal
		//----------------------------------------------------------------------------
		parentObj.ParentGoal = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ParentGoal
		//----------------------------------------------------------------------------
		parentObj.ParentGoalId = nil;

		//----------------------------------------------------------------------------
		// save the Goal
		//----------------------------------------------------------------------------
		return UpdateGoal(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more childGoalsIds as a ChildGoals to a Goal
//----------------------------------------------------------------------------
func AddChildGoalsToGoal ( goalId uint64, childGoalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

		// slice the ids on comma with no spaces
		ids := strings.Split( childGoalsIds, ",")

		for _, childGoalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Goal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Goal
			// with a matching childGoalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childGoalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ChildGoals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildGoals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildGoals", childGoalsId )
				return utils.RequestResult{false, msg, "unassignChildGoals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Goal from the gorm
		//----------------------------------------------------------------------------
		return GetGoal(goalId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more childGoalsIds as a ChildGoals from a Goal
//----------------------------------------------------------------------------
func RemoveChildGoalsFromGoal( goalId uint64, childGoalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Goal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoal(goalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Goal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Goal)

		// slice the ids on comma with no spaces
		ids := strings.Split( childGoalsIds, ",")

		for _, childGoalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Goal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Goal
			// with a matching childGoalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childGoalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GoalObj from the ChildGoals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildGoals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildGoals", childGoalsId )
				return utils.RequestResult{false, msg, "removeChildGoals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Goal from the gorm
		//----------------------------------------------------------------------------
		return GetGoal(goalId)

	} else {
		return parentRequestResult
	}
}

