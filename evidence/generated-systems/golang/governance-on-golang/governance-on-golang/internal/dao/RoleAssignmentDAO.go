package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RoleAssignmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRoleAssignment - creates a new db entry
//----------------------------------------------------------------------------
func CreateRoleAssignment(obj model.RoleAssignment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RoleAssignment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RoleAssignment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRoleAssignment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRoleAssignment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRoleAssignment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RoleAssignment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RoleAssignment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RoleAssignment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RoleAssignment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRoleAssignment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRoleAssignment - returns all
//----------------------------------------------------------------------------
func GetAllRoleAssignment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RoleAssignment

	//----------------------------------------------------------------------------
	// Request the ORM to find all RoleAssignment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RoleAssignment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RoleAssignment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRoleAssignment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRoleAssignment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRoleAssignment(obj model.RoleAssignment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RoleAssignment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RoleAssignment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRoleAssignment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRoleAssignment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRoleAssignment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRoleAssignment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RoleAssignment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RoleAssignment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRoleAssignment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Person on a RoleAssignment
//----------------------------------------------------------------------------
func AssignPersonToRoleAssignment( roleAssignmentId uint64, personId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Person

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Person with a
		// matching personId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, personId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Person	to the RoleAssignment
			//----------------------------------------------------------------------------
			parentObj.Person = &childObj

			//----------------------------------------------------------------------------
			// save the RoleAssignment
			//----------------------------------------------------------------------------
			return UpdateRoleAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Person", personId )
			return utils.RequestResult{false, msg, "assignPerson", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Person on a RoleAssignment
//----------------------------------------------------------------------------
func UnassignPersonFromRoleAssignment(roleAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Person to the Person
		//----------------------------------------------------------------------------
		parentObj.Person = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Person
		//----------------------------------------------------------------------------
		parentObj.PersonId = nil;

		//----------------------------------------------------------------------------
		// save the RoleAssignment
		//----------------------------------------------------------------------------
		return UpdateRoleAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Role on a RoleAssignment
//----------------------------------------------------------------------------
func AssignRoleToRoleAssignment( roleAssignmentId uint64, roleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Role

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Role with a
		// matching roleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, roleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Role	to the RoleAssignment
			//----------------------------------------------------------------------------
			parentObj.Role = &childObj

			//----------------------------------------------------------------------------
			// save the RoleAssignment
			//----------------------------------------------------------------------------
			return UpdateRoleAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Role", roleId )
			return utils.RequestResult{false, msg, "assignRole", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Role on a RoleAssignment
//----------------------------------------------------------------------------
func UnassignRoleFromRoleAssignment(roleAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Role to the Role
		//----------------------------------------------------------------------------
		parentObj.Role = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Role
		//----------------------------------------------------------------------------
		parentObj.RoleId = nil;

		//----------------------------------------------------------------------------
		// save the RoleAssignment
		//----------------------------------------------------------------------------
		return UpdateRoleAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a GovernanceBody on a RoleAssignment
//----------------------------------------------------------------------------
func AssignGovernanceBodyToRoleAssignment( roleAssignmentId uint64, governanceBodyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.GovernanceBody

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a GovernanceBody with a
		// matching governanceBodyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, governanceBodyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the GovernanceBody	to the RoleAssignment
			//----------------------------------------------------------------------------
			parentObj.GovernanceBody = &childObj

			//----------------------------------------------------------------------------
			// save the RoleAssignment
			//----------------------------------------------------------------------------
			return UpdateRoleAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GovernanceBody", governanceBodyId )
			return utils.RequestResult{false, msg, "assignGovernanceBody", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a GovernanceBody on a RoleAssignment
//----------------------------------------------------------------------------
func UnassignGovernanceBodyFromRoleAssignment(roleAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// assign an empty GovernanceBody to the GovernanceBody
		//----------------------------------------------------------------------------
		parentObj.GovernanceBody = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the GovernanceBody
		//----------------------------------------------------------------------------
		parentObj.GovernanceBodyId = nil;

		//----------------------------------------------------------------------------
		// save the RoleAssignment
		//----------------------------------------------------------------------------
		return UpdateRoleAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Organization on a RoleAssignment
//----------------------------------------------------------------------------
func AssignOrganizationToRoleAssignment( roleAssignmentId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

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
			// assign the Organization	to the RoleAssignment
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the RoleAssignment
			//----------------------------------------------------------------------------
			return UpdateRoleAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a RoleAssignment
//----------------------------------------------------------------------------
func UnassignOrganizationFromRoleAssignment(roleAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RoleAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRoleAssignment(roleAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RoleAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RoleAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the RoleAssignment
		//----------------------------------------------------------------------------
		return UpdateRoleAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}


