package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PositionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePosition - creates a new db entry
//----------------------------------------------------------------------------
func CreatePosition(obj model.Position)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Position with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Position", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePosition", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPosition - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPosition(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Position

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Position with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Position using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Position using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPosition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPosition - returns all
//----------------------------------------------------------------------------
func GetAllPosition()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Position

	//----------------------------------------------------------------------------
	// Request the ORM to find all Position
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Position" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Position", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPosition", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePosition - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePosition(obj model.Position)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Position using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Position using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePosition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePosition - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePosition(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPosition(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Position using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Position using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePosition", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Department on a Position
//----------------------------------------------------------------------------
func AssignDepartmentToPosition( positionId uint64, departmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Department

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Department with a
		// matching departmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, departmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Department	to the Position
			//----------------------------------------------------------------------------
			parentObj.Department = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Department", departmentId )
			return utils.RequestResult{false, msg, "assignDepartment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Department on a Position
//----------------------------------------------------------------------------
func UnassignDepartmentFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty Department to the Department
		//----------------------------------------------------------------------------
		parentObj.Department = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Department
		//----------------------------------------------------------------------------
		parentObj.DepartmentId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a JobProfile on a Position
//----------------------------------------------------------------------------
func AssignJobProfileToPosition( positionId uint64, jobProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.JobProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a JobProfile with a
		// matching jobProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, jobProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the JobProfile	to the Position
			//----------------------------------------------------------------------------
			parentObj.JobProfile = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobProfile", jobProfileId )
			return utils.RequestResult{false, msg, "assignJobProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a JobProfile on a Position
//----------------------------------------------------------------------------
func UnassignJobProfileFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty JobProfile to the JobProfile
		//----------------------------------------------------------------------------
		parentObj.JobProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the JobProfile
		//----------------------------------------------------------------------------
		parentObj.JobProfileId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CostCenter on a Position
//----------------------------------------------------------------------------
func AssignCostCenterToPosition( positionId uint64, costCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CostCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CostCenter with a
		// matching costCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, costCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CostCenter	to the Position
			//----------------------------------------------------------------------------
			parentObj.CostCenter = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CostCenter", costCenterId )
			return utils.RequestResult{false, msg, "assignCostCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CostCenter on a Position
//----------------------------------------------------------------------------
func UnassignCostCenterFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty CostCenter to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenterId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a Position
//----------------------------------------------------------------------------
func AssignLocationToPosition( positionId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Location

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Location with a
		// matching locationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, locationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Location	to the Position
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a Position
//----------------------------------------------------------------------------
func UnassignLocationFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty Location to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ManagerPosition on a Position
//----------------------------------------------------------------------------
func AssignManagerPositionToPosition( positionId uint64, managerPositionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Position

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Position with a
		// matching managerPositionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, managerPositionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ManagerPosition	to the Position
			//----------------------------------------------------------------------------
			parentObj.ManagerPosition = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ManagerPosition", managerPositionId )
			return utils.RequestResult{false, msg, "assignManagerPosition", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ManagerPosition on a Position
//----------------------------------------------------------------------------
func UnassignManagerPositionFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty Position to the ManagerPosition
		//----------------------------------------------------------------------------
		parentObj.ManagerPosition = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ManagerPosition
		//----------------------------------------------------------------------------
		parentObj.ManagerPositionId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more directReportsIds as a DirectReports to a Position
//----------------------------------------------------------------------------
func AddDirectReportsToPosition ( positionId uint64, directReportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		// slice the ids on comma with no spaces
		ids := strings.Split( directReportsIds, ",")

		for _, directReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching directReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , directReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DirectReports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DirectReports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DirectReports", directReportsId )
				return utils.RequestResult{false, msg, "unassignDirectReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Position from the gorm
		//----------------------------------------------------------------------------
		return GetPosition(positionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more directReportsIds as a DirectReports from a Position
//----------------------------------------------------------------------------
func RemoveDirectReportsFromPosition( positionId uint64, directReportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		// slice the ids on comma with no spaces
		ids := strings.Split( directReportsIds, ",")

		for _, directReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching directReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , directReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PositionObj from the DirectReports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DirectReports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DirectReports", directReportsId )
				return utils.RequestResult{false, msg, "removeDirectReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Position from the gorm
		//----------------------------------------------------------------------------
		return GetPosition(positionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more assignmentsIds as a Assignments to a Position
//----------------------------------------------------------------------------
func AddAssignmentsToPosition ( positionId uint64, assignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		// slice the ids on comma with no spaces
		ids := strings.Split( assignmentsIds, ",")

		for _, assignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentAssignment
			// with a matching assignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Assignments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assignments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assignments", assignmentsId )
				return utils.RequestResult{false, msg, "unassignAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Position from the gorm
		//----------------------------------------------------------------------------
		return GetPosition(positionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more assignmentsIds as a Assignments from a Position
//----------------------------------------------------------------------------
func RemoveAssignmentsFromPosition( positionId uint64, assignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		// slice the ids on comma with no spaces
		ids := strings.Split( assignmentsIds, ",")

		for _, assignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentAssignment
			// with a matching assignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmploymentAssignmentObj from the Assignments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assignments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assignments", assignmentsId )
				return utils.RequestResult{false, msg, "removeAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Position from the gorm
		//----------------------------------------------------------------------------
		return GetPosition(positionId)

	} else {
		return parentRequestResult
	}
}

