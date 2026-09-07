package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PersonDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePerson - creates a new db entry
//----------------------------------------------------------------------------
func CreatePerson(obj model.Person)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Person with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Person", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePerson", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPerson - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPerson(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Person

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Person with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Person using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Person using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPerson", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPerson - returns all
//----------------------------------------------------------------------------
func GetAllPerson()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Person

	//----------------------------------------------------------------------------
	// Request the ORM to find all Person
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Person" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Person", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPerson", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePerson - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePerson(obj model.Person)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Person using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Person using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePerson", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePerson - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePerson(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPerson(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Person)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Person using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Person using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePerson", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more roleAssignmentsIds as a RoleAssignments to a Person
//----------------------------------------------------------------------------
func AddRoleAssignmentsToPerson ( personId uint64, roleAssignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerson(personId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Person)

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
		// retrieve the modified Person from the gorm
		//----------------------------------------------------------------------------
		return GetPerson(personId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more roleAssignmentsIds as a RoleAssignments from a Person
//----------------------------------------------------------------------------
func RemoveRoleAssignmentsFromPerson( personId uint64, roleAssignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerson(personId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Person)

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
		// retrieve the modified Person from the gorm
		//----------------------------------------------------------------------------
		return GetPerson(personId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ownedPoliciesIds as a OwnedPolicies to a Person
//----------------------------------------------------------------------------
func AddOwnedPoliciesToPerson ( personId uint64, ownedPoliciesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerson(personId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Person)

		// slice the ids on comma with no spaces
		ids := strings.Split( ownedPoliciesIds, ",")

		for _, ownedPoliciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching ownedPoliciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedPoliciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OwnedPolicies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedPolicies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedPolicies", ownedPoliciesId )
				return utils.RequestResult{false, msg, "unassignOwnedPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Person from the gorm
		//----------------------------------------------------------------------------
		return GetPerson(personId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ownedPoliciesIds as a OwnedPolicies from a Person
//----------------------------------------------------------------------------
func RemoveOwnedPoliciesFromPerson( personId uint64, ownedPoliciesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerson(personId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Person)

		// slice the ids on comma with no spaces
		ids := strings.Split( ownedPoliciesIds, ",")

		for _, ownedPoliciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching ownedPoliciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedPoliciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the OwnedPolicies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedPolicies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedPolicies", ownedPoliciesId )
				return utils.RequestResult{false, msg, "removeOwnedPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Person from the gorm
		//----------------------------------------------------------------------------
		return GetPerson(personId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more correctiveActionsIds as a CorrectiveActions to a Person
//----------------------------------------------------------------------------
func AddCorrectiveActionsToPerson ( personId uint64, correctiveActionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerson(personId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Person)

		// slice the ids on comma with no spaces
		ids := strings.Split( correctiveActionsIds, ",")

		for _, correctiveActionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CorrectiveAction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CorrectiveAction
			// with a matching correctiveActionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , correctiveActionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CorrectiveActions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CorrectiveActions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveActions", correctiveActionsId )
				return utils.RequestResult{false, msg, "unassignCorrectiveActions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Person from the gorm
		//----------------------------------------------------------------------------
		return GetPerson(personId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more correctiveActionsIds as a CorrectiveActions from a Person
//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromPerson( personId uint64, correctiveActionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Person with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerson(personId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Person so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Person)

		// slice the ids on comma with no spaces
		ids := strings.Split( correctiveActionsIds, ",")

		for _, correctiveActionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CorrectiveAction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CorrectiveAction
			// with a matching correctiveActionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , correctiveActionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CorrectiveActionObj from the CorrectiveActions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CorrectiveActions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveActions", correctiveActionsId )
				return utils.RequestResult{false, msg, "removeCorrectiveActions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Person from the gorm
		//----------------------------------------------------------------------------
		return GetPerson(personId)

	} else {
		return parentRequestResult
	}
}

