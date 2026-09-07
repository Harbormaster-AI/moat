package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TerritoryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTerritory - creates a new db entry
//----------------------------------------------------------------------------
func CreateTerritory(obj model.Territory)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Territory with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Territory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTerritory", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTerritory - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTerritory(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Territory

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Territory with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Territory using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Territory using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTerritory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTerritory - returns all
//----------------------------------------------------------------------------
func GetAllTerritory()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Territory

	//----------------------------------------------------------------------------
	// Request the ORM to find all Territory
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Territory" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Territory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTerritory", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTerritory - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTerritory(obj model.Territory)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Territory using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Territory using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTerritory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTerritory - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTerritory(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTerritory(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Territory)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Territory using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Territory using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTerritory", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Territory
//----------------------------------------------------------------------------
func AssignOrganizationToTerritory( territoryId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerritory(territoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Territory)

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
			// assign the Organization	to the Territory
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Territory
			//----------------------------------------------------------------------------
			return UpdateTerritory(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Territory
//----------------------------------------------------------------------------
func UnassignOrganizationFromTerritory(territoryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerritory(territoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Territory)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Territory
		//----------------------------------------------------------------------------
		return UpdateTerritory(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Territory
//----------------------------------------------------------------------------
func AddAccountsToTerritory ( territoryId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerritory(territoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Territory)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Territory from the gorm
		//----------------------------------------------------------------------------
		return GetTerritory(territoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Territory
//----------------------------------------------------------------------------
func RemoveAccountsFromTerritory( territoryId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerritory(territoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Territory)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Territory from the gorm
		//----------------------------------------------------------------------------
		return GetTerritory(territoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more usersIds as a Users to a Territory
//----------------------------------------------------------------------------
func AddUsersToTerritory ( territoryId uint64, usersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerritory(territoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Territory)

		// slice the ids on comma with no spaces
		ids := strings.Split( usersIds, ",")

		for _, usersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.User

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a User
			// with a matching usersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , usersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Users using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Users").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Users", usersId )
				return utils.RequestResult{false, msg, "unassignUsers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Territory from the gorm
		//----------------------------------------------------------------------------
		return GetTerritory(territoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more usersIds as a Users from a Territory
//----------------------------------------------------------------------------
func RemoveUsersFromTerritory( territoryId uint64, usersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Territory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerritory(territoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Territory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Territory)

		// slice the ids on comma with no spaces
		ids := strings.Split( usersIds, ",")

		for _, usersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.User

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a User
			// with a matching usersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , usersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove UserObj from the Users array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Users").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Users", usersId )
				return utils.RequestResult{false, msg, "removeUsers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Territory from the gorm
		//----------------------------------------------------------------------------
		return GetTerritory(territoryId)

	} else {
		return parentRequestResult
	}
}

