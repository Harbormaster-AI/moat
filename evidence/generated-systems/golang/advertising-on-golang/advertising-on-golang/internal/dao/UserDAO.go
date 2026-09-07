package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing UserDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateUser - creates a new db entry
//----------------------------------------------------------------------------
func CreateUser(obj model.User)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a User with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a User", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateUser", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetUser - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetUser(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.User

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a User with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a User using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a User using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetUser", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllUser - returns all
//----------------------------------------------------------------------------
func GetAllUser()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.User

	//----------------------------------------------------------------------------
	// Request the ORM to find all User
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all User" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all User", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllUser", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateUser - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateUser(obj model.User)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a User using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a User using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateUser", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteUser - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteUser(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetUser(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.User)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a User using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a User using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteUser", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Agency on a User
//----------------------------------------------------------------------------
func AssignAgencyToUser( userId uint64, agencyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUser(userId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.User)

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
			// assign the Agency	to the User
			//----------------------------------------------------------------------------
			parentObj.Agency = &childObj

			//----------------------------------------------------------------------------
			// save the User
			//----------------------------------------------------------------------------
			return UpdateUser(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agency", agencyId )
			return utils.RequestResult{false, msg, "assignAgency", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Agency on a User
//----------------------------------------------------------------------------
func UnassignAgencyFromUser(userId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUser(userId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.User)

		//----------------------------------------------------------------------------
		// assign an empty Agency to the Agency
		//----------------------------------------------------------------------------
		parentObj.Agency = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Agency
		//----------------------------------------------------------------------------
		parentObj.AgencyId = nil;

		//----------------------------------------------------------------------------
		// save the User
		//----------------------------------------------------------------------------
		return UpdateUser(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more teamsIds as a Teams to a User
//----------------------------------------------------------------------------
func AddTeamsToUser ( userId uint64, teamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUser(userId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.User)

		// slice the ids on comma with no spaces
		ids := strings.Split( teamsIds, ",")

		for _, teamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Team

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Team
			// with a matching teamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , teamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Teams using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Teams").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Teams", teamsId )
				return utils.RequestResult{false, msg, "unassignTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified User from the gorm
		//----------------------------------------------------------------------------
		return GetUser(userId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more teamsIds as a Teams from a User
//----------------------------------------------------------------------------
func RemoveTeamsFromUser( userId uint64, teamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUser(userId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.User)

		// slice the ids on comma with no spaces
		ids := strings.Split( teamsIds, ",")

		for _, teamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Team

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Team
			// with a matching teamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , teamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TeamObj from the Teams array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Teams").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Teams", teamsId )
				return utils.RequestResult{false, msg, "removeTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified User from the gorm
		//----------------------------------------------------------------------------
		return GetUser(userId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more adAccountsIds as a AdAccounts to a User
//----------------------------------------------------------------------------
func AddAdAccountsToUser ( userId uint64, adAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUser(userId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.User)

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
		// retrieve the modified User from the gorm
		//----------------------------------------------------------------------------
		return GetUser(userId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more adAccountsIds as a AdAccounts from a User
//----------------------------------------------------------------------------
func RemoveAdAccountsFromUser( userId uint64, adAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the User with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUser(userId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.User so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.User)

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
		// retrieve the modified User from the gorm
		//----------------------------------------------------------------------------
		return GetUser(userId)

	} else {
		return parentRequestResult
	}
}

