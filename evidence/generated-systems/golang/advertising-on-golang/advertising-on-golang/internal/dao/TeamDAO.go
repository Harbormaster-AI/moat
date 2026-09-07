package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TeamDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTeam - creates a new db entry
//----------------------------------------------------------------------------
func CreateTeam(obj model.Team)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Team with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Team", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTeam", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTeam - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTeam(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Team

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Team with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Team using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Team using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTeam", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTeam - returns all
//----------------------------------------------------------------------------
func GetAllTeam()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Team

	//----------------------------------------------------------------------------
	// Request the ORM to find all Team
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Team" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Team", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTeam", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTeam - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTeam(obj model.Team)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Team using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Team using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTeam", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTeam - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTeam(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTeam(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Team)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Team using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Team using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTeam", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Agency on a Team
//----------------------------------------------------------------------------
func AssignAgencyToTeam( teamId uint64, agencyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTeam(teamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Team)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Agency

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Agency with a
		// matching agencyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, agencyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Agency	to the Team
			//----------------------------------------------------------------------------
			parentObj.Agency = &childObj

			//----------------------------------------------------------------------------
			// save the Team
			//----------------------------------------------------------------------------
			return UpdateTeam(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agency", agencyId )
			return utils.RequestResult{false, msg, "assignAgency", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Agency on a Team
//----------------------------------------------------------------------------
func UnassignAgencyFromTeam(teamId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTeam(teamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Team)

		//----------------------------------------------------------------------------
		// assign an empty Agency to the Agency
		//----------------------------------------------------------------------------
		parentObj.Agency = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Agency
		//----------------------------------------------------------------------------
		parentObj.AgencyId = nil;

		//----------------------------------------------------------------------------
		// save the Team
		//----------------------------------------------------------------------------
		return UpdateTeam(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more usersIds as a Users to a Team
//----------------------------------------------------------------------------
func AddUsersToTeam ( teamId uint64, usersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTeam(teamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Team)

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
		// retrieve the modified Team from the gorm
		//----------------------------------------------------------------------------
		return GetTeam(teamId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more usersIds as a Users from a Team
//----------------------------------------------------------------------------
func RemoveUsersFromTeam( teamId uint64, usersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTeam(teamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Team)

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
		// retrieve the modified Team from the gorm
		//----------------------------------------------------------------------------
		return GetTeam(teamId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more adAccountsIds as a AdAccounts to a Team
//----------------------------------------------------------------------------
func AddAdAccountsToTeam ( teamId uint64, adAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTeam(teamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Team)

		// slice the ids on comma with no spaces
		ids := strings.Split( adAccountsIds, ",")

		for _, adAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdAccount
			// with a matching adAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AdAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccounts", adAccountsId )
				return utils.RequestResult{false, msg, "unassignAdAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Team from the gorm
		//----------------------------------------------------------------------------
		return GetTeam(teamId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more adAccountsIds as a AdAccounts from a Team
//----------------------------------------------------------------------------
func RemoveAdAccountsFromTeam( teamId uint64, adAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Team with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTeam(teamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Team so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Team)

		// slice the ids on comma with no spaces
		ids := strings.Split( adAccountsIds, ",")

		for _, adAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdAccount
			// with a matching adAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AdAccountObj from the AdAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccounts", adAccountsId )
				return utils.RequestResult{false, msg, "removeAdAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Team from the gorm
		//----------------------------------------------------------------------------
		return GetTeam(teamId)

	} else {
		return parentRequestResult
	}
}

