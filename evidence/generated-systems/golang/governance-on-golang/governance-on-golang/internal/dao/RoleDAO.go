package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RoleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRole - creates a new db entry
//----------------------------------------------------------------------------
func CreateRole(obj model.Role)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Role with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Role", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRole", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRole - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRole(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Role

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Role with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Role using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Role using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRole", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRole - returns all
//----------------------------------------------------------------------------
func GetAllRole()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Role

	//----------------------------------------------------------------------------
	// Request the ORM to find all Role
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Role" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Role", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRole", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRole - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRole(obj model.Role)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Role using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Role using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRole", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRole - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRole(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Role with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRole(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Role so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Role)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Role using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Role using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRole", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more assignmentsIds as a Assignments to a Role
//----------------------------------------------------------------------------
func AddAssignmentsToRole ( roleId uint64, assignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Role with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRole(roleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Role so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Role)

		// slice the ids on comma with no spaces
		ids := strings.Split( assignmentsIds, ",")

		for _, assignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RoleAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RoleAssignment
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
		// retrieve the modified Role from the gorm
		//----------------------------------------------------------------------------
		return GetRole(roleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more assignmentsIds as a Assignments from a Role
//----------------------------------------------------------------------------
func RemoveAssignmentsFromRole( roleId uint64, assignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Role with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRole(roleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Role so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Role)

		// slice the ids on comma with no spaces
		ids := strings.Split( assignmentsIds, ",")

		for _, assignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RoleAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RoleAssignment
			// with a matching assignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RoleAssignmentObj from the Assignments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assignments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assignments", assignmentsId )
				return utils.RequestResult{false, msg, "removeAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Role from the gorm
		//----------------------------------------------------------------------------
		return GetRole(roleId)

	} else {
		return parentRequestResult
	}
}

