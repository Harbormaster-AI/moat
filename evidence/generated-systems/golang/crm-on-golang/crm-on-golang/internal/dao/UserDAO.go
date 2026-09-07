package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
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
// assigns a Organization on a User
//----------------------------------------------------------------------------
func AssignOrganizationToUser( userId uint64, organizationId uint64 )(utils.RequestResult){

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
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the User
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the User
			//----------------------------------------------------------------------------
			return UpdateUser(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a User
//----------------------------------------------------------------------------
func UnassignOrganizationFromUser(userId uint64)(utils.RequestResult) {

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
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

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
// adds one or more activitiesIds as a Activities to a User
//----------------------------------------------------------------------------
func AddActivitiesToUser ( userId uint64, activitiesIds string )(utils.RequestResult) {

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
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Activities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "unassignActivities", childObj}
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
// removes one or more activitiesIds as a Activities from a User
//----------------------------------------------------------------------------
func RemoveActivitiesFromUser( userId uint64, activitiesIds string )(utils.RequestResult) {
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
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ActivityObj from the Activities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "removeActivities", childObj}
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
// adds one or more ownedAccountsIds as a OwnedAccounts to a User
//----------------------------------------------------------------------------
func AddOwnedAccountsToUser ( userId uint64, ownedAccountsIds string )(utils.RequestResult) {

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
		ids := strings.Split( ownedAccountsIds, ",")

		for _, ownedAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching ownedAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OwnedAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedAccounts", ownedAccountsId )
				return utils.RequestResult{false, msg, "unassignOwnedAccounts", childObj}
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
// removes one or more ownedAccountsIds as a OwnedAccounts from a User
//----------------------------------------------------------------------------
func RemoveOwnedAccountsFromUser( userId uint64, ownedAccountsIds string )(utils.RequestResult) {
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
		ids := strings.Split( ownedAccountsIds, ",")

		for _, ownedAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching ownedAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the OwnedAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedAccounts", ownedAccountsId )
				return utils.RequestResult{false, msg, "removeOwnedAccounts", childObj}
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
// adds one or more ownedLeadsIds as a OwnedLeads to a User
//----------------------------------------------------------------------------
func AddOwnedLeadsToUser ( userId uint64, ownedLeadsIds string )(utils.RequestResult) {

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
		ids := strings.Split( ownedLeadsIds, ",")

		for _, ownedLeadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Lead

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Lead
			// with a matching ownedLeadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedLeadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OwnedLeads using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedLeads").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedLeads", ownedLeadsId )
				return utils.RequestResult{false, msg, "unassignOwnedLeads", childObj}
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
// removes one or more ownedLeadsIds as a OwnedLeads from a User
//----------------------------------------------------------------------------
func RemoveOwnedLeadsFromUser( userId uint64, ownedLeadsIds string )(utils.RequestResult) {
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
		ids := strings.Split( ownedLeadsIds, ",")

		for _, ownedLeadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Lead

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Lead
			// with a matching ownedLeadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedLeadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LeadObj from the OwnedLeads array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedLeads").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedLeads", ownedLeadsId )
				return utils.RequestResult{false, msg, "removeOwnedLeads", childObj}
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
// adds one or more ownedOpportunitiesIds as a OwnedOpportunities to a User
//----------------------------------------------------------------------------
func AddOwnedOpportunitiesToUser ( userId uint64, ownedOpportunitiesIds string )(utils.RequestResult) {

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
		ids := strings.Split( ownedOpportunitiesIds, ",")

		for _, ownedOpportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching ownedOpportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedOpportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OwnedOpportunities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedOpportunities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedOpportunities", ownedOpportunitiesId )
				return utils.RequestResult{false, msg, "unassignOwnedOpportunities", childObj}
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
// removes one or more ownedOpportunitiesIds as a OwnedOpportunities from a User
//----------------------------------------------------------------------------
func RemoveOwnedOpportunitiesFromUser( userId uint64, ownedOpportunitiesIds string )(utils.RequestResult) {
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
		ids := strings.Split( ownedOpportunitiesIds, ",")

		for _, ownedOpportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching ownedOpportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedOpportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityObj from the OwnedOpportunities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedOpportunities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedOpportunities", ownedOpportunitiesId )
				return utils.RequestResult{false, msg, "removeOwnedOpportunities", childObj}
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
// adds one or more ownedCasesIds as a OwnedCases to a User
//----------------------------------------------------------------------------
func AddOwnedCasesToUser ( userId uint64, ownedCasesIds string )(utils.RequestResult) {

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
		ids := strings.Split( ownedCasesIds, ",")

		for _, ownedCasesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Case_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Case_
			// with a matching ownedCasesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedCasesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OwnedCases using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedCases").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedCases", ownedCasesId )
				return utils.RequestResult{false, msg, "unassignOwnedCases", childObj}
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
// removes one or more ownedCasesIds as a OwnedCases from a User
//----------------------------------------------------------------------------
func RemoveOwnedCasesFromUser( userId uint64, ownedCasesIds string )(utils.RequestResult) {
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
		ids := strings.Split( ownedCasesIds, ",")

		for _, ownedCasesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Case_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Case_
			// with a matching ownedCasesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownedCasesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Case_Obj from the OwnedCases array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OwnedCases").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OwnedCases", ownedCasesId )
				return utils.RequestResult{false, msg, "removeOwnedCases", childObj}
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
// adds one or more quotesIds as a Quotes to a User
//----------------------------------------------------------------------------
func AddQuotesToUser ( userId uint64, quotesIds string )(utils.RequestResult) {

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
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Quotes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "unassignQuotes", childObj}
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
// removes one or more quotesIds as a Quotes from a User
//----------------------------------------------------------------------------
func RemoveQuotesFromUser( userId uint64, quotesIds string )(utils.RequestResult) {
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
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QuoteObj from the Quotes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "removeQuotes", childObj}
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
// adds one or more ordersIds as a Orders to a User
//----------------------------------------------------------------------------
func AddOrdersToUser ( userId uint64, ordersIds string )(utils.RequestResult) {

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
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Order

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Order
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Orders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "unassignOrders", childObj}
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
// removes one or more ordersIds as a Orders from a User
//----------------------------------------------------------------------------
func RemoveOrdersFromUser( userId uint64, ordersIds string )(utils.RequestResult) {
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
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Order

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Order
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OrderObj from the Orders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "removeOrders", childObj}
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
// adds one or more contractsIds as a Contracts to a User
//----------------------------------------------------------------------------
func AddContractsToUser ( userId uint64, contractsIds string )(utils.RequestResult) {

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
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contracts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "unassignContracts", childObj}
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
// removes one or more contractsIds as a Contracts from a User
//----------------------------------------------------------------------------
func RemoveContractsFromUser( userId uint64, contractsIds string )(utils.RequestResult) {
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
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ContractObj from the Contracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "removeContracts", childObj}
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
// adds one or more emailMessagesIds as a EmailMessages to a User
//----------------------------------------------------------------------------
func AddEmailMessagesToUser ( userId uint64, emailMessagesIds string )(utils.RequestResult) {

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
		ids := strings.Split( emailMessagesIds, ",")

		for _, emailMessagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailMessagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailMessagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EmailMessages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmailMessages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmailMessages", emailMessagesId )
				return utils.RequestResult{false, msg, "unassignEmailMessages", childObj}
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
// removes one or more emailMessagesIds as a EmailMessages from a User
//----------------------------------------------------------------------------
func RemoveEmailMessagesFromUser( userId uint64, emailMessagesIds string )(utils.RequestResult) {
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
		ids := strings.Split( emailMessagesIds, ",")

		for _, emailMessagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailMessagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailMessagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmailMessageObj from the EmailMessages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmailMessages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmailMessages", emailMessagesId )
				return utils.RequestResult{false, msg, "removeEmailMessages", childObj}
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

