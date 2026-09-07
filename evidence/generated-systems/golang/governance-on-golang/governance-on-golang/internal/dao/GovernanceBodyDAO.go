package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing GovernanceBodyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateGovernanceBody - creates a new db entry
//----------------------------------------------------------------------------
func CreateGovernanceBody(obj model.GovernanceBody)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a GovernanceBody with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a GovernanceBody", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateGovernanceBody", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetGovernanceBody - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetGovernanceBody(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.GovernanceBody

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a GovernanceBody with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a GovernanceBody using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a GovernanceBody using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetGovernanceBody", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllGovernanceBody - returns all
//----------------------------------------------------------------------------
func GetAllGovernanceBody()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.GovernanceBody

	//----------------------------------------------------------------------------
	// Request the ORM to find all GovernanceBody
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all GovernanceBody" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all GovernanceBody", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllGovernanceBody", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateGovernanceBody - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateGovernanceBody(obj model.GovernanceBody)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a GovernanceBody using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a GovernanceBody using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateGovernanceBody", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteGovernanceBody - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteGovernanceBody(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetGovernanceBody(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.GovernanceBody)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a GovernanceBody using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a GovernanceBody using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteGovernanceBody", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a GovernanceBody
//----------------------------------------------------------------------------
func AssignOrganizationToGovernanceBody( governanceBodyId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGovernanceBody(governanceBodyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GovernanceBody)

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
			// assign the Organization	to the GovernanceBody
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the GovernanceBody
			//----------------------------------------------------------------------------
			return UpdateGovernanceBody(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a GovernanceBody
//----------------------------------------------------------------------------
func UnassignOrganizationFromGovernanceBody(governanceBodyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGovernanceBody(governanceBodyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GovernanceBody)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the GovernanceBody
		//----------------------------------------------------------------------------
		return UpdateGovernanceBody(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more roleAssignmentsIds as a RoleAssignments to a GovernanceBody
//----------------------------------------------------------------------------
func AddRoleAssignmentsToGovernanceBody ( governanceBodyId uint64, roleAssignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGovernanceBody(governanceBodyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GovernanceBody)

		// slice the ids on comma with no spaces
		ids := strings.Split( roleAssignmentsIds, ",")

		for _, roleAssignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RoleAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RoleAssignment
			// with a matching roleAssignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , roleAssignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RoleAssignments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RoleAssignments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RoleAssignments", roleAssignmentsId )
				return utils.RequestResult{false, msg, "unassignRoleAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GovernanceBody from the gorm
		//----------------------------------------------------------------------------
		return GetGovernanceBody(governanceBodyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more roleAssignmentsIds as a RoleAssignments from a GovernanceBody
//----------------------------------------------------------------------------
func RemoveRoleAssignmentsFromGovernanceBody( governanceBodyId uint64, roleAssignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGovernanceBody(governanceBodyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GovernanceBody)

		// slice the ids on comma with no spaces
		ids := strings.Split( roleAssignmentsIds, ",")

		for _, roleAssignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RoleAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RoleAssignment
			// with a matching roleAssignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , roleAssignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RoleAssignmentObj from the RoleAssignments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RoleAssignments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RoleAssignments", roleAssignmentsId )
				return utils.RequestResult{false, msg, "removeRoleAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GovernanceBody from the gorm
		//----------------------------------------------------------------------------
		return GetGovernanceBody(governanceBodyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a GovernanceBody
//----------------------------------------------------------------------------
func AddPoliciesToGovernanceBody ( governanceBodyId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGovernanceBody(governanceBodyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GovernanceBody)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GovernanceBody from the gorm
		//----------------------------------------------------------------------------
		return GetGovernanceBody(governanceBodyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a GovernanceBody
//----------------------------------------------------------------------------
func RemovePoliciesFromGovernanceBody( governanceBodyId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the GovernanceBody with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGovernanceBody(governanceBodyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GovernanceBody so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GovernanceBody)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GovernanceBody from the gorm
		//----------------------------------------------------------------------------
		return GetGovernanceBody(governanceBodyId)

	} else {
		return parentRequestResult
	}
}

