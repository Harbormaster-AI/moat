package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ShiftDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateShift - creates a new db entry
//----------------------------------------------------------------------------
func CreateShift(obj model.Shift)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Shift with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Shift", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateShift", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetShift - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetShift(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Shift

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Shift with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Shift using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Shift using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetShift", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllShift - returns all
//----------------------------------------------------------------------------
func GetAllShift()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Shift

	//----------------------------------------------------------------------------
	// Request the ORM to find all Shift
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Shift" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Shift", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllShift", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateShift - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateShift(obj model.Shift)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Shift using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Shift using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateShift", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteShift - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteShift(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Shift with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetShift(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Shift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Shift)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Shift using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Shift using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteShift", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Plant on a Shift
//----------------------------------------------------------------------------
func AssignPlantToShift( shiftId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Shift with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShift(shiftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Shift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Shift)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Plant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Plant with a
		// matching plantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, plantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plant	to the Shift
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the Shift
			//----------------------------------------------------------------------------
			return UpdateShift(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a Shift
//----------------------------------------------------------------------------
func UnassignPlantFromShift(shiftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Shift with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShift(shiftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Shift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Shift)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the Shift
		//----------------------------------------------------------------------------
		return UpdateShift(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more assignmentsIds as a Assignments to a Shift
//----------------------------------------------------------------------------
func AddAssignmentsToShift ( shiftId uint64, assignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Shift with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShift(shiftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Shift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Shift)

		// slice the ids on comma with no spaces
		ids := strings.Split( assignmentsIds, ",")

		for _, assignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ShiftAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ShiftAssignment
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
		// retrieve the modified Shift from the gorm
		//----------------------------------------------------------------------------
		return GetShift(shiftId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more assignmentsIds as a Assignments from a Shift
//----------------------------------------------------------------------------
func RemoveAssignmentsFromShift( shiftId uint64, assignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Shift with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShift(shiftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Shift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Shift)

		// slice the ids on comma with no spaces
		ids := strings.Split( assignmentsIds, ",")

		for _, assignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ShiftAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ShiftAssignment
			// with a matching assignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ShiftAssignmentObj from the Assignments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assignments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assignments", assignmentsId )
				return utils.RequestResult{false, msg, "removeAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Shift from the gorm
		//----------------------------------------------------------------------------
		return GetShift(shiftId)

	} else {
		return parentRequestResult
	}
}

